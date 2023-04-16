package helpers

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	counter atomic.Int32
	once    sync.Once
)

func CounterAdd() {
	once.Do(func() {
		go func() {
			for {
				timer := time.NewTimer(time.Second * 10)

				<-timer.C
				logrus.Print("за 10 секунд обработано: ", counter.Load())

				counter.Store(0)
			}
		}()
	})

	counter.Add(1)
}
