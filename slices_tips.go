import "log"
import "strings"

alice = append(alice, bob...)


func Log(format string, v ...interface{}) {
	format = strings.TrimSuffix(format, "\n") + "\n"
	log.Printf(format, v...)
}
