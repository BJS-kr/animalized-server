package packet

import "errors"

func stripDelimiter(chunk *[]byte) error {
	c := *chunk
	l := len(c)

	if c[l-1] != INPUT_PACKET_DELIMITER {
		return errors.New("delimiter not on last position")
	}
	*chunk = c[:l-1]
	return nil
}
