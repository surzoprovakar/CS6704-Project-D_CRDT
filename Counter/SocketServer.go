package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

// Application constants, defining host, port, and protocol.
const (
	//connHost = "localhost"
	//connPort = "8080"
	connType = "tcp"
)

var hosts []string
var counter *Counter
var conns []net.Conn

func do_actions(actions []string) {

	//sleep for 10 secs, so other replicase
	//have time to get started
	time.Sleep(10 * time.Second)
	fmt.Println("Starting to do_actions")

	for _, action := range actions {
		if action == "Inc" {
			counter.Inc()
			counter.inc_list[counter.id-1] += 1
			counter.Print()

		} else if action == "Dec" {
			counter.Dec()
			counter.dec_list[counter.id-1] += 1
			counter.Print()
		} else if action == "Broadcast" {
			fmt.Println("processing Broadcast")
			if conns == nil { //establish connecitons on first broadcast
				conns = EstablishConnections(hosts)
			}
			fmt.Println("About to broadcast Counter")
			counter.Print()
			Broadcast(conns, counter.ToByteArray())
		} else { //assume it is delay
			var err error
			var number int
			if number, err = strconv.Atoi(action); err != nil {
				panic(err)
			}

			time.Sleep(time.Duration(number) * time.Second)
		}

	}
}

func main() {

	input := os.Args[1:]
	if len(input) != 4 {
		println("Usage: counter_id ip_address crdt_socket_server Replicas'_Addresses.txt Actions.txt")
		os.Exit(1)
	}

	var err error
	var id int
	if id, err = strconv.Atoi(input[0]); err != nil {
		panic(err)
	}
	counter = NewCounter(id)

	ip_address := input[1]

	hosts = ReadFile(input[2])
	actions := ReadFile(input[3])

	// Start the server and listen for incoming connections.
	fmt.Println("Starting " + connType + " server on " + ip_address)
	l, err := net.Listen(connType, ip_address)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()

	do_actions(actions)

	// run loop forever, until exit.
	for {
		// Listen for an incoming connection.
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Client connected.")

		// Print client connection address.
		fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

		// Handle connections concurrently in a new goroutine.
		go HandleConnection(c)
	}
}

// handleConnection handles logic for a single connection request.
func HandleConnection(conn net.Conn) {
	// Buffer client input until a newline.
	//buffer, err := bufio.NewReader(conn).ReadBytes('\n')

	buffer := make([]byte, 448)
	//c := bufio.NewReader(conn)
	fmt.Println("starting to read")
	n, err := conn.Read(buffer)
	// Close left clients.
	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}
	fmt.Println("Read bytes from socket", n)
	tempCounter := FromByteArray(buffer)
	counter.Merge(tempCounter)

	// Restart the process.
	//handleConnection(conn)
}
