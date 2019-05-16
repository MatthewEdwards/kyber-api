package app

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (app *application) handleGetAricles() http.HandlerFunc {
	type article struct {
		Title string `json:"title"`
		URL   string `json:"url"`
	}

	result := &article{
		Title: "WireGuard VPN review: A new type of VPN offers serious advantages",
		URL:   "https://arstechnica.com/gadgets/2018/08/wireguard-vpn-review-fast-connections-amaze-but-windows-support-needs-to-happen/",
	}

	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Returning all articles")
		response(w, result, 200)
	}
}
