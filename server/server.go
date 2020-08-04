package server

import (
	"github.com/fvbock/endless"
)

//Init the gin server
func Init() {
	r := setupRouter()

	srv := endless.NewServer(":8080", r)

	srv.ListenAndServe()

}
