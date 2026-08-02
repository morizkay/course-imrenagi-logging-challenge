package functions

import (
	"context"
	"net/http"
)

func Handler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	res, err := Greeting(ctx, name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(res))
}
