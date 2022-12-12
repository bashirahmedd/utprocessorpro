package playlist

import "fmt"

func PrintHello() {
	fmt.Println("Hello, Modules! This is mypackage speaking!")
}

//signal to exit after
func signalexit(status bool) bool{

	//rename _exit file to exit
	//if _exit and exit files dont exist, create exit file
	return true
}


func signalshutdown(status bool)  {

	//rename _shutdown file to shutdown
	//if _shutdown and shutdown files dont exist, create shutdown file
	// return false
}

