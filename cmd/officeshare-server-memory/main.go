package main

import "github.com/Tiffinger-Thiel-GmbH/projectshare-api/api"

func main() {
	s := api.Server{}
	s.With(api.InMemory())
	s.Init()
	err := s.Serve()
	panic(err)
}
