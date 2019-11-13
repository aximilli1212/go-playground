package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main(){
	path := flag.String("path","myapp.log","Path to log to analyze")
	level := flag.String("level","ERROR","Log level to search for  ERROR")

	f, err :=os.Open(*path)
	if err != nil{
		log.Fatal(err)
	}

	defer f.Close()

	r :=bufio.NewReader(f)
	for {
		 s, err := r.ReadString('\n')
		 if err != nil {
		 	break
		 }

		 if strings.Contains(s, *level){
		 	fmt.Println(s)
		 }
	}
}
