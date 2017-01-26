package server

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
	"github.com/meatballhat/negroni-logrus"
	"net/http"
	"time"
	"github.com/Sirupsen/logrus"
)

func StartServer() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println("starting server")

		router := httprouter.New()
		serverHandler := &Handler{}
		serverHandler.registerRoutes(router)

		n := negroni.New()
		n.Use(negronilogrus.NewMiddleware())
		n.UseHandler(router)

		var srv = http.Server{
			Addr: "127.0.0.1:5255",
			Handler: n,
			ReadTimeout: time.Second * 5,
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
}

func (h *Handler) registerRoutes(router *httprouter.Router) {

}