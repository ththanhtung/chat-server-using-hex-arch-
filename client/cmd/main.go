package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:3000")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("Error connecting to UDP:", err)
		return
	}

	defer conn.Close()

	fmt.Print("Enter your name (Input NOT allow space): ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = name[:len(name)-1] // remove newline character

	registerMessage := []byte(fmt.Sprintf("@register %s", name))
	_, err = conn.Write(registerMessage)
	if err != nil {
		fmt.Println("Error registering with server:", err)
		return
	}

	go receiveMessages(conn)

	for {
		message, _ := reader.ReadString('\n')
		message = message[:len(message)-1] // remove newline character

		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}

func receiveMessages(conn *net.UDPConn) {
	for {
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP:", err)
			continue
		}

		message := string(buffer[:n])
		fmt.Println(message)
	}
}
