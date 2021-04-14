package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

func main() {
	// Parsing Arguments
	var (
		port int
		key  string
		pass string
	)
	flag.IntVar(&port, "port", 22, "Port to connect to on the remote host.")
	flag.StringVar(&key, "i", "", "Selects a file from which the identity (private key) for public key authentication.")
	flag.StringVar(&pass, "pass", "", "The password is given on the command line.")

	flag.Parse()

	user := flag.Arg(0)
	host := flag.Arg(1)
	file := flag.Arg(2)

	auth := []ssh.AuthMethod{}

	// Try Password Method
	if pass != "" {
		auth = append(auth, ssh.Password(pass))
	}

	// Try PrivateKey Method
	if key != "" {
		buf, err := ioutil.ReadFile(key)
		if err != nil {
			log.Println(err)
			return
		}
		signer, err := ssh.ParsePrivateKey(buf)
		if err != nil {
			log.Println(err)
			return
		}
		auth = append(auth, ssh.PublicKeys(signer))
	}
	config := &ssh.ClientConfig{
		User:            user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            auth,
	}

	// Try connecting to remote host
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Setup SSH Session
	session, err := conn.NewSession()
	if err != nil {
		log.Println(err)
		return
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	// Setup StdinPipe
	stdin, err := session.StdinPipe()
	if err != nil {
		log.Println(err)
		return
	}
	defer stdin.Close()

	// Start a shell
	if err := session.Shell(); err != nil {
		log.Println(err)
		return
	}

	// Open File
	f, _ := os.Open(file)
	buf := bytes.Buffer{}
	buf.ReadFrom(f)
	buf.WriteTo(stdin)

	// Wait close ssion
	if err := session.Wait(); err != nil {
		log.Println(err)
		return
	}
}
