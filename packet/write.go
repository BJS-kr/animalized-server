package packet

// incomingBuf: ReadIncoming함수에서 buf의 size만큼 slice한 []byte
// ReadIncoming의 size가 0보다 클 때만 호출
func (ps *PacketStore) writeInput(size int) error {
	targetBuf := ps.incomingBuf[:size]

	if _, err := ps.inputBuf.Write(targetBuf); err != nil {
		return err
	}

	return nil
}
