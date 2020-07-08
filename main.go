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

	// countHandler := httptransport.NewServer(
	// 	makeCountEndpoint(svc),
	// 	decodeCountRequest,
	// 	encodeResponse,
	// )

	http.Handle("/createOrder", makeCreateHandler)
	// http.Handle("/count", countHandler)
	logger.Log("msg", "HTTP", "addr", ":8081")
	logger.Log("err", http.ListenAndServe(":8081", nil))
}
