package pipeline

import "net/http"

type Pipeline interface {
	handle(request *http.Request, response *http.Response, doNext func(r *http.Request))

	enable() bool
}
