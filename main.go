package main

import (
	"bufio"
	//for efficient reading of the stdin line by line
	"bytes"
	//imlements functions for manipulation of byte slices to capture command output
	"fmt"
	//to print text on the terminal
	"os"
	//for args and stdin / stdout
	"os/exec"
	// to run external prgrams eg quadA
	"sort"
	//to sort matches alphabetically
	"strings"
	//strings used for trimming and joining or conatenating
)

//helper funtion that runs ine quad executable passing it x and y and it returns what it printed plus any error
func runQuad(quadName string, x, y string) (string, error) {
	//builds a command like ./quadA 5 3 the ./means run the program in the current directory
	cmd := exec.Command("./" + quadName, x, y)

	//creates a buffer to capture the programs standard output	
	var out bytes.Buffer

	// redirects the command stdout into that buffer
	cmd.Stdout = &out

	//executes an external program and waits for it to finish any failure becomes err
	err := cmd.Run()

	//returns whatever the quad printed (as a string) and the error (if any)
	return out.String(), err
}

//entry point of the program
func main(){

	//reads the entire ascii rectangle piped into the program
	input, err := readStdin()

	//if reading stdin failed, report and stop
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	//removes triailing new lines carriage returns from the piped input so comparisons won't fail due to extra end -of-text (handles born /n and \r\n)
	input = strings.TrimRight(input, "\n\r")

	// validates you passed exctly two commands - line args (width and height if not prints usage and exits)
	if len(os.Args) != 3 {
		fmt.Println("Usage: quadchecker <width> <height>")
		return
	}

	// grabs the width and height as strins from cli arguments
	x := os.Args[1]
	y := os.Args[2]

	// the list of extenal executables to test
	quads := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}

	//empty slice to collect which quads match the input
	matches := []string{}

	for _, quad := range quads {
		//loop over each quad name
		output, err := runQuad(quad, x, y)
		if err != nil {
			continue
		}

		//try to run it. if it fails(eg fie missing), skip to the next one
		output = strings.TrimRight(output, "\n\r")

		if output == input {
			matches = append(matches, quad)
		}
	}

	// if nothing atched tell the user and exit
	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	// sort matching strings alphabetically
	sort.Strings(matches)

	//prepare to build the final message parts
	outParts := []string{}

	// for each match format a chunk like quad [5] [3]
	for _, m := range matches {
		outParts = append(outParts, fmt.Sprintf("[%s] [%s] [%s]", m, x, y))
	}

	fmt.Println(strings.Join(outParts, " || "))
}

//helper to read all the text from stdin efficiently into a string builder
func readStdin() (string, error) {
	var b strings.Builder

	// creates a line by line scanner over stdin
	scanner := bufio.NewScanner(os.Stdin)

	// reads each line without the new line appends it then mannually re adds \n
	for scanner.Scan() {
		b.WriteString(scanner.Text())
		b.WriteString("\n")
	}

	// if the scan loop encountered an error io issues report it
	if err := scanner.Err(); err != nil {
		return "", err
	}

	//return the full input as one string wit \n afer each line
	return b.String(), nil
}
