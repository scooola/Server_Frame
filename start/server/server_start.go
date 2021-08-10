//Auther: scola
//Date: 2021/08/08 20:03
//Description:
//Σ(っ °Д °;)っ

package server

import _ "Server_Frame/action"

func SeverStart(port int, ginMode string) {
	//config
	initConfig()

	//gin router
	initRouter(ginMode)
	initFrontend()
	initHealthCheck()
	routerRun(port)
}
