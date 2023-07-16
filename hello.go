package MyUtil

import "fmt"

func Hello() string {
	return "Hello,World!"
}

func HelloWorld() string {
	return "world"
}

func HelloV2(format string, s string) string {
	return fmt.Sprintf(format, s)
}
