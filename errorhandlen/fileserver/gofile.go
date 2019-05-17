package main

import (
	"github.com/gpmgo/gopm/modules/log"
	"google_go/errorhandlen/fileserver/filelistening"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer,request)
		log.Warn("error handling request %s",err.Error())
		if err != nil {
			code := http.StatusOK
			switch  {
			case os.IsNotExist(err):
				//http.Error(
				//	writer,
				//	http.StatusText(http.StatusNotFound),
				//	http.StatusNotFound)
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden

			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer,http.StatusText(code),code)
		}

	}
}

func main() {
	http.HandleFunc("/list/",errWrapper(filelistening.HandleFileList))
	errHttp := http.ListenAndServe(":8888",nil)
	if errHttp != nil {
		panic(errHttp)
	}
}

