package http

import (
	articleRouter "ximlr/go-lang-web/http/article"

	"github.com/julienschmidt/httprouter"
)

func HttpInit() *httprouter.Router {
	http := httprouter.New()

	articleRouter.Init(http)

	return http
}
