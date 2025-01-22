package handler

import (
	"image/png"
	"net/http"

	"github.com/osamikoyo/qrk/internal/utils"
	"github.com/osamikoyo/qrk/pkg/loger"
)

type Handler struct{}

func New() *Handler {
	return new(Handler)
}

type handler func (w http.ResponseWriter, r *http.Request) error 

func errorRoute(h handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r);err != nil{
			loger.New().Error().Err(err)
		}
	}
}

func (h *Handler) getQR(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet{
		r.Response.StatusCode = http.StatusMethodNotAllowed
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
	
	title := r.URL.Query().Get("link")
	code := utils.GetQR(title, &w)

	return png.Encode(w, code)
}

func RegisterRoute(mux *http.ServeMux){
	handler := New()
	mux.Handle("/qr", errorRoute(handler.getQR))
}