package reals

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Reteriever struct {
	UserAgent string
	TimeOut time.Duration
}

func (r *Reteriever) Get(url string) string  {
	resp,err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result , err := httputil.DumpResponse(resp,true)
	// close
	resp.Body.Close()
	if err != nil {
		panic(err)
	}
	return string(result)
}
