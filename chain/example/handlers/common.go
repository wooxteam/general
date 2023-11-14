package handlers

import (
	"fmt"

	"github.com/wooxteam/general/chain"
)

type SomeEvent struct {
	Code int
}

func wrapNotFoundErr(err chain.HandlerNotFoundError, event *SomeEvent) chain.HandlerNotFoundError {
	return fmt.Errorf("handle event id=[%d] skipped, %w", event.Code, err)
}
