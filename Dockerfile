FROM golang:1-stretch

# install tensorflow
RUN curl -L "https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-linux-x86_64-1.12.0.tar.gz" | tar -C /usr/local -xz 
RUN ldconfig

# go get tensorflow lib
RUN go get -d github.com/tensorflow/tensorflow/tensorflow/go

# python lib and tools like `saved_model_cli`
RUN curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py
RUN python get-pip.py
RUN pip install tensorflow