package taskrunner

import (
	"log"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	d := func(dc dataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("Dispatcher sent: %v\n", i)
		}
		return nil
	}

	e := func(dc dataChan) error {
	forLoop:
		for {
			select {
			case d := <-dc:
				log.Printf("Executor received: %v\n", d)
			default:
				break forLoop
			}
		}
		return nil
	}
	runner := NewRunner(30, false, d, e)
	go runner.StartAll()
	time.Sleep(3 * time.Second)
}
