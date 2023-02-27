package main

import (
	"github.com/gin-gonic/gin"
	_ "net/http/pprof"
)

func main() {
	//go kClient.KClient()

	r := gin.Default()

	initRouter(r)
	//go testPprof()

	err := r.Run()
	if err != nil {
		panic("run failed.")
	}
}

//func testPprof() {
//	err := http.ListenAndServe(":8888", nil)
//	if err != nil {
//		return
//	}
//}
