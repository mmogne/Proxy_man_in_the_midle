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
func redirect_conn(socket string)(net.Conn , error){
    conn,err := net.Dial("tcp" , socket)
    if err != nil {
        return nil , err 
    }
    return conn ,nil 
    
}


func get_info_in_requests(conn net.Conn) {
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

    requestParts := strings.Fields(data[0])
    hostparts := strings.Fields(data[1])

    if len(requestParts) < 3 {
        return
    }

    methode := requestParts[0]
    path := requestParts[1]
    version := requestParts[2]

    fmt.Println("Methode:", methode)
    fmt.Println("Path:", path)
    fmt.Println("Version:", version)
    
    //fmt.Printf("%T",conn)

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

	go get_info_in_requests(conn)
	
    }


}