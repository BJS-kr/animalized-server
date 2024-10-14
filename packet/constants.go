package packet

const (
	BUFFER_SIZE            = 4096
	INPUT_PACKET_DELIMITER = '$'
	INIT                   = iota + 1
	MOVE
	ATTACK
)
