package main

import "fmt"

/*
	Use it to group objects into tree structures and work with them as if they were separate objects.
	1. works when realizing tree-shape object structures
	2. each part can be processed recursively

	We add "c" prefix to each struct/interface to avoid conflict with other design_pattern.go
*/

type cNode interface {
	search(string)
}
type cFile struct {
	name string
}

func (c *cFile) search(key string) {
	fmt.Printf("search key %s in file %s \n", key, c.name)
}

type cFolder struct {
	nodes []cNode
	name  string
}

func (f *cFolder) search(key string) {
	fmt.Printf("search key %s in folder %s recursively\n", key, f.name)
	for _, node := range f.nodes {
		node.search(key)
	}
}
func (f *cFolder) add(c cNode) {
	f.nodes = append(f.nodes, c)
}

func RunComposite() {
	folder := &cFolder{
		name: "outer folder",
		nodes: []cNode{
			&cFile{
				name: "file 1",
			},
			&cFolder{
				name: "inner folder",
				nodes: []cNode{
					&cFile{
						name: "file A",
					},
					&cFile{
						name: "file B",
					},
				},
			},
			&cFile{
				name: "file 2",
			},
		},
	}
	folder.search("xxx")
}
