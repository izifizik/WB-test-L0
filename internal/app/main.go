package app

import (
	"github.com/nats-io/stan.go"
)

func Run() error {
	opt := stan.Option(nil)
	opts := &stan.Options{NatsURL: ""}
	err := opt(opts)
	if err != nil {
		return err
	}
	stanC, err := stan.Connect("testCluster", "testID", opt)
	if err != nil {
		return err
	}
	defer stanC.Close()
	return nil
}
