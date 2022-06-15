FROM golang:latest

RUN mkdir /autocomplete-test
ADD . /autocomplete-test
RUN cd /autocomplete-test && git clone https://github.com/markelmad/autocomplete-test-emapta.git
WORKDIR /autocomplete-test
RUN go build -o main .
CMD ["/autocomplete-test/main"]
# WORKDIR /autocomplete-test

# RUN export GO118MODULE=on 
# ADD . /autocomplete-test
# RUN cd /autocomplete-test && git clone https://github.com/markelmad/autocomplete-test-emapta.git

# RUN cd /autocomplete-test && go build

# EXPOSE 9000

# ENTRYPOINT ["/autocomplete-test/autocomplete-test-emapta"]