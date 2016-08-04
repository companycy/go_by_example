package main

import (
	"github.com/fatih/structs"
	."fmt"
)

type Alice struct {
	Id int
	Name string
	IsAdmin bool
	Bob // embed Bob struct
}

type Bob struct {
	Name string
}

func Fill(dest interface{}, src interface{}) {
	mSrc := structs.Map(src)
	mDest := structs.Map(dest)
	for key, val := range mSrc {
		if _, ok := mDest[key]; ok {
			structs.New(dest).Field(key).Set(val)
		}
	}
}

func main() {
	a := Alice{}
	b := Bob {Name : "Bobby"}
	Fill(&a, b)
	Println(a)
	Printf("%v\n", a)
	// Sprintf("%s", a)
}
