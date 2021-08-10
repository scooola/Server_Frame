//Auther: scola
//Date: 2021/08/08 20:01
//Description:
//Σ(っ °Д °;)っ

package start

import "Server_Frame/start/server"

func Run(port int, ginMode string) {
	server.SeverStart(port, ginMode)
}
