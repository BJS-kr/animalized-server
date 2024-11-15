package packet

const (
	UP = iota + 1
	DOWN
	LEFT
	RIGHT
)

const (
	BUFFER_SIZE            = 4096
	INPUT_PACKET_DELIMITER = '$'
)
