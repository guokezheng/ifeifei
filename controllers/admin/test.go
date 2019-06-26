package admin

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/",handler)
	http.ListenAndServe("localhost:8080",nil)
}

func handler(rw http.ResponseWriter, req *http.Request){
		req.ParseForm();
		if len(req.Form["name"]) > 0 {
			fmt.Fprint(rw,"hello,",req.Form["name"][0])
		}else{
			fmt.Fprint(rw,"hello world")
		}
}
