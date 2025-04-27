package initializers

import (
	"log"
	"net/http"
	"strconv"
)

// register the prometheus metrics
func init() {

}

func InitializeServer(port int64) {
	portString := strconv.FormatInt(port, 10)
	log.Fatal(http.ListenAndServe("localhost:"+portString, Handler{}))
}
