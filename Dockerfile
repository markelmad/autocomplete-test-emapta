FROM golang:latest

RUN mkdir /autocomplete-test
WORKDIR /autocomplete-test

RUN export GO118MODULE=on 
RUN go get github.com/markelmad/autocomplete-test-emapta
RUN cd /autocomplete-test && git clone https://github.com/markelmad/autocomplete-test-emapta.git

RUN cd /autocomplete-test && go build

EXPOSE 9000

ENTRYPOINT ["/autocomplete-test/main"]