package main

import "CacheCLI/cmd"

func main() {
	//comment out one of the below lines for demo depending on if project is running server or CLI
	//NOTE: import will change as well
	
	//go server.StartServer("8080")
	cmd.Execute()

}
