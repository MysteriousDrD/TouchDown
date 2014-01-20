package main 

import "os"
import "fmt"
import "bufio"
import "regexp"
import "strings"

func main(){
	input := readLines("test.md")

	for i := 0; i < len(input); i ++{
		if checkForLists(input[i]){
			fmt.Println(input[i])
			modifyListLine(input[i])
		}
		
	}
}

func checkForLists(line string) (bool){
	var isListElement = regexp.MustCompile(`\*`) //use raw strings here
	return isListElement.MatchString(line)
}

func modifyListLine(line string) (string){
	loc := strings.Index(line, "*")
	if len(line) > 1 { //if it's not just a single trailing *
		slice := line[loc+1:len(line)]
		s := []string{"<li>", slice, "</li>"}
		fmt.Println(strings.Join(s, ""))
	}


	return line
}


func readLines(fileName string) ([]string){
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