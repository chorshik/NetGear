package netgear

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func (nGear *NetGear) sendREQ(data []byte, host string)  string{
	_, err := nGear.socket.SendBytes(data, 0)
	if err != nil {
		log.Infof("Can not send data to server\nError: %v", err)
		return "nil"
	}
	time.Sleep(time.Microsecond)
	str, err := nGear.socket.Recv(0)
	if err != nil {
		log.Infof("Can not receive data from server\nError: %v", err)
		return ""
	}
	log.Info(str)
	time.Sleep(time.Microsecond)
	return str
}

func (nGear *NetGear) sendREP(data []byte , host string)  {
	nGear.socket.SendBytes(data, 0)
}
