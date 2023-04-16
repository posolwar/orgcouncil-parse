package helpers

import (
	"time"

	"github.com/sirupsen/logrus"
)

func Counter(name string, i *int) {
	time.Sleep(time.Second * 10)
	logrus.Printf("%s is count %d", name, *i)
}
