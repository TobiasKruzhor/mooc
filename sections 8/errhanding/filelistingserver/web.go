package main

import (
	"log"
	"net/http"
	"os"

	"imooc.com/TobiasKruzhor/sectionEight/errhanding/filelistingserver/filelisting"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// panic
		defer func() {
			r := recover()
			log.Printf("Panic: %v", r)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}()
		err := handler(writer, request)
		if err != nil {
			log.Printf("Errot handling handling request: %s", err.Error())
			// user error
			if userErr, ok := err.(userError); ok {
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}
			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
