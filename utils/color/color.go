package color

import "fmt"

const (
	TEXT_BLACK = iota + 30
	TEXT_RED
	TEXT_GREEN
	TEXT_YELLOW
	TEXT_BLUE
	TEXT_PURPLE
	TEXT_CYAN
	TEXT_WHITE
)

const (
	BG_BLACK = iota + 40
	BG_RED
	BG_GREEN
	BG_YELLOW
	BG_BLUE
	BG_PURPLE
	BG_CYAN
	BG_WHITE
)

func PrintWith(s interface{}, color int) string {
	return fmt.Sprintf("\033[%dm%s\033[0m", color, s)
}
