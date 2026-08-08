package urlutils 
import ( 		
	
		"strings"
		"net/url"
)







func ManageURLs ( rawurl string) ([]string, error) {
	port :=""
	list_url_split:=[]string{}

	url_parse,err:= url.Parse(rawurl)
	
	if err != nil {
		return nil , err 
		}

	host_url:=url_parse.Host
	
	
	if strings.Contains(host_url,":"){
	
		list_url_split := strings.Split(host_url,":")
	
		return list_url_split,nil 

	}else if   url_parse.Scheme == "https"{
	
		port = "443"
		list_url_split=[]string{url_parse.Host,port}
		return list_url_split,nil

	}else if url_parse.Scheme == "http"{
		port = "80"
		list_url_split=[]string{url_parse.Host,port}
		return list_url_split,nil 

	} 
	return list_url_split , nil 
	}





	


