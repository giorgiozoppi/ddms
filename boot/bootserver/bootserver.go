package bootserver

import (
	"context"
	"crypto/tls"
	"github.com/giorgiozoppi/ddms/pkg/common"
	"github.com/lucas-clemente/quic-go"
	"log"
)

type SuperNode struct {
	NodeId     common.ID
	members    *common.NodeList
	siblings   *common.NodeList
	scoreboard chan quic.Stream
	IpAddress  string
	Context    context.Context
	CertFile   string
	KeyFile    string
}

func NewSuperNode(address string, certFile string, keyFile string) *SuperNode {
	errorNode, nodeId := common.NewID()
	if errorNode != nil {
		panic(errorNode)
	}
	return &SuperNode{
		NodeId:     *nodeId,
		IpAddress:  address,
		CertFile:   certFile,
		KeyFile:    keyFile,
		scoreboard: make(chan quic.Stream),
		members:    common.NewEmptyList(),
		siblings:   common.NewEmptyList(),
	}
}
func (node *SuperNode) Run() {
	cert, err := tls.LoadX509KeyPair(node.CertFile, node.KeyFile)
	if err != nil {
		log.Fatal(err)
	}
	cfg := &tls.Config{Certificates: []tls.Certificate{cert}}
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	listener, err := quic.ListenAddr(node.IpAddress, cfg, nil)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	for {
		sess, err := listener.Accept(context.Background())
		if err != nil {
			log.Fatalf("Error %v", err)
		}
		quickStream, err := sess.AcceptStream(context.Background())
		if err != nil {
			log.Fatalf("Error %v", err)
		}
		node.scoreboard <- quickStream
	}
}
