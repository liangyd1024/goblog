//AppName: goblog
//Version: V1.0.0
//User: marco
//Date: 2019-01-31

package crypt

import (
	"fmt"
	"testing"
)

func TestGetMd5(t *testing.T) {
	fmt.Println(GetMd5("goblog"))
}
