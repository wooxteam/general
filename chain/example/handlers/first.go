package handlers

import (
	"fmt"

	"github.com/wooxteam/general/chain"
)

type first struct {
	chain.CommonHandler[*SomeEvent]
}

func NewFirst() *first {
	canHandleFn := func(event *SomeEvent) bool {
		if event.Code == 1 {
			return true
		}

		return false
	}

	return &first{
		CommonHandler: chain.NewCommonHandler[*SomeEvent](canHandleFn, wrapNotFoundErr),
	}
}

func (f *first) Handle(event *SomeEvent) chain.HandlerNotFoundError {
	return f.CommonHandler.Handle(event, f.handle)
}

func (f *first) handle(event *SomeEvent) {
	fmt.Println("first")
}
