package main 

import "os"
import "fmt"
import "bufio"


func main(){
	readInput("test.md")
}

func readInput(fileName string){
 file,err := os.Open(fileName)

 if err != nil {
 	fmt.Println(err)
 }
 scanner := bufio.NewScanner(file)

 var lines []string
 for(scanner.Scan()){

 	lines = append(lines, scanner.Text())
 }
 return lines
}