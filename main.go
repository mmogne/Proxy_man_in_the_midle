package main


import (
	"fmt"
	"bufio"
	"net"
	"strings"
    "Proxy_man_in_the_midle/urlutils"
    "Proxy_man_in_the_midle/redirect_request"
)


func redirect_conn(host string,port string)(net.Conn , error){
    target := net.JoinHostPort(host, port) 
    conn,err := net.Dial("tcp" , target)

    //fmt.Println(host)
    if err != nil {
        return nil , err 
    }

    return conn ,nil 
    
}

type ReadRequests struct {
    Methode string 
    Path string
    Version string
    Url_request string
    Requests_Data []string 


}




func get_info_in_requests(conn net.Conn)ReadRequests {
    var info ReadRequests

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
        return info
    }

    requestParts := strings.Fields(data[0])
    //hostparts := strings.Fields(data[1])
    //fmt.Println(requestParts[1])
    
    
    
    
    if len(requestParts) < 3 {
        return info
    }
    url_request := requestParts[1]

    methode := requestParts[0]
    path := requestParts[1]
    version := requestParts[2]
    info = ReadRequests{
        Methode : methode,
        Path: path,
        Version: version,
        Url_request: url_request,
        Requests_Data:data,

    }
    return info


}
func HandleConnection(conn net.Conn){
    defer conn.Close()
    
    header_requests:= get_info_in_requests(conn)
    fmt.Println(
    "Methode:", header_requests.Methode,
    "Target:", header_requests.Url_request,
)
if header_requests.Methode == "CONNECT" {
    return
}

    list_url_port,err:=urlutils.ManageURLs(header_requests.Url_request)
    if err != nil {
        fmt.Println(nil) 
        return  
    }
    if len(list_url_port) < 2 {
        fmt.Println("URL invalide ou port introuvable")
        return
}

    host:=list_url_port[0]
    port:=list_url_port[1]
    serv_conn_distant,error := redirect_conn(host,port )
       if error != nil {
        fmt.Println(error)
        return 
    }
    
    redirect_request.Redirect_request(serv_conn_distant,header_requests.Requests_Data)
    redirect_request.Get_response(serv_conn_distant , conn)
    //_=data_response
    //fmt.Println("data response: ",data_response)

    //fmt.Println("header_requests: " , header_requests.Requests_Data)
 
    return 

    
    

    

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

	go HandleConnection(conn)
	
    }


}