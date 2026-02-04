package log

import "log"

var (
	std = log.Default()
)

func Info(msg string, fields ...Field) {
	std.Println(format("[INFO]", msg, fields...))
}

func Error(msg string, fields ...Field) {
	std.Println(format("[ERROR]", msg, fields...))
}

func TitleInfo(title string, msg string, fields ...Field) {
	std.Println("================[" + title + "]================")
	std.Println(format("[INFO]", msg, fields...))
	std.Println("================[" + title + "]================")
}
