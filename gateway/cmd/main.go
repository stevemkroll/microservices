package main

import (
	"gateway/service"
	"log"
)

func main() {
	s := service.NewService()
	if err := s.Run(); err != nil {
		log.Panic(err)
	}
}
