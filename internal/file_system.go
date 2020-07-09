/**
 * @Author sgfoot.com
 * @date 2020/7/7 19:40
 * @Project_name yezihack
 */
package internal

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func FileSystem(dirname string, port int) error {
	handle := http.FileServer(http.Dir(dirname))
	s := &http.Server{
		Handler:        handle,
		Addr:           fmt.Sprintf(":%d", port),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("%d http://0.0.0.0:%d\n", os.Getpid(), port)
	return s.ListenAndServe()
}
