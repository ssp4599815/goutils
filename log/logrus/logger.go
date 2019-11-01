package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path"
)

func init() {
	// 设置日志格式为 json格式
	log.SetFormatter(&log.JSONFormatter{})

	// 设置日志的输出位置
	baseDir, _ := os.Getwd()
	logFile, err := os.Create(path.Join(baseDir, "log/logrus/my.log"))
	if err != nil {
		log.Fatal("failed create file, err:", err)
	}
	log.SetOutput(logFile)

	// 设置日志级别为warn以上
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.Errorf("haha")
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Warn("A group of walrus emerges from the ocean")
}
