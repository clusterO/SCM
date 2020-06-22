package main

import (
	auth "auth/pkg/service"
	db "db/pkg/service"
	"log"
	"net/http"

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

	// connect to db

	flag.Parse()
	ctx := context.Background()

	//pipe to an error channel
	errs := make(chan error)

	go func ()  {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	*/

	/* Transport logging
	logger := log.NewLogfmtLogger(os.Stderr)

	var saveuser endpoint.Endpoint
	saveuser = MakeSaveUserEndpoint(dbs)
	saveuser = loggingMiddleware(log.With(logger, "method", "uppercase"))(uppercase)

	// getbyid
	// getbyusername
	*/

	/* Application logging */


	dbs := db.DBService{}
	oths := auth.Auth{}

	AuthenticateHandler := httptransport.NewServer(
		auth.MakeAuthenticateEndpoint(oths),
		auth.DecodeAuthenticateRequest,
		auth.EncodeResponse,
		httptransport.ServerErrorEncoder(auth.ErrorEncoder),
	)

	AuthorizeHandler := httptransport.NewServer(
		auth.MakeAuthorizeEndpoint(oths),
		auth.DecodeAuthorizeRequest,
		auth.EncodeResponse,
		httptransport.ServerErrorEncoder(auth.ErrorEncoder),
	)

	ValidateTokenHandler := httptransport.NewServer(
		auth.MakeValidateTokenEndpoint(oths),
		auth.DecodeValidateTokenRequest,
		auth.EncodeResponse,
		httptransport.ServerErrorEncoder(auth.ErrorEncoder),
	)

	EncryptionHandler := httptransport.NewServer(
		auth.MakeEncryptionEndpoint(oths),
		auth.DecodeEncryptionRequest,
		auth.EncodeResponse,
		httptransport.ServerErrorEncoder(auth.ErrorEncoder),
	)

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

	http.Handle("/authenticate", AuthenticateHandler)
	http.Handle("/authorize", AuthorizeHandler)
	http.Handle("/validate_token", ValidateTokenHandler)
	http.Handle("/encryption", EncryptionHandler)

	print("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
func loggingMiddleware(logger log.Logger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			logger.Log("msg", "calling endpoint")
			defer logger.Log("msg", "called endpoint")
			return next(ctx, request)
		}
	}
}
*/

/* Application instrumentation

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           DbService
}

func (mw instrumentingMiddleware) SaveUser(user *User) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "saveuser", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.DbService(s)
	return
}

// GetUserByID
// GetUserByUsername

*/

/* And wire it into our service.

import (
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/metrics"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "db_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "db_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "db_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	var svc DbService
	svc = DBService{}
	svc = loggingMiddleware{logger, svc}
	svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	SaveUserHandler := httptransport.NewServer(
		db.MakeSaveUserEndpoint(dbs),
		db.DecodeSaveUserRequest,
		db.EncodeResponse,
	)

	// GetById handler
	// GetByUsername handler

	http.Handle("/save", SaveUserHandler)
	http.Handle("/count", countHandler)
	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}

*/