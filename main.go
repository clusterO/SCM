package main

import (
	db "db/pkg/service"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	/* LOG
	var logger log.Logger {
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "db",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	*/

	// connect to db

	flag.Parse()
	// ctx := context.Background()

	// pipe to an error channel
	errs := make(chan error)

	go func ()  {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	dbs := db.DBService{}

	SaveUserHandler := httptransport.NewServer(
		db.MakeSaveUserEndpoint(dbs),
		db.DecodeSaveUserRequest,
		db.EncodeResponse,
	)

	GetUserByIDHandler := httptransport.NewServer(
		db.MakeGetUserByIDEndpoint(dbs),
		db.DecodeGetUserByIDRequest,
		db.EncodeResponse,
	)

	GetUserByUsernameHandler := httptransport.NewServer(
		db.MakeGetUserByUsernameEndpoint(dbs),
		db.DecodeGetUserByUsernameRequest,
		db.EncodeResponse,
	)

	http.Handle("/save", SaveUserHandler)
	http.Handle("/get_by_id", GetUserByIDHandler)
	http.Handle("/get_by_username", GetUserByUsernameHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	print("listening on port 8080")
}
