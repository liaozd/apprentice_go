package main

import "net/http"

type Server interface {
	Route(pattern string, handleFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

type Heaer map[string][]string
