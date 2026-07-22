package handler

import ()

type Service interface {
}

type Handler struct {
	service Service
}

func New(svc Service) *Handler {
	return &Handler{service: svc}
}
