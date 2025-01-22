package handler

import (
	"errors"
	"fmt"
	"net/http"

	home "github.com/osamikoyo/qrk/internal/view/page"
)

func (h Handler) homePageHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet{
		r.Response.Status = fmt.Sprintf("%d", http.StatusMethodNotAllowed)
		http.Error(w, "method must be get", http.StatusMethodNotAllowed)
		return errors.New("method must be get")
	}

	return home.HomePage().Render(r.Context(), w)
}

