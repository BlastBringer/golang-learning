package peers 

import(
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)


type Peer struct{
	IP net.IP
	Port uint16 
}

//unmarshal parses peer IP addreses and the ports from a buffer
func Unmarshal(peersBin []byte) ([]Peer, error){
	const Peersize = 6 //4 for ip , 2 for byte 
	numPeers := len(peersBin)
	if len(peersBin)%Peersize != 0{
		err := fmt.Errorf("Received malformed papers")
		return nil, err 
	}

	peers := make([]Peer, numPeers)
	for i := 0; i <numPeers; i++{
		offset := i * Peersize
		peers[i].IP = net.IP(peersBin[offset : offset+4])
		peers[i].Port = binary.BigEndian.Uint16([]byte(peersBin[offset+4 : offset+6]))
	}
	return peers, nil 
}

func (p Peer) String() string{
	return net.JoinHostPort(p.IP.String(), strconv.Itoa(int(p.Port)))
}