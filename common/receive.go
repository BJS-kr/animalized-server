package common

import (
	"log/slog"
)

func (b *Base) Receive(handler Handler) {
	for input := range b.InputChannel {
		select {
		case <-b.Stop:
			return
		default:
			input, err := handler(input)

			if err != nil {
				slog.Error(err.Error())
				continue
			}

			b.Inputs.Enqueue(input)
		}
	}
}
