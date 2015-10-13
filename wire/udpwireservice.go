package wire

import (
	"fmt"
	"github.com/nstehr/go-mansbridge/collections"
	"log"
	"net"
	"time"
)

type UdpWireService struct {
	port    int
	encoder WireEncoder
	address string
}

func (service *UdpWireService) GetAddress() string {
	return service.address
}

func getAddress(port int) string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fullAddress := fmt.Sprintf("%s:%d", ipnet.IP.String(), port)
				return fullAddress
			}
		}
	}
	return ""
}

func NewUdpWireService(port int, wireEncoder WireEncoder) *UdpWireService {
	return &UdpWireService{port: port, encoder: wireEncoder, address: getAddress(port)}
}

func (service *UdpWireService) SendNews(correspondent string, entries []collections.Entry) {
	correspondentAddr, err := net.ResolveUDPAddr("udp", correspondent)
	if err != nil {
		log.Println("Error resolving: " + correspondent)
		return
	}

	conn, err := net.DialUDP("udp", nil, correspondentAddr)

	defer conn.Close()

	msg := WireMessage{Entries: entries, CurrentTime: time.Now(), Source: service.address}
	data, err := service.encoder.Encode(msg)
	if err != nil {
		log.Println("Error encoding data")
		return
	}

	_, sendErr := conn.Write(data)
	if sendErr != nil {
		log.Println(sendErr)
	}

}

func (service *UdpWireService) GetNews() <-chan WireMessage {

	out := make(chan WireMessage)
	go service.startListening(service.port, out)

	return out
}

func (service *UdpWireService) startListening(port int, out chan WireMessage) {
	defer close(out)
	ServerAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Println(err)
	}

	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		log.Panic(err)
	}
	defer ServerConn.Close()

	buf := make([]byte, 8024)

	for {
		n, _, err := ServerConn.ReadFromUDP(buf)

		wireMsg, err := service.encoder.Decode(buf, n)
		if err != nil {
			log.Println(err)
			return
		}

		out <- wireMsg

		if err != nil {
			log.Println("Error: ", err)
		}
	}
}
