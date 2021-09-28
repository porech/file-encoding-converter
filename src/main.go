package main

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/rgeoghegan/tabulate"
	"io/ioutil"
	"os"
	"strings"
)

var inputEncDec EncoderDecoder
var outputEncDec EncoderDecoder

func convertFile(inputFile, outputFile string) error {
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("cannot read file: %w", err)
	}
	contentStr, err := inputEncDec.Decode(content)
	if err != nil {
		return fmt.Errorf("cannot decode input: %v", err)
	}
	encodedBytes, err := outputEncDec.Encode(contentStr)
	if err != nil {
		return fmt.Errorf("cannot encode string: %w", err)
	}
	err = ioutil.WriteFile(outputFile, encodedBytes, os.ModePerm)
	if err != nil {
		return fmt.Errorf("cannot write to output file: %w", err)
	}
	return nil
}

func printSupportedEncodings() {
	encoders := make([][]string, len(availableEncodersDecoders))
	for i, encoder := range availableEncodersDecoders {
		names := strings.Join(encoder.Names, ", ")
		encoders[i] = []string{encoder.Description, names}
	}

	table, _ := tabulate.Tabulate(encoders, &tabulate.Layout{
		Format:      tabulate.FancyGridFormat,
		HideHeaders: false,
		Headers:     []string{"Description", "Accepted names"},
	})

	fmt.Println(table)
}

func printHelp() {
	fmt.Printf("Convert files from one encoding to another, optionally watching the input file to convert it on any change.\r\n\r\n")
	fmt.Println("Usage:")
	flag.PrintDefaults()
	fmt.Printf("\r\n")
	fmt.Println("Supported encodings:")
	printSupportedEncodings()
}

func main() {
	var inputFile string
	var inputEncoding string
	var outputFile string
	var outputEncoding string
	var watch bool
	flag.StringVar(&inputFile, "inputFile", "", "The input file to read from")
	flag.StringVar(&inputEncoding, "inputEncoding", "", "The current input file encoding")
	flag.StringVar(&outputFile, "outputFile", "", "The file to write the converted string to")
	flag.StringVar(&outputEncoding, "outputEncoding", "", "The wanted output file encoding")
	flag.BoolVar(&watch, "watch", false, "Keep the process running and watch the input file for changes")
	flag.Parse()

	if inputFile == "" || inputEncoding == "" || outputFile == "" || outputEncoding == "" {
		printHelp()
		os.Exit(1)
	}

	inputEncDec = getEncoderDecoder(inputEncoding)
	if inputEncDec == nil {
		fmt.Printf("Invalid input encoding: %s\r\n", inputEncoding)
		printHelp()
		os.Exit(1)
	}

	outputEncDec = getEncoderDecoder(outputEncoding)
	if outputEncDec == nil {
		fmt.Printf("Invalid output encoding: %s\r\n", outputEncoding)
		printHelp()
		os.Exit(1)
	}

	err := convertFile(inputFile, outputFile)
	if err != nil {
		fmt.Printf("Error converting file: %v\r\n", err)
	}

	if !watch {
		return
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("Error creating watcher: %v\r\n", err)
		os.Exit(1)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			// watch for events
			case <-watcher.Events:
				err := convertFile(inputFile, outputFile)
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
