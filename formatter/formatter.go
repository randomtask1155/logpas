/*
	Interface for log parsing format
 */

package formatter


var (
	Colors = map[string]string{
		"yellow": "\033[33m",
		"red": "\033[31m",
		"blue": "\033[34m",
		"light-gray": "\033[37m",
		"light-red": "\033[91m",
		"green": "\033[92m",
		"magenta": "\033[35m",
		"light-magenta": "\033[95m",
		"white": "\033[97m",
		"light-blue": "\033[94m",
		"light-yellow": "\033[93m",
		"RESET": "\033[0m"}
)

type Formatter  interface {
	Reformat([]byte) ([]byte, error)
	Read() error
	Write([]byte) error
}


func ParseLog(f Formatter) error {
	return f.Read()
}