package main

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"io/ioutil"
	"os"
)

func getDecoder(enc string) *encoding.Decoder {
	return charmap.Windows1252.NewDecoder()
}

func getEncoder(encoding string) *encoding.Encoder {
	return charmap.Windows1252.NewEncoder()
}

func convertFile(inputFile, inputEncoding, outputFile, outputEncoding string) error {
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("cannot read file: %w", err)
	}
	contentStr := string(content)
	encoder := getEncoder(outputEncoding)
	encodedStr, err := encoder.String(contentStr)
	if err != nil {
		return fmt.Errorf("cannot encode string: %w", err)
	}
	err = ioutil.WriteFile(outputFile, []byte(encodedStr), os.ModePerm)
	if err != nil {
		return fmt.Errorf("cannot write to output file: %w", err)
	}
	return nil
}

func main() {
	var inputFile string
	var inputEncoding string
	var outputFile string
	var outputEncoding string
	flag.StringVar(&inputFile, "inputFile", "", "The file to watch")
	flag.StringVar(&inputEncoding, "inputEncoding", "", "The input file encoding")
	flag.StringVar(&outputFile, "outputFile", "", "The file to write the converted string to")
	flag.StringVar(&outputEncoding, "outputEncoding", "", "The output file encoding")
	flag.Parse()
	if inputFile == "" || inputEncoding == "" || outputFile == "" || outputEncoding == "" {
		fmt.Println("Usage:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("Error creating watcher: %v\r\n", err)
		os.Exit(1)
	}
	defer watcher.Close()

	//
	done := make(chan bool)

	//
	go func() {
		for {
			select {
			// watch for events
			case <-watcher.Events:
				err := convertFile(inputFile, inputEncoding, outputFile, outputEncoding)
				if err != nil {
					fmt.Printf("Error converting file: %v\r\n", err)
				}

				// watch for errors
			case err := <-watcher.Errors:
				fmt.Printf("Error in watcher: %v\r\n", err)
				os.Exit(1)
			}
		}
	}()

	if err := watcher.Add(inputFile); err != nil {
		fmt.Printf("Error watching file: %v\r\n", err)
		os.Exit(1)
	}

	<-done
}