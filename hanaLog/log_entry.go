package hanalog

import "time"

type LogEntry struct {
	LogLevel  LogLevel  //日志等级
	Module    string    //日志所属模块
	Service   string    //日志所属服务
	TimeStamp time.Time //日志发生时间
	Message   string    //日志内容
	Field     []Field   //其他日志字段 -写法待定
}
