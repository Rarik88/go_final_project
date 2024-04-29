package go_final_project

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) (err error) {
	op := "status.Server.Run"

	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	err = s.httpServer.ListenAndServe()
	logrus.Println(op, "start http server on port", port)
	return
}
