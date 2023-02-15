package builder

import (
	"context"
	"io"
	"net/http"
)

type builder struct{
	headers map[string][]string
	url 	string
	method string
	body io.Reader
	close bool
	ctx context.Context
}

func (b *builder) AddHeader(name, value string) HttpBuilder {
	values, ok := b.headers[name]

	if !ok{
		values = make([]string,0)
	}
	values = append(values,value)
	return b
}

func (b *builder) Body(r io.Reader) HttpBuilder {
	b.body = r
	return b
}

func (b *builder) Method(method string) HttpBuilder {
	b.method = method
	return b
}

func (b *builder) Close(close bool) HttpBuilder {
	b.close = close
	return b
}

func (b *builder) Build() (*http.Request, error) {
	req, err := http.NewRequestWithContext(b.ctx, b.url, b.method, b.body)
	if err != nil{
		return nil, err
	}

	for k, v := range b.headers{
		for _, val := range v{
			req.Header.Add(k,val)
		}
	}

	req.Close = b.close
	return req, nil
}

type HttpBuilder interface{
	AddHeader(name, value string) HttpBuilder
	Body(r io.Reader) HttpBuilder
	Method(method string) HttpBuilder
	Close(close bool) HttpBuilder
	Build()(*http.Request, error)
}
func NewBuilder(url string)HttpBuilder{
	return &builder{
		headers: map[string][]string{},
		url:    url,
		method:  http.MethodGet,
		body:    nil,
		close:   false,
		ctx:     nil,
	}
}
