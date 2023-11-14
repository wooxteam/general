package handlers

import (
	"fmt"

	"github.com/wooxteam/general/chain"
)

type second struct {
	chain.CommonHandler[*SomeEvent]
}

func NewSecond() *second {
	canHandleFn := func(event *SomeEvent) bool {
		if event.Code == 2 {
			return true
		}

		return false
	}

	return &second{CommonHandler: chain.NewCommonHandler[*SomeEvent](canHandleFn, wrapNotFoundErr)}
}

func (f *second) Handle(event *SomeEvent) chain.HandlerNotFoundError {
	return f.CommonHandler.Handle(event, f.handle)
}

func (f *second) handle(event *SomeEvent) {
	fmt.Println("second")
}
