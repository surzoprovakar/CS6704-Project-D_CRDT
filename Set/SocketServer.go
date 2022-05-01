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
var set *Set

var conns []net.Conn

var id2trust map[int]float64

func do_actions(actions []Operation2Value) {

	//sleep for 10 secs, so other replicase
	//have time to get started
	time.Sleep(10 * time.Second)
	fmt.Println("Starting to do_actions")

	for _, act2val := range actions {
		if act2val.action == "Add" {
			set.Add(act2val.operator)
			set.Print()

		} else if act2val.action == "Remove" {
			set.Remove(act2val.operator)
			set.Print()

		} else if act2val.action == "Broadcast" {
			fmt.Println("processing Broadcast")
			if conns == nil { //establish connecitons on first broadcast
				conns = EstablishConnections(hosts)
			}
			fmt.Print("About to broadcast Counter")
			set.Print()
			Broadcast(conns, set.ToByteArray())
		} else { //assume it is delay
			var err error
			var number int
			if number, err = strconv.Atoi(act2val.action); err != nil {
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

	//establish connections using the addresses from the first input file
	//read the execution steps from the second input file
	//execute the script
	var err error
	var id int
	if id, err = strconv.Atoi(input[0]); err != nil {
		panic(err)
	}
	set = NewSet(id)

	ip_address := input[1]

	hosts = ReadFile(input[2])
	actions := ReadFile_Actions(input[3])
	//id2trust = ReadFile_TrustGradient(input[4])
	//trusts_list := get_trust_from_map(hosts)

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

	buffer := make([]byte, 704)
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

	tempSet := FromByteArray(buffer[0:n])

	set.Merge(tempSet)

	// Restart the process.
	//handleConnection(conn)
}
