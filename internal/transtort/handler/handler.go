package handler

import (
	"net/http"

	"github.com/osamikoyo/qrk/pkg/loger"
)

type handler func (w http.ResponseWriter, r *http.Request) error 

func errorRoute(h handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r);err != nil{
			loger.New().Error().Err(err)
		}
	}
}