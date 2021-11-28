package main

import (
	"log"
	"notification/service"
)

func main() {
	s := service.NewService()
	if err := s.Run(); err != nil {
		log.Panic(err)
	}
}
