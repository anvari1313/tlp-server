package log

import "fmt"

func Error(error string) {
	fmt.Errorf(error)
}
