package coresharedcontroller

import (
	"net/http"
)

func NotImplementedHandler(w http.ResponseWriter, _ *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
