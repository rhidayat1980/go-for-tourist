package logger

import "fmt"

// Info functions
func Info(s string) {
	fmt.Println("[INFO]", s)
}

func init() {
	fmt.Println("logger.init() - 1 from info.go")
}
