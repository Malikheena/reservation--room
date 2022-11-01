package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/malikheena/heena2-main/internal/config"
)

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a

}

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.Errorlog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

}
