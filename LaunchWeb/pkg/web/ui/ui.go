package ui

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/epyphite/realip"
	"github.com/spf13/viper"
)

func UI() http.Handler {

	assetsPath := viper.GetString("assets-path")

	var filesystem http.FileSystem
	if assetsPath != "" {
		log.Println("using ui assets path:", assetsPath)
		filesystem = http.Dir(assetsPath)
	} else {
		log.Println("using the built-in ui assets")
		filesystem = assetFS()
	}

	h := http.FileServer(filesystem)
	return &idpUI{h, h}
}

type idpUI struct {
	h             http.Handler
	prefixHandler http.Handler
}

func (s *idpUI) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	clientIP := realip.FromRequest(req)
	if "/favicon.ico" == req.URL.Path {
		s.h.ServeHTTP(w, req)
		return
	}
	log.Info(req.URL.Path)
	log.Info(clientIP)

	if "/ui/index.html" == req.URL.Path {
		// 5 minute cache for the login HTML page
		log.Println(req.URL.Path)
		w.Header().Add("Cache-Control", "public, max-age=600")
	} else {
		log.Println("UI")
		// Encourage caching of UI
		w.Header().Add("Cache-Control", "public, max-age=31536000")
	}
	s.prefixHandler.ServeHTTP(w, req)
}
