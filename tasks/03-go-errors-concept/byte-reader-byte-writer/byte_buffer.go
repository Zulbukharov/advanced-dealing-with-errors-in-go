package bytebuffer

import "fmt"

const bufferMaxSize = 1024

type MaxSizeExceededError struct {
	desiredLen int
}

func (e *MaxSizeExceededError) Error() string {
	return fmt.Sprintf("buffer max size exceeded: %d > %d", e.desiredLen, bufferMaxSize)
}

type EndOfBufferError struct{}

func (e *EndOfBufferError) Error() string {
	return "end of buffer"
}

type ByteBuffer struct {
	// buffer представляет собой непосредственно буфер: содержит какие-то данные.
	buffer []byte
	// offset представляет собой смещение, указывающее на первый непрочитанный байт.
	offset int
}

// Необходимо сделать так, чтобы тип *ByteBuffer реализовывал интерфейсы io.ByteWriter и io.ByteReader.
//
// Метод WriteByte должен возвращать ошибку *MaxSizeExceededError при попытке записи в буфер,
// если в нём уже больше bufferMaxSize байт.
//
// Метод ReadByte должен возвращать ошибку *EndOfBufferError при попытке чтения из буфера,
// если ранее буфер уже был вычитан полностью.

func (b *ByteBuffer) WriteByte(c byte) error {
	if len(b.buffer)+1 > bufferMaxSize {
		return &MaxSizeExceededError{desiredLen: len(b.buffer) + 1}
	}
	b.buffer = append(b.buffer, c)
	return nil
}

func (b *ByteBuffer) ReadByte() (byte, error) {
	if b.offset == len(b.buffer) {
		return 0, &EndOfBufferError{}
	}
	c := b.buffer[b.offset]
	b.offset += 1
	return c, nil
}
