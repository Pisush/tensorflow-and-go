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

	// Run captcha through the TensorFlow model
	feedsOutput := tf.Output{
		Op:    savedModel.Graph.Operation("CAPTCHA/input_image_as_bytes"),
		Index: 0,
	}
	feedsTensor, err := tf.NewTensor(string(captchaImage))
	if err != nil {
		log.Fatal(err)
	}
	feeds := map[tf.Output]*tf.Tensor{feedsOutput: feedsTensor}

	fetches := []tf.Output{
		{
			Op:    savedModel.Graph.Operation("CAPTCHA/prediction"),
			Index: 0,
		},
	}

	captchaText, err := savedModel.Session.Run(feeds, fetches, nil)
	if err != nil {
		log.Fatal(err)
	}

	text := captchaText[0].Value().(string)
	if printLogs {
		log.Println("for file" + path + "the text is" + text)
	}

	return text
}

func main() {
	printLogs := flag.Bool("printlog", false, "set to true for printing all log lines on the screen")
	flag.Parse()

	// Always make a log file
	logfile, err := os.OpenFile("run.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening a log file: %v", err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	// Load TensorFlow model
	savedModel, err := tf.LoadSavedModel("./tensorflow_savedmodel_captcha", []string{"serve"}, nil)
	if err != nil {
		log.Println("failed to load model", err)
		return
	}

	captchaText := captchaToText(inputImage, savedModel, *printLogs)
	fmt.Println(captchaText)
}
