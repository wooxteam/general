package main

import (
	"errors"
	"fmt"
)

var ErrHandlerNotFound HandlerNotFoundError = errors.New("handler not found")

type HandlerNotFoundError interface {
	error
}

type (
	handleFn[T any]          func(event T)
	canHandleFn[T any]       func(event T) bool
	wrapNotFoundErrFn[T any] func(err HandlerNotFoundError, event T) HandlerNotFoundError
)

type Handler[T any] interface {
	Handle(event T) HandlerNotFoundError
	Next() Handler[T]
	SetNext(handler Handler[T]) error
}

type CommonHandler[T any] struct {
	nextHandler       Handler[T]
	canHandleFn       func(event T) bool
	wrapNotFoundErrFn func(err HandlerNotFoundError, event T) HandlerNotFoundError
}

func NewCommonHandler[T any](canHandleFn canHandleFn[T], wrapNotFoundErrFn wrapNotFoundErrFn[T]) CommonHandler[T] {
	return CommonHandler[T]{
		canHandleFn:       canHandleFn,
		wrapNotFoundErrFn: wrapNotFoundErrFn,
	}
}

func (h *CommonHandler[T]) Handle(event T, fn handleFn[T]) HandlerNotFoundError {
	if h.canHandleFn(event) {
		fn(event)
		return nil
	} else if h.nextHandler != nil {
		return h.nextHandler.Handle(event)
	}

	if h.wrapNotFoundErrFn != nil {
		return h.wrapNotFoundErrFn(ErrHandlerNotFound, event)
	}

	return ErrHandlerNotFound
}

func (h *CommonHandler[T]) SetNext(handler Handler[T]) error {
	if h.nextHandler != nil {
		return errors.New("next handler already set")
	}

	h.nextHandler = handler

	return nil
}

func (h *CommonHandler[T]) Next() Handler[T] {
	return h.nextHandler
}

func MakeChainHandler[T any](handlers ...Handler[T]) (Handler[T], error) {
	for i := 0; i < len(handlers)-1; i++ {
		err := handlers[i].SetNext(handlers[i+1])
		if err != nil {
			return nil, fmt.Errorf("make chain handlers fail: %w", err)
		}
	}

	return handlers[0], nil
}
