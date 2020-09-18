package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)
func saySayHello (w http.ResponseWriter, r *http.Request) {
	_=r.ParseForm()//3解析参数，默认是不会解析的
	fmt.Println(r.Form)//4输出到服务器端的打印信息
	fmt.Println("Path:",r.URL.Path)
	fmt.Println("Host:",r.Host)

	for k, v := range r.Form {
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v," "))
	}
	_,_ = fmt.Fprintf(w,"Hello Web, %s!",r.Form.Get("name"))//5写入到w的是输出到客户端的信息
}

func main() {
	http.HandleFunc("/",saySayHello)//1设置访问路由
	err := http.ListenAndServe(":8080",nil)//2设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe:",err)
	}
}
