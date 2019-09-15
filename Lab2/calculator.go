package main
import (
	"fmt"
	"os"
	"strings"
	"strconv"
)
func main() {
	input := strings.Split(os.Args[1], " ")
	in1, err1 := strconv.ParseInt(input[0], 10, 64)
	in2, err2 := strconv.ParseInt(input[2], 10, 64)
	if err1 == nil && err2 == nil{
		var out int64
		if input[1]== "+" {
			out = in1 + in2
		} else if input[1] == "-" { 
			out = in1 - in2 
		} else if input[1] == "*" { 
			out = in1 * in2 
		} else { 
			out = in1 / in2 
		}

		fmt.Println(out)
	}
}

