package packet

const (
	INIT = iota + 1
	MOVE
	ATTACK
	SERVER_STATE
	JOIN
	QUIT
)

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
