//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2019-04-01

package component

import (
	"fmt"
	"testing"
	"time"
)

func TestGoRoutine(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("TestGoRoutine err:[%+v]", err)
		}
	}()

	GoRoutine(func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("GoRoutine err:[%+v]", err)
			}
		}()

		a := 0
		a = 1 / a
	})
	time.Sleep(time.Second * 2)
}
