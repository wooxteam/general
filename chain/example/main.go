package main

import (
	"errors"
	"fmt"
	"github.com/wooxteam/general/chain"
	"github.com/wooxteam/general/chain/example/handlers"
)

func main() {
	events := make([]*handlers.SomeEvent, 3)
	for i := 0; i < 3; i++ {
		events[i] = &handlers.SomeEvent{Code: i + 1}
	}

	hs := []chain.Handler[*handlers.SomeEvent]{handlers.NewFirst(), handlers.NewSecond()}
	h, _ := chain.MakeChainHandler(hs...)

	for _, e := range events {
		if err := h.Handle(e); err != nil {
			if errors.Is(err, chain.ErrHandlerNotFound) {
				fmt.Println("warning:", err)
			} else {
				fmt.Println("error:", err)
			}

		}
	}
}
