package main

import (
    "bytes"
    "flag"
    "fmt"
    "log"
    "os"

    "golang.org/x/crypto/ssh"
)

func main() {
    var (
      port int
      //key  string
      pass string
    )
    flag.IntVar(&port, "port", 22, "Port to connect to on the remote host.")
    //key  = flag.String("i", "", "Selects a file from which the identity (private key) for public key authentication.")
    flag.StringVar(&pass, "pass", "", "The password is given on the command line.")

    flag.Parse()

    user := flag.Arg(0)
    host := flag.Arg(1)
    file := flag.Arg(2)

    config := &ssh.ClientConfig{
        User: user,
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
        Auth: []ssh.AuthMethod{
            ssh.Password(pass),
        },
    }

    addr := fmt.Sprintf("%s:%d", host, port)
    conn, err := ssh.Dial("tcp", addr, config)
    if err != nil {
        log.Println(err)
        return
    }
    defer conn.Close()

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

    if err := session.Wait(); err != nil {
        log.Println(err)
        return
    }
}
