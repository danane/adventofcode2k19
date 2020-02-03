package alarm

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func ParseIntcode(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	return scanner.Text()
}

func IntcodeToSlice(intcode string) (intCodeSlice []int) {
	stringCode := strings.Split(intcode, ",")
	for _, v := range stringCode {
		i, _ := strconv.Atoi(v)
		intCodeSlice = append(intCodeSlice, i)
	}
	return
}

func Exec(cmds []int, index int) (terminated bool) {
	switch cmds[index] {
	case 1:
		cmds[cmds[index+3]] = cmds[cmds[index+1]] + cmds[cmds[index+2]]
	case 2:
		cmds[cmds[index+3]] = cmds[cmds[index+1]] * cmds[cmds[index+2]]
	case 99:
		terminated = true
		fmt.Fprintf(ioutil.Discard, "exiting program...")
	default:
		panic("something went wrong")
	}
	return
}

func ComputeOutput(intcodes []int) int {
	codeIndex := 0
	var terminated bool

	for {
		terminated = Exec(intcodes, codeIndex)
		if terminated {
			break
		}
		codeIndex += 4
	}
	return intcodes[0]
}
