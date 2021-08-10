//Auther: scola
//Date: 2021/08/08 19:35
//Description:
//Σ(っ °Д °;)っ

package main

import (
	"Server_Frame/start"
	"flag"
)

func main() {
	port := flag.Int("port", 12138, "port")
	ginMode := flag.String("gin_mode", "debug", "gin mode: debug/release")
	flag.Parse()
	start.Run(*port, *ginMode)
}
