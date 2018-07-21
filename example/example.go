package main

import (
	"fmt"
	"net"
)

func main() {
	// factory is the function that create connection
	factory = func() (net.Conn, error) { return net.Dial(network, address) }

	// poolConfig is the config of gpool
	poolConfig = &PoolConfig{
		InitCap: 5,
		MaxCap:  30,
		Factory: factory,
	}

	// create a new conn pool
	p, err := NewGPool(poolConfig)
	if err != nil {
		fmt.Println("new pool error is", err)
	}

	// release pool
	defer p.Close()

	// get a new connection from pool
	conn, err := p.Get()
	if err != nil {
		t.Errorf("Get error: %s", err)
	}

	_, ok := conn.(*GConn)
	if !ok {
		t.Errorf("Conn is not of type GConn")
	}

	// return connection to pool
	conn.Close()

	// return len of the pool
	fmt.Println("len=", p.Len())

}