package server

import (
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/arisro/talos/config"
	"github.com/julienschmidt/httprouter"
	"github.com/meatballhat/negroni-logrus"
	"github.com/spf13/cobra"
	"github.com/urfave/negroni"
)

func StartServer() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		router := httprouter.New()
		serverHandler := &Handler{}
		serverHandler.registerRoutes(router)

		n := negroni.New()
		n.Use(negronilogrus.NewMiddleware())
		n.UseHandler(router)

		var srv = http.Server{
			Addr:         "127.0.0.1:5255",
			Handler:      n,
			ReadTimeout:  time.Second * 5,
			WriteTimeout: time.Second * 10,
		}

		logrus.Infof("Setting up http server")
		err := srv.ListenAndServe()

		if err != nil {
			logrus.Error("Cannot start server %s.", err)
		}
	}
}

type Handler struct {
	Config *config.Config
}

func (h *Handler) registerRoutes(router *httprouter.Router) {
	// c := h.Config
	// ctx := c.Context

	router.GET("/health", func(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		rw.WriteHeader(http.StatusNoContent)
	})
}
