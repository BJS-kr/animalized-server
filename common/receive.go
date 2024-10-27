package common

import (
	"log/slog"
)

func (d *Distributable) Receive(handler Handler) {
	for input := range d.InputChannel {
		select {
		case <-d.Stop:
			return
		default:
			input, err := handler(input)

			if err != nil {
				slog.Error(err.Error())
				continue
			}

			if input == nil {
				continue
			}

			d.Inputs.Enqueue(input)
		}
	}
}
