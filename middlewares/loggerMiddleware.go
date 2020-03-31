package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logrus.Info(fmt.Sprintf("-REQ INFO- %s %s", r.URL.String(), r.Method))
		begin := time.Now().UnixNano()
		next(w, r, ps)
		end := time.Now().UnixNano()
		logrus.Info(fmt.Sprintf("-RES INFO- PROCESSED IN %dns", (end - begin)))
	}
}
