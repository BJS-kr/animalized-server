package producer

import (
	"animalized/message"
	"animalized/packet"
	"animalized/user"
	"bytes"
	"errors"
	"time"
)

// 유저로부터 수집된 인풋들을 중계 스택으로 쌓는다.
// 복수형인 이유는 패킷파싱할 때 일단 커넥션 타고 있는 것들은 싹 순회하고 넣을 예정이라서
// 일단 지금 생각으로는 buffered channel로 넣으면 될 것 같다.
// 패킷 타입은 하나로 통일한다. 로직을 간단화하고 시간순서 맞추기도 편하다.
func ProduceInput(u *user.User, buf *[]byte, inputBuf *bytes.Buffer, inputProduceChannel chan<- *message.Input) error {
	input, err := packet.ParseInput(u.Conn, buf, inputBuf)

	if err != nil {
		return err
	}

	if input.UserId != u.Id {
		return errors.New("user id not matched")
	}

	if err := u.Conn.SetReadDeadline(time.Now().Add(READ_DEADLINE)); err != nil {
		return err
	}

	inputProduceChannel <- input

	return nil
}
