package encrypt

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// Reader decrypts and encrypted datastream from Reader.
type Reader struct {
	aeadCommon

	// Underlying io.Reader
	Reader io.Reader

	// Caches leftovers from unread blocks
	backlog *bytes.Reader

	// Last index of the byte the user visited.
	// (Used to avoid re-reads in Seek())
	// This does *not* equal the seek offset of the underlying stream.
	lastSeekPos int64
}

// Read from source and decrypt.
//
// This method always decrypts one block to optimize for continous reads. If
// dest is too small to hold the block, the decrypted text is cached for the
// next read.
func (r *Reader) Read(dest []byte) (int, error) {
	readBytes := 0

	// Try our best to fill len(dest)
	for readBytes < len(dest) {
		if r.backlog.Len() == 0 {
			if _, err := r.readBlock(); err != nil {
				return readBytes, err
			}
		}

		n, _ := r.backlog.Read(dest[readBytes:])
		readBytes += n
		r.lastSeekPos += int64(n)
	}

	return readBytes, nil
}

// Fill internal buffer with current block
func (r *Reader) readBlock() (int, error) {
	if n, err := r.Reader.Read(r.nonce); err != nil {
		return 0, err
	} else if n != r.aead.NonceSize() {
		return 0, fmt.Errorf("Nonce size mismatch. Should: %d. Have: %d",
			r.aead.NonceSize(), n)
	}

	// Read the *whole* block from the fs
	N := MaxBlockSize + r.aead.Overhead()
	n, err := io.ReadAtLeast(r.Reader, r.encBuf[:N], N)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return 0, err
	}

	r.decBuf, err = r.aead.Open(r.decBuf[:0], r.nonce, r.encBuf[:n], nil)
	if err != nil {
		return 0, err
	}

	r.backlog = bytes.NewReader(r.decBuf)
	return len(r.decBuf), nil
}

// Seek into the encrypted stream.
//
// Note that the seek offset is relative to the decrypted data,
// not to the underlying, encrypted stream.
//
// Mixing SEEK_CUR and SEEK_SET might not a good idea,
// since a seek might involve reading a whole encrypted block.
// Therefore relative seek offset
func (r *Reader) Seek(offset int64, whence int) (int64, error) {
	// Check if seeking is supported:
	seeker, ok := r.Reader.(io.ReadSeeker)
	if !ok {
		return 0, fmt.Errorf("Seek is not supported by underlying datastream")
	}

	// Constants and assumption on the stream below:
	blockSize := int64(MaxBlockSize)
	blockHeaderSize := int64(r.aead.NonceSize())
	totalBlockSize := blockHeaderSize + blockSize + int64(r.aead.Overhead())

	// absolute Offset in the decrypted stream
	absOffsetDec := int64(0)

	// Convert possibly relative offset to absolute offset:
	switch whence {
	case os.SEEK_CUR:
		absOffsetDec = r.lastSeekPos + offset
	case os.SEEK_SET:
		absOffsetDec = offset
	case os.SEEK_END:
		// We have no idea when the stream ends.
		return 0, fmt.Errorf("SEEK_END is not supported for encrypted data")
	}

	if r.lastSeekPos == absOffsetDec {
		// Nothing changed, why bother?
		return r.lastSeekPos, nil
	}

	if absOffsetDec < 0 {
		return 0, fmt.Errorf("Negative seek index")
	}

	// Caller wanted to know only the current stream pos:
	if absOffsetDec == r.lastSeekPos {
		return absOffsetDec, nil
	}

	// Convert decrypted offset to encrypted offset
	absOffsetEnc := headerSize + ((absOffsetDec / blockSize) * totalBlockSize)

	// Check if we're still in the same block as last time:
	blockNum := absOffsetEnc / totalBlockSize
	lastBlockNum := r.lastSeekPos / blockSize
	r.lastSeekPos = absOffsetDec

	if lastBlockNum != blockNum {
		// Seek to the beginning of the encrypted block:
		if _, err := seeker.Seek(absOffsetEnc, os.SEEK_SET); err != nil {
			return 0, err
		}

		// Make read consume the current block:
		if _, err := r.readBlock(); err != nil {
			return 0, err
		}
	}

	// Reslice the backlog, so Read() does not return skipped data.
	r.backlog.Seek(absOffsetDec%blockSize, os.SEEK_SET)
	return absOffsetDec, nil
}

// Close does finishing work.
// If supported by the underlying data stream, it closes it too.
// This is currently a No-Op, but you should not rely on that.
func (r *Reader) Close() error {
	// Check
	closer, ok := r.Reader.(io.Closer)
	if ok {
		return closer.Close()
	}

	return nil
}

// NewReader creates a new encrypted reader and validates the file header.
// The key is required to be KeySize bytes long.
func NewReader(r io.Reader, key []byte) (*Reader, error) {
	reader := &Reader{
		Reader:  r,
		backlog: bytes.NewReader([]byte{}),
	}

	header := make([]byte, headerSize)
	n, err := reader.Reader.Read(header)
	if err != nil {
		return nil, err
	}

	if n != headerSize {
		return nil, fmt.Errorf("No valid header found, damaged file?")
	}

	version, ciperType, keylen, _, err := ParseHeader(header)
	if err != nil {
		return nil, err
	}

	if version != 1 {
		return nil, fmt.Errorf("This implementation does not support versions != 1")
	}

	if uint32(len(key)) != keylen {
		return nil, fmt.Errorf("Key length differs: file=%d, user=%d", keylen, len(key))
	}

	if err := reader.initAeadCommon(key, ciperType); err != nil {
		return nil, err
	}

	return reader, nil
}