# NetGear

NetGear implements a wrapper around pebbe/zmq4 golang library that contains golang bindings for ZeroMQ.

***

#### NetGear as of now seamlessly supports three ZeroMQ messaging patterns:

- zmq4.REQ/zmq4.REP (ZMQ Request/Reply Pattern)

whereas the supported protocol only tcp.

## Sample usage:
    package main

    import (
        netgear "github.com/chorshik/NetGear"
        "sync"
    )

    func main() {
        var wg sync.WaitGroup
    
        srv := netgear.NewNetGear("127.0.0.1", "TCP", "REP", "", "", 5551, 0, true)
        client := netgear.NewNetGear("127.0.0.1", "TCP", "REQ", "", "", 5551, 0, false)
    
        wg.Add(2)
    
        go client.Send([]byte("data"), "host1")
        go srv.Recv()
    
        wg.Wait()

    }
