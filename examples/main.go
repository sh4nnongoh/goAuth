package main

// https://jacobmartins.com/2016/02/29/getting-started-with-oauth2-in-go/

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	//kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	//stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const htmlIndex = `
	<html><body>
	<a href="/GoogleLogin">Log in with Google</a>
	</body></html>
`

func main() {

	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	flag.Parse()
	logger := log.NewLogfmtLogger(os.Stderr)

	r := mux.NewRouter()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		//httptransport.ServerErrorEncoder(encodeError),
	}

	r.Methods("POST").Path("/generateReport").Handler(httptransport.NewServer(
		//githubstats2.MakeGenerateReportEndpoint(svc),
		//githubstats2.DecodeGenerateReportRequest,
		//githubstats2.EncodeResponse,
		options...,
	))

	r.Methods("GET").Path("/metrics").Handler(promhttp.Handler())

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		server := &http.Server{
			Addr:    *httpAddr,
			Handler: r,
		}
		errs <- server.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)
}
