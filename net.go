package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
)

const (
	HOST = "127.0.0.1"
	PORT = "8888"
)

func main() {

}

//创建tcp服务器
func ListenTcp() {

	/*
		addr, e := net.ResolveTCPAddr("tcp", ":8083")
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println(addr)
		listener, i := net.ListenTCP("tcp", addr)
		if i != nil {
			fmt.Println(i)
		}
	 */
	listener, el := net.Listen("tcp", HOST+":"+PORT)
	if el != nil {
		log.Fatal(el.Error())
	}
	defer listener.Close()
	for {
		conn, ec := listener.Accept()
		fmt.Println(conn.RemoteAddr())
		if ec != nil {
			log.Fatal(ec.Error())
		}
		go func(conn2 net.Conn) {
			buf := make([]byte, 512)
			_, err := conn.Read(buf) //接收客户端的数据
			if err != nil {
				fmt.Println("error reading", err)
				return
			}
			fmt.Println("client send to server ", string(buf))
			s := "server send to client " + string(buf)
			conn.Write([]byte(s)) //给客户端发送数据
			conn.Close()
		}(conn)
	}
}

//创建tcp客户端
func DialTcp() {
	/*
		addr, e := net.ResolveTCPAddr("tcp", ":8083")
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println(addr)
		conn, i := net.DialTCP("tcp", nil, addr)
		if i != nil {
			fmt.Println(i)
		}
	*/
	conn, e := net.Dial("tcp", "127.0.0.1:8888")
	if e != nil {
		log.Fatal(e)
		return
	}
	defer conn.Close()
	buf := []byte("hello world" )
	read := make([]byte, 512)
	_, err := conn.Write(buf) //往conn中写入数据
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err2 := conn.Read(read) //往conn中读取数据
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(string(read))
	conn.Close()
}

//创建http服务
func HttpServer() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "hello world\n")
	})

	http.HandleFunc("/error", func(writer http.ResponseWriter, request *http.Request) {
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	})

	http.HandleFunc("/redirect", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "http://www.baidu.com", http.StatusMovedPermanently)
	})

	http.HandleFunc("/nofound", func(writer http.ResponseWriter, request *http.Request) {
		http.NotFound(writer, request)
	})

	http.HandleFunc("/file", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "http.txt")
	})

	http.HandleFunc("/cookie", func(writer http.ResponseWriter, request *http.Request) {
		cookie := &http.Cookie{
			Name:     "cookie",
			Value:    "zhangsan",
			Path:     "/",
			Domain:   "127.0.0.1",
			MaxAge:   0,
			Secure:   true,
			HttpOnly: true,
		}
		http.SetCookie(writer, cookie)
	})

	serve := http.ListenAndServe("127.0.0.1:8081", nil)
	if serve != nil {
		fmt.Println(serve)
	}
}

//模拟get数据
func MockRequestByDoGet() {
	client := &http.Client{}
	request, e := http.NewRequest(http.MethodGet, "http://www.baidu.com", nil)
	if e != nil {
		log.Fatal(e)
	}
	response, i := client.Do(request)
	if i != nil {
		log.Fatal(i)
	}
	bytes, i2 := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if i2 != nil {
		log.Fatal(i2)
	}
	fmt.Printf("%s\n", bytes)
}

//模拟post数据
func MockRequestByDoPost() {
	client := &http.Client{}
	request, e := http.NewRequest(
		http.MethodPost,
		"http://devaapi.com/app/like.json",
		strings.NewReader("cid=2"),
	)
	if e != nil {
		log.Fatal(e)
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Game-Imei", "112")
	response, i := client.Do(request)
	if i != nil {
		log.Fatal(i)
	}
	bytes, i2 := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if i2 != nil {
		log.Fatal(i2)
	}
	fmt.Printf("%s\n", bytes)
}
