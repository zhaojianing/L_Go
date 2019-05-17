package mock

type Reteriever struct {
	Contents string
}

func (r Reteriever) Get(url string) string  {
	return r.Contents
}
