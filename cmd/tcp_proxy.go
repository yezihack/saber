/**
 * @date 2020/7/9 16:35
 * @Project_name yezihack
 */
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/yezihack/saber/internal"
)

var (
	listenAddr, remoteAddr string
	TcpProxy               = &cobra.Command{
		Use:     "proxy",
		Example: "saber proxy --listen localhost:6379 --remote localhost:7379",
		Short:   "Tcp proxy",
		Long:    "Tcp proxy is traffic forwarding",
		Args: func(cmd *cobra.Command, args []string) error {
			if listenAddr == "" {
				return fmt.Errorf("listen address is nil")
			}
			if remoteAddr == "" {
				return fmt.Errorf("remote address is nil")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithCancel(context.Background()) // 创建一个上下文, 控制资源释放
			internal.NewTcpProxy(ctx, listenAddr, remoteAddr)       // 启动 tcp proxy
			waitSignal()                                            // 等待终端关闭信号
			cancel()                                                // 关闭上下文
			log.Println("tcp proxy is stop")
		},
	}
)

func init() {
	TcpProxy.PersistentFlags().StringVarP(&listenAddr,
		"listen", "l", "", "input your to the outside to server host:port 输入对外提供的服务地址")
	TcpProxy.PersistentFlags().StringVarP(&remoteAddr,
		"remote", "r", "", "input your want to listen to server host:port 输入需要监听的服务地址")
}

// 阻塞，只有执行信号才执行
func waitSignal() {
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-osSignals
}
