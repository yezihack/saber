/**
 * @Author sgfoot.com
 * @date 2020/7/7 19:40
 * @Project_name yezihack
 */
package internel

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func FileSystem(dirname string, port int) error {
	http.FileServer(http.Dir(dirname))
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
		// Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("%d http://0.0.0.0:%d\n", os.Getpid(), port)
	return s.ListenAndServe()
}
