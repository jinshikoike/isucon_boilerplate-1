package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
)

type Logger struct {
	Time    *time.Time
	File    *os.File
	Comment string
}

var file = Initialize()

func MeasureStart(comment string) *Field {
	now := time.Now()

	logger := Logger{
		Time:    &now,
		File:    file,
		Comment: comment,
	}
	return &logger
}

func MeasureEnd(logger *Logger) {
	delta := getDeltaTime(logger.Time)
	funcName := getFuncName()

	log := fmt.Sprintf("%v: %v: %v\n", logger.Comment, funcName, delta)

	writer := bufio.NewWriter(logger.File)
	_, err := writer.WriteString(log)

	if err != nil {
		fmt.Println(err)
	}

	writer.Flush()
}

func Initialize() *os.File {
	file, err := os.Create(fmt.Sprintf("/tmp/%v", time.Now()))

	if err != nil {
		fmt.Println(err)
	}

	return file
}

func getFuncName() string {
	pc, _, _, _ := runtime.Caller(2)
	fn := runtime.FuncForPC(pc)
	return fn.Name()
}

func getDeltaTime(t *time.Time) time.Duration {
	return time.Now().Sub(*t)
}

func chDir(path string) {
	_, err := os.Stat(path)

	if err != nil {
		err := os.Mkdir(path, 0777)

		if err != nil {
			fmt.Println(err)
		}
	}

	os.Chdir(path)
}
