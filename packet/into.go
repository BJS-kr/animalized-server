package packet

import "google.golang.org/protobuf/proto"

func into[M proto.Message](target M, stripped []byte) error {
	if err := proto.Unmarshal(stripped, target); err != nil {
		return err
	}

	return nil
}
