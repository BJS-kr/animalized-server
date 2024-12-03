package state

import "time"

const (
	MAP_SIZE                     = 20
	CLIENT_CELL_SIZE             = 27
	SERVER_STATE_SIGNAL_INTERVAL = time.Second
	MAX_SPACE                    = MAP_SIZE * CLIENT_CELL_SIZE
)
