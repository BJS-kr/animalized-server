package packet

const (
	INIT = iota + 1
	MOVE
	ATTACK
	BUFFER_SIZE            = 4096
	INPUT_PACKET_DELIMITER = '$'
)
