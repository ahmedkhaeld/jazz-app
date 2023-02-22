package handlers

import (
	"github.com/ahmedkhaeld/jazz"
	"myapp/data"
	"net/http"
	"time"
)

type Handlers struct {
	*jazz.Jazz
	data.Models
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	defer h.LoadTime(time.Now())
	err := h.Render.JetPage(w, r, "home", nil, nil)
	if err != nil {
		h.ErrorLog.Println("error rendering:", err)
	}
}
