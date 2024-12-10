package packet

import (
	"animalized/message"
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"

	"google.golang.org/protobuf/proto"
)

type PacketStore struct {
	incomingBuf   []byte
	inputBuf      *bytes.Buffer
	packetSizeBuf []byte
	targetBuf     []byte
}

func NewStore() *PacketStore {
	ps := new(PacketStore)
	ps.incomingBuf = make([]byte, BUFFER_SIZE)
	ps.inputBuf = bytes.NewBuffer(nil)
	ps.packetSizeBuf = make([]byte, 2)
	ps.targetBuf = make([]byte, BUFFER_SIZE)

	return ps
}

func (ps *PacketStore) ParseInput(conn net.Conn) (*message.Input, error) {
	chunk, err := ps.makeChunk(conn)

	if err != nil {
		return nil, err
	}

	input := new(message.Input)
	if err := proto.Unmarshal(chunk, input); err != nil {
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

		if size == 0 {
			return chunk, errors.New("connection closed")
		}

		if err != nil {
			return ps.incomingBuf, err
		}

		if err := ps.writeInput(size); err != nil {
			return ps.incomingBuf, err
		}
	}
}

func (ps *PacketStore) cutChunk() ([]byte, error) {
	readSizeOfPacket, err := ps.inputBuf.Read(ps.packetSizeBuf)

	if err != nil {
		if errors.Is(err, io.EOF) {
			ps.inputBuf.Write(ps.packetSizeBuf[:readSizeOfPacket])
		}

		return ps.packetSizeBuf, err
	}

	packetSize := binary.BigEndian.Uint16(ps.packetSizeBuf)

	if packetSize > BUFFER_SIZE {
		return ps.targetBuf, errors.New("received packet size is bigger than BUFFER_SIZE")
	}

	targetBuf := ps.targetBuf[:packetSize]
	readPacketSize, err := ps.inputBuf.Read(targetBuf)

	if err != nil {
		if errors.Is(err, io.EOF) {
			ps.inputBuf.Write(targetBuf[:readPacketSize])
		}

		return targetBuf, err
	}

	return targetBuf, nil
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
