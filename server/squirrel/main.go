package squirrel

import "fmt"

func init() {
	printLogo()
}

func printLogo() {
	fmt.Printf("                 _               _ \n ___  __ _ _   _(_)_ __ _ __ ___| |\n/ __|/ _` | | | | | '__| '__/ _ \\ |\n\\__ \\ (_| | |_| | | |  | | |  __/ |\n|___/\\__, |\\__,_|_|_|  |_|  \\___|_|\n        |_|                        \n")
}

func Main(args []string) {
	checkSupportedArch()

	loadConfig()

	startSquirrel(args)
}
