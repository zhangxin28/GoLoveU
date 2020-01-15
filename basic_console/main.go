package main

import (
	"fmt"
	"goloveu/goinactioncode"
	_ "goloveu/goinactioncode/executers"
	gua "goloveu/utils/avatar"
	"os"
)

func main() {
	fmt.Println("Fuck go First")
	
	pngBytes, err := gua.Generate(123123)
	if err != nil {
		fmt.Println(err)
	}
	fi, _ := os.Create("123123.png")
	fi.Write(pngBytes)
	fi.Close()	

	goinactioncode.RunSample("fdfsdfsdf")
}
