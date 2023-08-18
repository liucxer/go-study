package main

import (
	"github.com/sirupsen/logrus"
	"time"
)

func main() {

	ch := make(chan *int, 4)
	go func() {
		for i:=1; i < 100; i++ {
			a := i
			ch <- &a
			time.Sleep(1)
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		close(ch)
	}()

	go func() {
		for {
			select {
			case item := <-ch :
				if item == nil {
					logrus.Infof("item == nil")
					continue
				}

				logrus.Info("item:%d", *item)
			}
			time.Sleep(time.Second)
		}
	}()


	for {
		time.Sleep(time.Second)
	}

}
