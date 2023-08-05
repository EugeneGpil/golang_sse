package sseServer

import "github.com/r3labs/sse/v2"

var sseServer *sse.Server

func Get() *sse.Server {
	if sseServer != nil {
		return sseServer
	}

	sseServer = sse.New()

	return sseServer
}
