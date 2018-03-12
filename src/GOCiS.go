package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "flag"
)


const VERSION = "0.0.1"

var info bool
var input, output string

func init() {
	flag.BoolVar(&info, "info", false, "Version")
	flag.StringVar(&input, "in", "", "Select the input CSV")
	flag.StringVar(&output, "out", "", "Select the output CSV name")
}

func main() {

  flag.Parse()

	if info {
		fmt.Println("GOCiS " + VERSION)
		os.Exit(0)
	}

  if input == "" || output == "" {
		fmt.Println("Select input CSV")
		os.Exit(1)
	}

  fileinput, _ := os.Open(input)
  defer fileinput.Close()
  fileScanner := bufio.NewScanner(fileinput)

  fileout, _ := os.Create(output)
  defer fileout.Close()

  for fileScanner.Scan() {
    z := strings.SplitN(fileScanner.Text(), ",", -1)

		for i:=0; i < len(z); i++ {

			var str = z[i]

			if str[:1] == "=" {
        z[i] = strings.Replace(str, "=", " =", 1)
			}
			if str[:1] == "+" {
        z[i] = strings.Replace(str, "+", " +", 1)
			}
			if str[:1] == "-" {
        z[i] = strings.Replace(str, "-", " -", 1)
			}
			if str[:1] == "@" {
        z[i] = strings.Replace(str, "@", " @", 1)
			}
    }
     fileout.WriteString(strings.Join(z, ","))
     fileout.WriteString("\n")
	}
}
