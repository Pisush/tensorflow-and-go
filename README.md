# tensorflow-and-go
This is an example of using the Go API to read CAPTCHAs with TensorFlow. 

Log is generated for each run in the file run.log.  
For printing the logs, run with flag -printlog=true

### Structure
- `template.go` is the template with instructions in the comments
- The `tensorflow_savedmodel_captcha` folder contains the model in a protobuf format.
- The `captcha` folder has several CAPTCHAs that can be used for serving, alongside the Python script used to generate them.
- And, well, the solutions are in the `solution` folder. There is a standard one and the detailed one, if you want to read through the full solution with the detailed instructions.
