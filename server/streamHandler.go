package main

import (
	"context"
	"fmt"
	common "gochat/common"
	chat "gochat/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const tokenHeader = "x-chat-token"

func (s *server) openStream(tk string) chan chat.StreamResponse {
	stream := make(chan chat.StreamResponse, 100)

	s.usrLock.Lock()
	s.users[tk].stream = stream
	s.usrLock.Unlock()

	return stream
}

func (s *server) closeStream(tk string) error {
	s.usrLock.Lock()
	defer s.usrLock.Unlock()
	u, ok := s.users[tk]
	if !ok {
		return fmt.Errorf("stream for token %s not found", tk)
	}
	close(u.stream)
	return nil
}

func extractToken(ctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md[tokenHeader]) == 0 {
		return "", false
	}

	return md[tokenHeader][0], true
}

// broadcastMessages opens a user's stream and sends the
// StreamResponses that are pushed on that stream
func (s *server) broadcastMessages(srv chat.Chat_StreamServer, tk string) {
	stream := s.openStream(tk)
	defer s.closeStream(tk)
	for {
		select {
		case <-srv.Context().Done():
			return
		case res := <-stream:
			err := srv.Send(&res)
			if s, ok := status.FromError(err); ok {
				switch s.Code() {
				case codes.OK:
					// all fine, no op
				case codes.Unavailable, codes.Canceled, codes.DeadlineExceeded:
					common.Debugf("client (%s) terminated the connection", tk)
					return
				default:
					common.ServerLogf("failed to send to client (%s)", tk)
					return
				}
			} else {
				common.Errorf("ERROR: %v\n", err)
			}
		}
	}
}

func (s *server) broadcast() {
	for res := range s.broadcastCh {
		s.usrLock.RLock()
		for tk, u := range s.users {
			select {
			case u.stream <- res:
				// response sent to client's channel!
			default:
				common.ServerLogf("client (%s) stream full, dropping message", tk)
			}
		}
		s.usrLock.RUnlock()
	}
}
