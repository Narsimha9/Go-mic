package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	// fieldKeys := []string{"method", "error"}

	var svc OrderService
	svc = orderService{}
	svc = loggingMiddleware{logger, svc}

	makeCreateHandler := httptransport.NewServer(
		makeCreateEndpoint(svc),
		decodeCreateRequest,
		encodeResponse,
	)

	http.Handle("/createOrder", makeCreateHandler)

	logger.Log("msg", "HTTP", "addr", ":8083")
	logger.Log("err", http.ListenAndServe(":8083", nil))
}
