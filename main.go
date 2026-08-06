package main

import (
	"fmt"
	"bufio"
	"net"
	"strings"
)

//func get_host(conn){
//
//}

func receive_http_request(conn net.Conn) {
    defer conn.Close()

    scanner := bufio.NewScanner(conn)

    data := []string{}

    for scanner.Scan() {
        request := scanner.Text()
		

        if request == "" {
            break
        }

        data = append(data, request)
    }

    if len(data) == 0 {
        return
    }

    champs := strings.Fields(data[0])

    if len(champs) < 3 {
        return
    }

    methode := champs[0]
    path := champs[1]
    version := champs[2]

    fmt.Println("Methode:", methode)
    fmt.Println("Path:", path)
    fmt.Println("Version:", version)
}

func main() {
    ln, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Erreur au niveau de l'écoute tcp")
	}
	for {
		conn ,err:= ln.Accept()

		if err != nil {
    	   continue

	}

	go receive_http_request(conn)
	
    }


}