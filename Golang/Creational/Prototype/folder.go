package main

import "fmt"

type Folder struct {
	children []INode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() INode {
	var tempChildren []INode
	cloneFolder := &Folder{
		children: nil,
		name:     f.name + "_clone",
	}

	for _, i := range f.children {
		childCopy := i.clone()
		tempChildren = append(tempChildren, childCopy)
	}
	cloneFolder.children = tempChildren

	return cloneFolder
}
