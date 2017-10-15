// Pretty printer for parsing large url query values into Go url.Values
package main

import "flag"
import "fmt"

import "net/url"
import "strings"

func main() {
	query := flag.String("query", "A=B&C=D+E&A=F", "URL query to pretty print")
	flag.Parse()

	values, err := url.ParseQuery(*query)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	for out := range formatValues(values) {
		fmt.Println(out)
	}
}

func formatValues(v url.Values) <-chan string {
	out := make(chan string, len(v))

	go func() {
		out <- fmt.Sprint("v := url.Values{")
		for key, val := range v {
			line := fmt.Sprintf("\t\"%s\": {\"%s\"},", key, strings.Join(val, "\",\""))
			out <- line
		}
		out <- fmt.Sprint("}")
		close(out)
	}()

	return out
}
