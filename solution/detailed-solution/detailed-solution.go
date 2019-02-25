package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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
		Op:    savedModel.Graph.Operation("CAPTCHA/input_image_as_bytes"),
		Index: 0,
	}

	// Define the input Tensor: the string versino of the CAPTCHA
	feedsTensor, err := tf.NewTensor(string(captchaImage))
	if err != nil {
		log.Fatal(err)
	}

	// Set a map from the operation we will apply to the input it will be applied on
	feeds := map[tf.Output]*tf.Tensor{feedsOutput: feedsTensor}

	// Define the Output struct that will be fetched as the model's result:
	// the operation here is the output of the prediction node, the index is 0
	fetches := []tf.Output{
		{
			Op:    savedModel.Graph.Operation("CAPTCHA/prediction"),
			Index: 0,
		},
	}

	// Run the data through the graph and receive the output: the captcha prediction
	captchaText, err := savedModel.Session.Run(feeds, fetches, nil)
	if err != nil {
		log.Fatal(err)
	}

	// We only need the first element of the tensor
	var text string
	text = captchaText[0].Value().(string)

	if printLogs {
		log.Println("for file" + path + "the text is" + text)
	}

	return text
}

func main() {
	printLogs := flag.Bool("printlog", false, "set to true for printing all log lines on the screen")
	flag.Parse()

	// Always make the log file run.log
	logfile, err := os.OpenFile("run.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening a log file: %v", err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	// Creates a new SavedModel by loading a model previously exported to a directory on disk
	// the tag is "serve", no special options
	savedModel, err := tf.LoadSavedModel("./tensorflow_savedmodel_captcha", []string{"serve"}, nil)
	if err != nil {
		log.Println("failed to load model", err)
		return
	}

	captchaText := captchaToText(inputImage, savedModel, *printLogs)
	fmt.Println(captchaText)
}
