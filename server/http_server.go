package server

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"home_api.totote05.ar/temperature"
)

type HttpServer struct {
	server *http.Server
}

func (s *HttpServer) RegisterTemperatureHandler(handler temperature.TemperatureService) {
	http.HandleFunc("/temperature/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(handler.HandleLastValue())
	})

	http.HandleFunc("/temperature/history", func(w http.ResponseWriter, r *http.Request) {
		values := []temperature.Record{}
		handler.HandleTemperatureHistory(func(c <-chan temperature.Record) {
			for value := range c {
				values = append(values, value)
			}
		})
		if data, err := json.Marshal(values); err != nil {
			log.Println("can't get temperature hidtory , reason:", err)
			http.Error(w, "", http.StatusInternalServerError)
		} else {
			w.Write(data)
		}
	})

	http.HandleFunc("/temperature/add", func(w http.ResponseWriter, r *http.Request) {
		if data, err := io.ReadAll(r.Body); err != nil {
			log.Println("can't register the temperature, reason:", err)
			http.Error(w, "Invalid payload", http.StatusBadRequest)
		} else if err := handler.HandleRegisterTemperature(data); err != nil {
			log.Println("can't register the temperature, reason:", err)
			http.Error(w, "Internar server error", http.StatusInternalServerError)
			log.Println("")
		} else {
			w.Write([]byte("ok!"))
		}
	})
}

func (s *HttpServer) Start() {
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
		<-sig
		s.server.Shutdown(context.Background())
	}()

	log.Println("Server started on", s.server.Addr)
	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("Server stopper incorrectly", err)
	} else {
		log.Println("Server stopped successfuly")
	}
}

func NewHttpServer(host string) *HttpServer {
	server := &http.Server{
		Addr: host,
	}

	return &HttpServer{server: server}
}

var _ Server = (*HttpServer)(nil)
