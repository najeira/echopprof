package echopprof

import (
	"github.com/labstack/echo"
	"net/http"
	"sync"
)

type customEchoHandler struct {
	httpHandler http.Handler

	wrappedHandleFunc echo.HandlerFunc
	once              sync.Once
}

func (ceh *customEchoHandler) Handle(c echo.Context) error {
	ceh.once.Do(func() {
		ceh.wrappedHandleFunc = echo.WrapHandler(ceh.httpHandler)
	})
	return ceh.wrappedHandleFunc(c)
}

func fromHTTPHandler(httpHandler http.Handler) *customEchoHandler {
	return &customEchoHandler{httpHandler: httpHandler}
}

func fromHandlerFunc(serveHTTP func(w http.ResponseWriter, r *http.Request)) *customEchoHandler {
	return &customEchoHandler{httpHandler: &customHTTPHandler{serveHTTP: serveHTTP}}
}
