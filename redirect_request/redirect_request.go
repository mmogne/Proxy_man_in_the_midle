package redirect_request
import (
	"fmt"
	"net"
	"io"

)



func Redirect_request (conn net.Conn , data []string){
	for _ , line := range data {
		fmt.Fprint(conn , line+"\r\n")
	}
	fmt.Fprint(conn, "\r\n")

}
func Get_response(conn net.Conn , conn_navigateur net.Conn){
	bytes,err :=io.Copy(conn_navigateur, conn)
	if err != nil {
		fmt.Println("Erreur pendant la copie :", err)
        return
    }

    fmt.Println("Octets transférés :", bytes)
}






