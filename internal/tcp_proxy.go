/**
 * @Author sgfoot.com
 * @date 2020/7/8 11:52
 * @Project_name yezihack
 */
package internal

import (
	"context"
	"io"
	"log"
	"net"
	"syscall"
	"time"
)

var (
	maxDelayTime = time.Second * 1 // 设置最大延时.
)

type TcpProxy struct {
	timeout    time.Duration // 连接超时
	listenAddr string        // 本地的 host:port 开启一个监听 ip:port 对外提供服务
	remoteAddr string        // 远程的 host:port 监听某个已有的服务
	ctx        context.Context
}

func NewTcpProxy(ctx context.Context, localAddr, remoteAddr string) *TcpProxy {
	tcp := &TcpProxy{
		listenAddr: localAddr,
		remoteAddr: remoteAddr,
		ctx:        ctx,
		timeout:    3 * time.Second,
	}
	log.Printf("listen:%s, remote:%s\n", localAddr, remoteAddr)
	go func() {
		if err := tcp.start(); err != nil {
			log.Printf("tcp.start, err:%v\n", err)
		}
	}()
	return tcp
}

// 设置超时
func (t *TcpProxy) SetTimeout(timeout time.Duration) {
	t.timeout = timeout
}

// 开始打开一个监听服务并将流量转发
func (t *TcpProxy) start() (err error) {
	// 开启一个监听端口, 对外提供服务
	ln, err := net.Listen("tcp", t.listenAddr)
	if err != nil {
		return
	}
	// 监听上下文关闭操作.并释放资源
	go func() {
		select {
		case <-t.ctx.Done():
			ln.Close()
			log.Printf("tcp server is close, listen addr:%s", t.listenAddr)
		}
	}()

	var tempDelay time.Duration
	for {
		// Accept等待并将下一个连接返回到侦听器。
		conn, err := ln.Accept()
		if err != nil {
			/*如果错误是暂时的,那么sleep一定时间在提供服务,否则就直接 return退出程序*/
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if tempDelay > maxDelayTime {
					tempDelay = maxDelayTime
				}
				time.Sleep(tempDelay)
				continue
			}
			return err
		}
		go t.handle(conn)
	}
}

// 拨号需要监听的 ip:port
func (t *TcpProxy) handle(conn net.Conn) {
	// 向已有的服务进行拨号连接.
	remote, err := net.DialTimeout("tcp", t.remoteAddr, t.timeout)
	if err != nil {
		if ne, ok := err.(*net.OpError); ok &&
			(ne.Err == syscall.ENFILE || ne.Err == syscall.EMFILE) {
			log.Printf("remote:%s, too many open files,err:%v", t.remoteAddr, err)
		} else {
			log.Printf("remote:%s, connection is fail, err:%v", t.remoteAddr, err)
		}
		return
	}
	// 进行双通道流量转发
	go t.pipeCopy(conn, remote)
	t.pipeCopy(remote, conn)
}

// src读取数据写入dst
func (t *TcpProxy) pipeCopy(src, dst net.Conn) {
	buf := bufPool.Get()
	defer func() {
		dst.Close()
		bufPool.Put(buf)
	}()
	var (
		err error
		n   int
	)
	for {
		n, err = src.Read(buf)
		if err != nil || n == 0 {
			// Always "use of closed network connection", but no easy way to
			// identify this specific error. So just leave the error along for now.
			// More info here: https://code.google.com/p/go/issues/detail?id=4373
			if err != io.EOF {
				log.Printf("err != nil  n == 0, err:%v", err)
			}
			break
		}
		if n > 0 {
			_, err := dst.Write(buf[:n])
			if err != nil {
				log.Printf("src<%s>->dst<%s> write err:%v\n", src.LocalAddr().String(), dst.LocalAddr().String(), err)
				break
			}
			if __conf.Debug {
				log.Printf("src<%s> -> dst<%s> write byte length:%d", src.LocalAddr().String(), dst.LocalAddr().String(), n)
			}
		}
	}
}
