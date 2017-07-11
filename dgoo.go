package main

import (
	"fmt"
	"net/http"

	"github.com/kardianos/service"
	"log"
)

var logger service.Logger

type program struct{}
func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}
func (p *program) run() {
	http.HandleFunc(`/`, indexHandler)
	http.ListenAndServe(":3000", nil)
}
func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "GoServiceExampleSimple",
		DisplayName: "Go Service Example",
		Description: "This is an example Go service.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "New test Message")
}
