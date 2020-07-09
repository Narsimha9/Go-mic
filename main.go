package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	// fieldKeys := []string{"method", "error"}
	db := GetMongoDB()
	// var svcs OrderService
	var svc OrderService
	{
		repository, err := NewRepo(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		svc = NewService(repository, logger)
	}
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

// var svc order.Service
// {
//    repository, err := cockroachdb.New(db, logger)
//    if err != nil {
// 	  level.Error(logger).Log("exit", err)
// 	  os.Exit(-1)
//    }
//    svc = ordersvc.NewService(repository, logger)
// }
