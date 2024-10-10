package packet

import "time"

const (
	READ_DEADLINE          = time.Duration(time.Minute)
	BUFFER_SIZE            = 4096
	INPUT_PACKET_DELIMITER = '$'
)


