package data

import (
	"context"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	recordV1 "github.com/ninely/go-work/final/api/record/v1"
	searchV1 "github.com/ninely/go-work/final/api/search/v1"
	solveV1 "github.com/ninely/go-work/final/api/solve/v1"
	"github.com/ninely/go-work/final/app/bff/homework/internal/conf"
	traceSDK "go.opentelemetry.io/otel/sdk/trace"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewRecordRepo,
	NewSolveRepo,
	NewSearchRepo,
	NewDiscovery,
	NewRegistrar,
	NewSearchClient,
	NewSolveClient,
	NewRecordClient,
)

type Data struct {
	log          *log.Helper
	searchClient searchV1.SearchClient
	solveClient  solveV1.SolveClient
	recordClient recordV1.RecordClient
}

// NewData .
func NewData(
	conf *conf.Data,
	logger log.Logger,
	searchClient searchV1.SearchClient,
	solveClient solveV1.SolveClient,
	recordClient recordV1.RecordClient,
) (*Data, error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	return &Data{log: l, searchClient: searchClient, solveClient: solveClient, recordClient: recordClient}, nil
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewSearchClient(ac *conf.Auth, r registry.Discovery, tp *traceSDK.TracerProvider) searchV1.SearchClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///final.search.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
			jwt.Client(func(token *jwt2.Token) (interface{}, error) {
				return []byte(ac.ServiceKey), nil
			}, jwt.WithSigningMethod(jwt2.SigningMethodHS256)),
		),
	)
	if err != nil {
		panic(err)
	}
	c := searchV1.NewSearchClient(conn)
	return c
}

func NewSolveClient(r registry.Discovery, tp *traceSDK.TracerProvider) solveV1.SolveClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///final.solve.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return solveV1.NewSolveClient(conn)
}

func NewRecordClient(r registry.Discovery, tp *traceSDK.TracerProvider) recordV1.RecordClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///final.record.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			tracing.Client(tracing.WithTracerProvider(tp)),
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return recordV1.NewRecordClient(conn)
}
