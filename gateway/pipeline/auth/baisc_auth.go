package auth

import "net/http"

type BasicAuth struct {
	username string
	password string
}

func (b *BasicAuth) handle(request *http.Request, response *http.Response, doNext func(*http.Request, *http.Response)) {
	username, password, ok := request.BasicAuth()

	if ok && username == b.username && password == b.password {
		doNext(request, response)
		return
	}
	response.StatusCode = 403
	//response.Body
}

func (b *BasicAuth) enable() bool {
	panic("implement me")
}
