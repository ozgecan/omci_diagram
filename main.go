package main

import (
	"bufio"
	"fmt"
	"github.com/CrowdSurge/banner"
	"os"
	"strings"
	example "github/omci_diagram/example"
)

func main() {
	fmt.Print("OMCI Diagram")
	createBanner()
	//path := scanConsole()
	//scanFile(path)
	example.CreateDiagram()
}
func scanConsole() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("----------Please Give Capture Path--------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if isValid(text) {
			fmt.Println("Your file path:", text)
			return text
		}
	}
}

func isValid(fp string) bool {
	if _, err := os.Stat(fp); err == nil {
		return true
	}else{
		return false
	}

}
func createBanner(){
	banner.Print("omci diagram")
}
