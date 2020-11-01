package example

import (
	"fmt"
	"github.com/xlab/treeprint"
)

func CreateDiagram(){
	tree := treeprint.New()
	// create a new branch in the root
	one := tree.AddBranch("PPTP")
	// create a new sub-branch
	shareEntity := 	one.AddBranch("MAC Conf Data").
		AddBranch("Mac Service Profile") // add some nodes

	shareEntity.AddBranch("MAC Conf Data"). // add a new sub-branch
		AddBranch("802.1p Mapper").
		AddBranch("GEM IW TP").
		AddBranch("GEM CTP").AddBranch("TCONT")

	shareEntity.AddBranch("MAC Conf Data").
		AddBranch("Multicast IW TP").
		AddBranch("GEM CTP").
		AddBranch("TCONT")

	fmt.Println(tree.String())
}