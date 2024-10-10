package packet

import (
	"animalized/message"
	"bytes"
	"net"
)

// 유저로부터 수집된 인풋들을 중계 스택으로 쌓는다.
// 복수형인 이유는 패킷파싱할 때 일단 커넥션 타고 있는 것들은 싹 순회하고 넣을 예정이라서
// 일단 지금 생각으로는 buffered channel로 넣으면 될 것 같다.
// 패킷 타입은 하나로 통일한다. 로직을 간단화하고 시간순서 맞추기도 편하다.
func ParseInput(conn *net.TCPConn)(*message.Input, error) {
	inputBuf := bytes.NewBuffer(nil)
	buf := make([]byte, BUFFER_SIZE)
	chunk, err := makeChunk(conn, buf, inputBuf)

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