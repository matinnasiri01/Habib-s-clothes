package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	

	m := make(map[string][]string)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		key, val, found := strings.Cut(line, ":")
		valuesStr := strings.TrimSpace(val)
		values := strings.Fields(valuesStr)
		if !found {
			m["season"] = []string{line}
			break
		}
		m[strings.TrimSpace(key)] = values
	}

	season := m["season"][0]

	for _, s := range m["SHIRT"] {
		for _, p := range m["PANTS"] {
			
			coats := []string{""}
			if season == "SPRING" || season == "FALL" || season == "WINTER" {
				for _, co := range m["COAT"] {
					
					if season == "FALL" && (co == "yellow" || co == "orange") {
						continue
					}
					coats = append(coats, co)
				}
			}

			caps := []string{""}
			if season == "SUMMER" {
				caps = m["CAP"]
			} else if season == "SPRING" || season == "FALL" {
				caps = append(caps, "") 
				caps = append(caps, m["CAP"]...)
			}

			
			jackets := []string{""}
			if season == "WINTER" {
				jackets = append(jackets, m["JACKET"]...)
			}

			for _, co := range coats {
				for _, ca := range caps {
					for _, ja := range jackets {
						
						if season == "WINTER" {
							if co != "" && ja != "" {
								continue
							}
							if co == "" && ja == "" {
								continue
							}
						}

						line := ""
						if co != "" {
							line += fmt.Sprintf("COAT: %s ", co)
						}
						line += fmt.Sprintf("SHIRT: %s PANTS: %s", s, p)
						if ca != "" {
							line += fmt.Sprintf(" CAP: %s", ca)
						}
						if ja != "" {
							line += fmt.Sprintf(" JACKET: %s", ja)
						}
						fmt.Println(line)
					}
				}
			}
		}
	}
}
