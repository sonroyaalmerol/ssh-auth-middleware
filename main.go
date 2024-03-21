package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"net"

	"golang.org/x/crypto/ssh"
)

func main() {
	// Listen for incoming SSH connections
	listener, err := net.Listen("tcp", "localhost:2222")
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}
	defer listener.Close()
	log.Println("SSH proxy started on localhost:2222")

	for {
		// Accept incoming connection
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		log.Printf("Accepted connection from: %s", conn.RemoteAddr())

		// Perform custom authentication
		authenticated := performSSHAuth(conn)
		if !authenticated {
			conn.Close()
			continue
		}

		// Connect to the actual SSH server
		// For this example, we're not connecting to the SSH server

		// Proxy SSH traffic bidirectionally
		// For this example, we're not proxying traffic
	}
}

func performSSHAuth(conn net.Conn) bool {
	config := &ssh.ServerConfig{
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			// Implement your authentication logic here
			// For example, you can check the username and password against a database
			if c.User() == "user23" && string(pass) == "password" {
				return nil, nil // Authentication successful
			}
			return nil, fmt.Errorf("authentication failed")
		},
	}
	config.AddHostKey(generateRSAHostKeys())

	// Accept the connection as a new SSH connection
	sshConn, chans, reqs, err := ssh.NewServerConn(conn, config)
	if err != nil {
		log.Printf("Failed to accept SSH connection: %v", err)
		return false
	}
	defer sshConn.Close()

	// Discard out-of-band requests
	go ssh.DiscardRequests(reqs)

	// Accept channels
	go func() {
		for newChannel := range chans {
			// Reject all incoming channels
			newChannel.Reject(ssh.UnknownChannelType, "not supported")
		}
	}()

	// SSH authentication successful
	return true
}

func generateRSAHostKeys() (privateKey ssh.Signer) {
	// Generate RSA private key
	rsaPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Failed to generate RSA private key: %v", err)
	}

	// Encode private key to PEM format
	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(rsaPrivateKey),
	}
	privateKeyBytes := pem.EncodeToMemory(privateKeyPEM)

	// Parse private key
	privateKey, err = ssh.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	return privateKey
}
