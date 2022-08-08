package netgear

import (
	zmq4 "github.com/pebbe/zmq4"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

var validSecurityMode = map[int]string{1: "IronHouse"}

type NetGear struct {
	ctx *zmq4.Context
	socket *zmq4.Socket
	address string
	port int
	protocol string
	pattern string
	receiveMod bool
	secureMode int
	authPublicKeysDir string
	authSecretKeysDir string
}

func (nGear *NetGear) Send(data []byte, hostname string)  {
	switch nGear.pattern {
	case "REQ":
		for {
			nGear.sendREQ(data, hostname)
			time.Sleep(time.Second)
		}
	}
}

func (nGear *NetGear) Recv()  {
	switch nGear.pattern {
	case "REP":
		for {
			nGear.receiveRep()
			time.Sleep(time.Second)
		}
	}
}

// NewNetGear ...
func NewNetGear (addr, protocol, pattern, pubKeys, secKeys string, port, secMode int, recv bool) *NetGear{
	nGear := new(NetGear)

	nGear.address = addr
	nGear.port = port
	nGear.protocol = protocol
	nGear.pattern = pattern
	nGear.receiveMod = recv
	nGear.secureMode = secMode
	nGear.authPublicKeysDir = pubKeys
	nGear.authSecretKeysDir = secKeys

	context, err := zmq4.NewContext()
	if err != nil {
		return nil
	}
	nGear.ctx = context

	switch nGear.pattern {
	case "REQ":
		nGear.socket, err = nGear.ctx.NewSocket(zmq4.REQ)
		if err != nil {
			log.Fatalf("%s", err)
		}
	case "REP":
		nGear.socket, err = nGear.ctx.NewSocket(zmq4.REP)
		if err != nil {
			log.Fatalf("%s", err)
		}
	default:
		log.Fatalf("Pattern %s no exist", nGear.pattern)
	}

	if nGear.receiveMod {
		if err := nGear.socket.Bind("tcp://" + nGear.address + ":" + strconv.Itoa(nGear.port)); err != nil {
			log.Fatalf("Can not bind on %s,\nErro: %s",nGear.address + ":" + strconv.Itoa(nGear.port), err)
		}
	} else {
		if err := nGear.socket.Connect("tcp://" + nGear.address + ":" + strconv.Itoa(nGear.port)); err != nil {
			log.Infof("Can not conection to tcp://\"%s,\nErro: %s",nGear.address + ":" + strconv.Itoa(nGear.port), err)
		}
	}


	return nGear
}
