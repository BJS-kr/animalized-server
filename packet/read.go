package packet

import (
	"errors"
	"io"
	"net"
)

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
