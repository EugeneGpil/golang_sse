package globals

import (
	"github.com/r3labs/sse/v2"
)

var server = sse.New()

func GetServer() *sse.Server {
	return server
}
