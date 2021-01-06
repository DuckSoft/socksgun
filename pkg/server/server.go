package server

import (
	"errors"
	"google.golang.org/grpc/metadata"
	"io"
	"log"
	"net"
	"socksgun/pkg/proto"
)

type SocksGunServiceImpl struct {

}

func RelayStreamToConn(halt chan struct{}, conn net.Conn, stream proto.SocksGunService_TunServer) {
	defer conn.Close()
	for {
		select {
		case <-halt:
			return
		default:
			recv, err := stream.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					log.Printf("conn %v: client read eof", conn.RemoteAddr())
				} else {
					log.Printf("conn %v: client read aborted: %v", conn.RemoteAddr(), err)
				}
				close(halt)
			}
			n, err := conn.Write(recv.Content)
			if err != nil {
				if errors.Is(err, io.EOF) {
					log.Printf("conn %v: remote write eof", conn.RemoteAddr())
				} else {
					log.Printf("conn %v: remote write aborted", conn.RemoteAddr())
				}
				close(halt)
			} else {
				log.Printf("conn %v: %d bytes written to remote", conn.RemoteAddr(), n)
			}
		}
	}
}

func RelayConnToStream(halt chan struct{}, conn net.Conn, stream proto.SocksGunService_TunServer) {
	defer conn.Close()
	buf := make([]byte, 3200)
	for {
		select {
		case <-halt:
			return
		default:
			nRecv, err := conn.Read(buf)
			if err != nil {
				if errors.Is(err, io.EOF) {
					log.Printf("conn %v: remote read eof", conn.RemoteAddr())
				} else {
					log.Printf("conn %v: remote read aborted: %v", conn.RemoteAddr(), err)
				}
				close(halt)
			}
			err = stream.Send(&proto.Packet{Content: buf[:nRecv]})
			if err != nil {
				if errors.Is(err, io.EOF) {
					log.Printf("conn %v: client write eof", conn.RemoteAddr())
				} else {
					log.Printf("conn %v: client write aborted", conn.RemoteAddr())
				}
				close(halt)
			} else {
				log.Printf("conn %v: %d bytes written to client", conn.RemoteAddr(), nRecv)
			}
		}
	}
}

func (s SocksGunServiceImpl) Tun(stream proto.SocksGunService_TunServer) error {
	headers, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("failed to fetch any metadata")
	}

	remote, err := net.Dial("tcp", headers["addr"][0])
	if err != nil {
		return errors.New("failed to parse IP")
	}

	halt := make(chan struct{})
	go RelayStreamToConn(halt, remote, stream)
	RelayConnToStream(halt, remote, stream)

	return nil
}




