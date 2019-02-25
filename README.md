# tensorflow-and-go
This is an example of using the Go API to read CAPTCHAs with TensorFlow. 
Log is generated for each run in the file run.log.  
For printing the logs, run with flag -printlog=true

### Structure
- The model in a protobuf version can be found in the `tensorflow_savedmodel_captcha` folder
- - Several CAPTCHAs that can be used for serving can be found in the `captcha` folder, alongside the Python script used to generate them
- The template with instructions in the comments is `tempalte.go`
- And, well, the solutions are in the `solution` folder. There is a standard one and the detailed one, if you want to read through the full solution with the detailed instructions.
