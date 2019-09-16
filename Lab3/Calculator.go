package main
import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "regexp"
)

func main() {
    input := strings.Split(os.Args[1], " ")
    match, _ := regexp.MatchString(`(?m)\d*(\d\.|\.\d)\d*`, input[0])
    match2, _ := regexp.MatchString(`(?m)\d*(\d\.|\.\d)\d*`, input[2])
    if match == true || match2 == true{
    in1, err1 := strconv.ParseFloat(input[0], 64)
    in2, err2 := strconv.ParseFloat(input[2], 64)
    if err1 == nil && err2 == nil{
        var out float64
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
    } else{
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
}