package main

import "fmt"

/*
	The prototype pattern delegates the cloning process to the actual object being cloned.
	The pattern declares a common interface for all objects that support cloning,
		which allows you to clone objects without coupling your code to the object's owning class.
	Typically, such an interface contains only one clone method.
*/

var clone_suffix = "_clone"

type inode interface {
	print(string)
	clone() inode
}

//concrete object 1
type file struct {
	name string
}

func (f *file) print(indentation string) {
	fmt.Println(indentation + f.name)
}
func (f *file) clone() inode {
	return &file{
		name: f.name + clone_suffix,
	}
}

//concrete object 2
type folder struct {
	name     string
	children []inode
}

func (f *folder) print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, v := range f.children {
		v.print(indentation + indentation)
	}
}
func (f *folder) clone() inode {
	cloneFolder := &folder{name: f.name + clone_suffix}
	cloneFolder.children = make([]inode, 0, len(f.children))
	for _, v := range f.children {
		cloneFolder.children = append(cloneFolder.children, v.clone())
	}
	return cloneFolder
}

func RunPrototype() {
	folderOrigin := &folder{
		name: "folder origin",
		children: []inode{
			&folder{
				name: "folder inner",
			},
			&file{
				name: "file 1",
			},
			&file{
				name: "file 2",
			},
			&file{
				name: "file 3",
			},
		},
	}
	fmt.Printf("print folder origin:\n")
	folderOrigin.print(" ")
	folderClone := folderOrigin.clone()
	fmt.Printf("print folder clone:\n")
	folderClone.print(" ")
}
