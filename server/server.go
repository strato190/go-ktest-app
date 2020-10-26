package server

import (
	"syscall"

	"github.com/fvbock/endless"
)

//Init the gin server
func Init() {
	//logging.Setup()
	r := setupRouter()

	srv := endless.NewServer(":8080", r)
	srv.SignalHooks[endless.POST_SIGNAL][syscall.SIGUSR1] = append(
		srv.SignalHooks[endless.POST_SIGNAL][syscall.SIGUSR1],
		preStopSleepHook)
	srv.ListenAndServe()

}
