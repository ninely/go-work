package main

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func indexHandler(resp http.ResponseWriter, req *http.Request) {
	_, _ = resp.Write([]byte("hello world"))
}

type myServer struct {
	ctx    context.Context
	cancel func()
	*http.Server
}

func NewMyServer() *myServer {
	http.HandleFunc("/", indexHandler)
	srv := &http.Server{Addr: ":80"}
	ctx, cancel := context.WithCancel(context.Background())
	return &myServer{
		ctx:    ctx,
		cancel: cancel,
		Server: srv,
	}
}

func (s *myServer) Start() error {
	err := s.Server.ListenAndServe()

	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *myServer) Stop() error {
	defer s.cancel()
	return s.Server.Shutdown(context.TODO())
}

func main() {
	mySrv := NewMyServer()
	eg, ctx := errgroup.WithContext(mySrv.ctx)

	eg.Go(func() error {
		return mySrv.Start()
	})

	c := make(chan os.Signal, 1)
	sigs := []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT}
	signal.Notify(c, sigs...)

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				if err := mySrv.Stop(); err != nil {
					log.Println("stop server failed:", err)
					return err
				}
			}
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		log.Printf("server end with err: %v", err)
	}
	log.Println("server stop")
}
