package handler_test

// import (
// 	"animalized/handler"
// 	"animalized/message"
// 	"animalized/packet"
// 	"animalized/queue"
// 	"animalized/state"
// 	"animalized/users"
// 	"net"
// 	"testing"
// 	"time"

// 	"google.golang.org/protobuf/proto"
// )

// // TODO handler는 종단이다. 하위 모듈에서 변경이 일어나고 있기 때문에 구조가 완성된 후 handler의 테스트를 수정한다.
// func TestHandlers(t *testing.T) {
// 	users := new(users.Users)
// 	gameState := new(state.GameState)
// 	server1, client1 := net.Pipe()
// 	server2, client2 := net.Pipe()
// 	inputQueue := queue.New[*message.Input]()
// 	inputProduceChan := make(chan *message.Input)

// 	// 글로벌 핸들러 시작
// 	go handler.ReceiveGameInput(inputQueue, gameState, inputProduceChan)
// 	go handler.Propagate(inputQueue, users)

// 	// 유저 핸들러 시작
// 	go func() {
// 		handler.StartHandlers(users, gameState, server1, inputProduceChan)
// 		handler.StartHandlers(users, gameState, server2, inputProduceChan)
// 	}()

// 	// 두 명의 유저가 접속합니다.
// 	userInit1 := &message.Input{
// 		Type:   packet.INIT,
// 		UserId: "test1",
// 	}
// 	userInit2 := &message.Input{
// 		Type:   packet.INIT,
// 		UserId: "test2",
// 	}

// 	message1, _ := proto.Marshal(userInit1)
// 	message2, _ := proto.Marshal(userInit2)

// 	_, err := client1.Write(append(message1, packet.INPUT_PACKET_DELIMITER))

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	_, err = client2.Write(append(message2, packet.INPUT_PACKET_DELIMITER))

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// 첫 번째 유저(test1)가 이동을 합니다.
// 	direction := int32(1)
// 	userMove := &message.Input{
// 		Type:      packet.MOVE,
// 		UserId:    "test1",
// 		Direction: &direction,
// 	}
// 	moveMessage, _ := proto.Marshal(userMove)
// 	_, err = client1.Write(append(moveMessage, packet.INPUT_PACKET_DELIMITER))

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// 두 번째 유저(test2)유저가 첫 번째 유저의 움직임 패킷을 받습니다.
// 	buf := make([]byte, packet.BUFFER_SIZE)
// 	receivedInput := new(message.Input)

// 	for {
// 		client2.SetReadDeadline(time.Now().Add(time.Second * 2))

// 		size, err := client2.Read(buf)

// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		if size > 0 {
// 			proto.Unmarshal(buf[:size-1], receivedInput)
// 			break
// 		}
// 	}

// 	// 두 번째 유저가 받은 패킷이 첫 번째 유저가 보낸 인풋과 일치하는지 검사합니다.
// 	if receivedInput.UserId != "test1" ||
// 		receivedInput.Type != packet.MOVE ||
// 		*receivedInput.Direction != 1 {
// 		t.Errorf("data not matched. actual: %s, %d, %d", receivedInput.UserId, receivedInput.Type, *receivedInput.Direction)
// 	}
// }
