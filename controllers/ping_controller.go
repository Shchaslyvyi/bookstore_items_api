package controllers

import (
	"net/http"
)

const (
	pong = "pong"
)

//
var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type pingController struct{}

// Ping func respondes with the pong
func (c *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}
