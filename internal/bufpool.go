/**
 * @Author sgfoot.com
 * @date 2020/7/8 14:23
 * @Project_name yezihack
 */
package internal

import "sync"

var bufPool = newBufferPool()

type bufferPool struct {
	sync.Pool
}

func newBufferPool() *bufferPool {
	return &bufferPool{
		Pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, 4096)
			},
		},
	}
}

func (p *bufferPool) Get() []byte {
	return p.Pool.Get().([]byte)
}

func (p *bufferPool) Put(b []byte) {
	p.Pool.Put(b)
}
