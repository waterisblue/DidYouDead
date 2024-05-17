package udpServer

import (
	"dyd/dao"
	"dyd/log"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func Start() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})
	if err != nil {
		log.Error.Println("UDP服务开启失败", err)
		return
	}
	defer listen.Close()
	for {
		var data [1024]byte
		n, _, err := listen.ReadFromUDP(data[:])
		if err != nil {
			log.Warning.Println("READ UDP FAILED:", err)
			continue
		}
		receiveData := string(data[:n])
		names := strings.Split(receiveData, ",")
		atoi, _ := strconv.Atoi(names[1])
		dao.InsertDetection(names[0], atoi)
	}
}

func RandomWrite() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 9090,
	})

	if err != nil {
		log.Warning.Println(err)
	}

	defer socket.Close()
	name := []string{"zjntest", "qhytest", "aaatest"}
	for {
		sendData := []byte(fmt.Sprintf("%s,%d", name[rand.Intn(len(name)-1)], rand.Intn(15)*10))
		_, err = socket.Write(sendData)
		if err != nil {
			log.Warning.Println(err)
		}
		time.Sleep(time.Second * 2)
	}

}
