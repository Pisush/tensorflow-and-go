# tensorflow-and-go
This is an example of using the Go API to read CAPTCHAs with TensorFlow. 

Log is generated for each run in the file run.log.  
For printing the logs, run with flag -printlog=true

### Structure
- `template.go` is the template with instructions in the comments
- The `tensorflow_savedmodel_captcha` folder contains the model in a protobuf format.
- The `captcha` folder has several CAPTCHAs that can be used for serving, alongside the Python script used to generate them.
- `Dockerfile` is a docker container template will all the tooling installed (Go compiler and Tensorflow dependencies) and ready to use
- `Makefile` targets to easy built and run the docker container
- And, well, the solutions are in the `solution` folder. There is a standard one and the detailed one, if you want to read through the full solution with the detailed instructions.


## Develop From Docker

If you don't have Go or Tensorflow installed on your system you can use the Docker image that is ready to use with:

```
# Download and login into the docker container
$ make login

# you can run go or tools like 'saved_model_cli'
root@c7223ab43fbf:/go/src/github.com/Pisush/tensorflow-and-go# go --version

```



