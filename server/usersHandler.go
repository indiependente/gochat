package main

import (
	"crypto/rand"
	"fmt"
	"gochat/common"
)

func (s *Server) addUser(token string, u *User) error {
	s.usrLock.Lock()
	defer s.usrLock.Unlock()
	if err := s.addToken(u.name, token); err != nil {
		return err
	}
	s.users[token] = u
	return nil
}

func (s *Server) getUser(token string) (*User, error) {
	s.usrLock.RLock()
	defer s.usrLock.RUnlock()
	u, ok := s.users[token]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func (s *Server) deleteUser(token string) (*User, error) {
	s.usrLock.Lock()
	defer s.usrLock.Unlock()
	u, ok := s.users[token]
	if !ok {
		common.Errorf("%s %v", ok, u)
		return nil, fmt.Errorf("user not found")
	}
	if err := s.deleteToken(u.name); err != nil {
		return u, err
	}
	delete(s.users, token)
	return u, nil
}

func (s *Server) addToken(name, token string) error {
	s.tkLock.Lock()
	defer s.tkLock.Unlock()
	if _, ok := s.tokens[name]; ok {
		return fmt.Errorf("user already existent")
	}
	s.tokens[name] = token
	return nil
}

func (s *Server) getToken(name string) (string, error) {
	s.tkLock.RLock()
	defer s.tkLock.RUnlock()
	tk, ok := s.tokens[name]
	if !ok {
		return "", fmt.Errorf("user not found")
	}
	return tk, nil
}

func (s *Server) deleteToken(name string) error {
	s.tkLock.Lock()
	defer s.tkLock.Unlock()
	if _, ok := s.tokens[name]; !ok {
		return fmt.Errorf("token not found")
	}
	delete(s.tokens, name)
	return nil
}

func createToken(name string) string {
	tkn := make([]byte, 4)
	rand.Read(tkn)
	return fmt.Sprintf("%x", tkn)
}
