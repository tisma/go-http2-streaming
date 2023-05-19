package http2

import (
	"io"
	"log"
	"net"
	"net/http"

	httpRouter "github.com/julienschmidt/httprouter"
)

type Server struct {
	router *httpRouter.Router
}

func (s *Server) Initialize() error {
	s.router = httpRouter.New()
	s.router.POST("/", s.handler)

	server := &http.Server{
		Handler: s.router,
	}

	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		return err
	}

	log.Println("HTTP server is listening...")
	return server.ServeTLS(listener, "./http2/certs/server.crt", "./http2/certs/server.key")
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request, _ httpRouter.Params) {
	if r.ProtoMajor != 2 {
		log.Println("Not a HTTP/2 request, rejected!")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	buf := make([]byte, 4*1024)

	for {
		n, err := r.Body.Read(buf)
		if n > 0 {
			w.Write(buf[:n])
		}

		if err != nil {
			if err == io.EOF {
				w.Header().Set("Status", "200 OK")
				r.Body.Close()
			}
			break
		}
	}
}
