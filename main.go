package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
	file, err := os.Open("test_input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
     }
     defer file.Close()
	
	
	
	m := make(map[string][]string)

     scanner := bufio.NewScanner(file)
    
	for i := 0; scanner.Scan(); i++ {
          line := scanner.Text()
		b, a, f := strings.Cut(line, ":")
	     valuesStr := strings.TrimSpace(a)
		values := strings.Fields(valuesStr)
		if !f {
			m["season"] = []string{b}
			break
		}
		m[b] = values
        
     }
	fmt.Print(m)

	
	
}