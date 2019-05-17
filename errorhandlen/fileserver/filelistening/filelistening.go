package filelistening

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter,
request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		//panic(err)
		//http.Error(writer,err.Error(),http.StatusInternalServerError)
		return err
	}
	defer file.Close()
	all, errRead := ioutil.ReadAll(file)
	if errRead != nil {
		//panic(errRead)
		return err
	}

	writer.Write(all)
	return nil
}