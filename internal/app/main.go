package app

import (
	"WB-test-L0/internal/config"
	"WB-test-L0/internal/delivery/api/user"
	"WB-test-L0/internal/errors"
	"WB-test-L0/internal/repository/cache"
	"WB-test-L0/internal/repository/database"
	"WB-test-L0/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	natsPath = "nats://"
	colon    = ":"
)

func Run() error {
	log.Println("init config")
	cfg := config.NewConfig()

	log.Println("nats-streaming connect")
	sConn, err := StanConnect("sub", cfg.NatsStreaming.Host, cfg.NatsStreaming.Port)
	if err != nil {
		return errors.StanConnectError
	}
	pConn, err := StanConnect("pub", cfg.NatsStreaming.Host, cfg.NatsStreaming.Port)
	if err != nil {
		return errors.StanConnectError
	}

	log.Println("init gin Engine")
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	server := &http.Server{
		Addr:           cfg.App.Host + ":" + cfg.App.Port,
		Handler:        router,
		ReadTimeout:    time.Second * 15,
		WriteTimeout:   time.Second * 15,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("init cache")
	c := cache.NewCache()

	log.Println("init repository")
	d := database.NewRepository(cfg.Database.DB)

	log.Println("init service")
	s := service.NewUserService(c, d, pConn)

	err = s.GetAllEntity()
	if err != nil {
		log.Println()
	}

	//start subscribe
	ch := make(chan struct{})
	go stanSub(ch, sConn, s, cfg.NatsStreaming.Time)

	log.Println("init handlers")
	h := user.NewHandler(s)

	log.Println("register handlers")
	h.Register(router)

	go gracefulShutdown(ch, []os.Signal{syscall.SIGABRT, syscall.SIGQUIT, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM}, server, sConn, cfg.Database.DB)

	log.Println("start serve")
	return server.ListenAndServe()
}

func StanConnect(clientID, natsURL, natsPort string) (stan.Conn, error) {
	opt := stan.NatsURL(natsPath + natsURL + colon + natsPort)
	stanC, err := stan.Connect("test-cluster", clientID, opt)
	if err != nil {
		return nil, err
	}

	log.Println("nats-streaming connect successfully")
	return stanC, nil
}

func stanSub(ch chan struct{}, conn stan.Conn, s service.Service, duration time.Duration) {
	var stanOpt stan.SubscriptionOption
	stanOpt = stan.StartAtTimeDelta(duration)
	sub, err := conn.QueueSubscribe("wb", "wb", func(m *stan.Msg) {
		log.Println("GET FROM NATS ", string(m.Data))
		s.SetEntity(m.Data)
	}, stanOpt, stan.DurableName("wb"))
	if err != nil {
		log.Fatal(err.Error())
	}

	defer func() {
		sub.Unsubscribe()
		sub.Close()
	}()

	<-ch
}

func gracefulShutdown(ch chan struct{}, signals []os.Signal, closeItems ...io.Closer) {
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, signals...)
	sig := <-sign
	log.Printf("Caught signal %s. Shutting down...", sig)
	ch <- struct{}{}
	for _, closer := range closeItems {
		err := closer.Close()
		if err != nil {
			fmt.Printf("failed to close %v: %v", closer, err)
		}
	}
}
