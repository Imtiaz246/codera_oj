package handler

import (
	"github.com/imtiaz246/codera_oj/app/store"
)

type Handler struct {
	*store.Store
}

func NewHandler() (*Handler, error) {
	s, err := store.NewStore()
	if err != nil {
		return nil, err
	}

	return &Handler{
		s,
	}, nil
}
