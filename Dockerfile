FROM golang:latest

RUN mkdir /autocomplete-test
ADD . /autocomplete-test
RUN cd /autocomplete-test && git clone https://github.com/markelmad/autocomplete-test-emapta.git
WORKDIR /autocomplete-test
RUN go build -o main .
CMD ["/autocomplete-test/main"]