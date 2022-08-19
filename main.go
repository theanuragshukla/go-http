package main
import (
	"fmt"
	"net/http"
	"io"
)

func getRoot(res http.ResponseWriter,req *http.Request){
	fmt.Println("getRoot called")
	io.WriteString(res,"Hello World")
}

func main(){
	http.HandleFunc("/",getRoot)
	http.ListenAndServe(":3000",nil)
}


