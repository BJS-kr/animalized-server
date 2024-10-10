package packet

import "bytes"

// incomingBuf: ReadIncoming함수에서 buf의 size만큼 slice한 []byte
// ReadIncoming의 size가 0보다 클 때만 호출
func writeInput(incomingBuf []byte, inputBuf *bytes.Buffer) error {
	if _, err := inputBuf.Write(incomingBuf); err != nil {
		return err
	}

	return nil
}