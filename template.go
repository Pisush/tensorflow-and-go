package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

const (
	inputImage = "./captcha/3KC6YY.png"
)

func captchaToText(path string, savedModel *tf.SavedModel, printLogs bool) string {
	// Read captcha
	captchaImage, err := ioutil.ReadFile(inputImage)
	if err != nil {
		log.Fatal(err)
	}

	// This is where the action happens: run captcha through tensorflow model!

	// Define the Output struct that will be fed in the model:
	// the operation here is providing the input image as bytes of the prediction node, the index is 0
	feedsOutput := tf.Output{
		//...
	}

	// Define the input Tensor: the string versino of the CAPTCHA
	//...

	// Set a map from the operation we will apply to the input it will be applied on
	//	feeds := map[tf.Output]*tf.Tensor{feedsOutput: feedsTensor}

	// Define the Output struct that will be fetched as the model's result:
	// the operation here is the output of the prediction node, the index is 0
	//...

	// Run the data through the graph using the two Output structs
	// and receive the output: the captcha prediction
	//...

	// We only need the first element of the tensor
	var text string
	//...

	if printLogs {
		log.Println("for file" + path + "the text is" + text)
	}

	return text
}

func main() {
	printLogs := flag.Bool("printlog", false, "set to true for printing all log lines on the screen")
	flag.Parse()

	// Always make the log file run.log
	//...

	// Creates a new SavedModel by loading a model previously exported to a directory on disk
	// the tag is "serve", no special options

	var s captchaText
	// captchaText = captchaToText(inputImage, savedModel, *printLogs)
	fmt.Println(captchaText)
}
