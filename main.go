package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gorilla/mux"

	"github.com/rxrav/loan_app/src/constant"
	"github.com/rxrav/loan_app/src/depInject"
	"github.com/rxrav/loan_app/src/handler"
	"github.com/rxrav/loan_app/src/middleware"
)

func main() {
	var wait time.Duration

	// initializing logger, basic config to set logging level to info
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// info logging example
	log.Info().Msg("Hello Gophers, welcome to the loan application service...")

	// creating a new router from gorilla mux package
	r := mux.NewRouter()
	// adding the middleware, telling router to use ErrorMiddleware to handle errors across the app
	r.Use(middleware.ErrorMiddleware)

	// using a wrapper handler function to reach actual handler function where we
	// inject the loan application required to process this route
	r.HandleFunc(constant.ApplyRoute, func(writer http.ResponseWriter, request *http.Request) {
		handler.ApplyLoanHandler(writer, request, depInject.CreateLoanApplicationService())
	}).Methods(http.MethodPost)

	r.HandleFunc(constant.GetAllLoansRoute, func(writer http.ResponseWriter, request *http.Request) {
		handler.GetAllLoansHandler(writer, request, depInject.CreateLoanApplicationService())
	}).Methods(http.MethodGet)

	r.HandleFunc(constant.GetLoanRoute, func(writer http.ResponseWriter, request *http.Request) {
		handler.GetLoanHandler(writer, request, depInject.CreateLoanApplicationService())
	}).Methods(http.MethodGet)

	// TODO: refactor following two lines as above and use in integrationtests tests with tc
	r.HandleFunc(constant.GetUserRoute, handler.GetUserHandler).Methods(http.MethodGet)
	r.HandleFunc(constant.CreateUserRoute, handler.CreateUserHandler).Methods(http.MethodPost)

	// if err := http.ListenAndServe(localAddr, r); err != nil {
	// 	log.Error().Msg(fmt.Sprintf("%v", err))
	// }

	server := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Error().Msg(fmt.Sprintf("%v", err))
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	err := server.Shutdown(ctx)
	if err != nil {
		return
	}

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Info().Msg("shutting down...")
	os.Exit(0)
}
