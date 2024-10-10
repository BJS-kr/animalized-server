package packet

import "errors"

func stripDelimiter(chunk []byte) ([]byte, error) {
	l := len(chunk)
	
	if chunk[l-1] != INPUT_PACKET_DELIMITER {
		return chunk, errors.New("delimiter not on last position")
	}

	return chunk[:l-1], nil
}