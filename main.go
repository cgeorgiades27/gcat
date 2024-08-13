package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	nums   = flag.Bool("n", false, "print number lines")
	space  = flag.Bool("b", false, "print blank space after line")
	reader = os.Stdin
)

func main() {
	flag.Parse()

	var files []*os.File
	for _, arg := range os.Args[1:] {
		if arg[0] == '-' {
			continue
		}
		file, err := os.Open(arg)
		if err != nil {
			log.Fatalf("bad file: %v", err)
		}
		files = append(files, file)
	}

	if len(files) == 0 {
		files = append(files, reader)
	}

	lineNum := 1
	for _, file := range files {
		scanner := bufio.NewScanner(file)
		writer := bufio.NewWriter(os.Stdout)
		for scanner.Scan() {
			line := scanner.Text()
			if *nums {
				line = fmt.Sprintf("%d %s\n", lineNum, line)
			}
			writer.WriteString(line)
			if *space {
				writer.WriteByte('\n')
			}
			writer.Flush()
			lineNum++
		}
		file.Close()
	}
}
