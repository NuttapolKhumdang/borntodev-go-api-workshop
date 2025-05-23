package main

import (
	"fmt"
	"strings"
	"time"
)

type Logger struct {
	DEBUG func(v ...any)
	TRACE func(v ...any)
	INFO  func(v ...any)
	WARN  func(v ...any)
	ERROR func(v ...any)
}

// func CreateLogger(namespace string) func(v ...any) {
// 	fmt.Printf("[%v] logger: start loggin '%s'\n", time.Now(), namespace)

// 	return func(v ...any) {
// 		var r []string

// 		for _, i := range v {
// 			r = append(r, fmt.Sprint(i))
// 		}

// 		fmt.Printf("[%v] %s: %s\n", time.Now(), namespace, strings.Join(r, " "))
// 	}
// }

func print(namespace string, level string, v []any) {
	var r []string

	for _, i := range v {
		r = append(r, fmt.Sprint(i))
	}

	fmt.Printf("[%v]\n[%s] %s: %s\n", time.Now(), level, namespace, strings.Join(r, " "))
}

func CreateLogger(namespace string) *Logger {
	print("Logger", "LOG", []any{"start logger", namespace})

	return &Logger{
		DEBUG: func(v ...any) {
			print(namespace, "DEBUG", v)
		},
		TRACE: func(v ...any) {
			print(namespace, "TRACE", v)
		},
		INFO: func(v ...any) {
			print(namespace, "INFO", v)
		},
		WARN: func(v ...any) {
			print(namespace, "WARN", v)
		},
		ERROR: func(v ...any) {
			print(namespace, "ERROR", v)
		},
	}
}
