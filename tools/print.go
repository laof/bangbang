package tools

import "fmt"

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
)

func PrintError(msg string) {
	fmt.Println(Red + msg + Reset)
}

func PrintSuccess(msg string) {
	fmt.Println(Green + msg + Reset)

}
