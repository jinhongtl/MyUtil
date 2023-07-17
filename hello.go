package MyUtil

import "fmt"

func HelloWorld() string {
	return "world"
}

func HelloV2(format string, s string) string {
	return fmt.Sprintf(format, s)
}

func A() {}
