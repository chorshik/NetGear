package netgear

import (
	log "github.com/sirupsen/logrus"
	"time"
)

func (nGear *NetGear) receiveRep()  []byte{
	data, err := nGear.socket.RecvBytes(0)
	if err != nil {
		log.Infof("Can not receive data from client\nError: %v", err)
		return nil
	}
	log.Info("from server: ", string(data))
	time.Sleep(time.Microsecond)
	_, err = nGear.socket.Send("ok", 0)
	if err != nil {
		log.Infof("Can not send data to client\nError: %v", err)
		return nil
	}
	time.Sleep(time.Microsecond)

	return data
}