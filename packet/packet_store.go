package packet

import (
	"animalized/message"
	"bytes"
	"errors"
	"io"
	"net"
)

type PacketStore struct {
	incomingBuf []byte
	inputBuf    *bytes.Buffer
}

func NewStore() *PacketStore {
	ps := new(PacketStore)
	ps.incomingBuf = make([]byte, BUFFER_SIZE)
	ps.inputBuf = bytes.NewBuffer(nil)

	return ps
}

func (ps *PacketStore) ParseInput(conn net.Conn) (*message.Input, error) {
	chunk, err := ps.makeChunk(conn)

	if err != nil {
		return nil, err
	}

	stripped, err := stripDelimiter(chunk)

	if err != nil {
		return nil, err
	}

	input := new(message.Input)
	err = into(input, stripped)

	if err != nil {
		return nil, err
	}

	return input, nil
}

func (ps *PacketStore) makeChunk(conn net.Conn) ([]byte, error) {
	for {
		chunk, err := ps.cutChunk()

		if err == nil {
			return chunk, nil
		}

		if !errors.Is(err, io.EOF) {
			return chunk, err
		}

		size, err := ps.readInput(conn)

		if err != nil {
			return ps.incomingBuf, err
		}

		if err := ps.writeInput(size); err != nil {
			return ps.incomingBuf, err
		}
	}
}

func (ps *PacketStore) cutChunk() ([]byte, error) {
	chunk, err := ps.inputBuf.ReadBytes(INPUT_PACKET_DELIMITER)

	if err != nil {
		if errors.Is(err, io.EOF) {
			ps.inputBuf.Write(chunk)
		}

		return chunk, err
	}

	return chunk, nil
}

func (ps *PacketStore) readInput(conn net.Conn) (int, error) {
	size, err := conn.Read(ps.incomingBuf)

	if err != nil {
		if errors.Is(err, io.EOF) {
			// 예측된 에러이고, 버퍼 내용은 buf에 쌓였음
			return size, nil
		}

		return size, err
	}

	if size > BUFFER_SIZE {
		return size, errors.New("read size cannot exceed predefined buffer size")
	}

	return size, nil
}

// incomingBuf: ReadIncoming함수에서 buf의 size만큼 slice한 []byte
// ReadIncoming의 size가 0보다 클 때만 호출
func (ps *PacketStore) writeInput(size int) error {
	targetBuf := ps.incomingBuf[:size]

	if _, err := ps.inputBuf.Write(targetBuf); err != nil {
		return err
	}

	return nil
}
