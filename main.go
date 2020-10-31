package main

import (
	"fmt"
	"github.com/CrowdSurge/banner"
)

func main() {
	fmt.Print("OMCI Diagram")
	createBanner()
	scanFile("/home/ozgecan/CHT/1/omci.log")
}
func createBanner(){
	banner.Print("omci diagram")
}
