package routes

import "net/http"

type Routes struct {
	ServeMux *http.ServeMux
}

func (r *Routes) Init() {
	basePath := "/api/v1"
	r.ServeMux.HandleFunc(basePath+"/ping", ping)
}

func ping(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(http.StatusText(http.StatusMethodNotAllowed)))
		return
	}

	response.Write([]byte("pong\n"))
}
