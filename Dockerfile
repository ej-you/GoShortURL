FROM golang:alpine3.20

# install make util for Makefile
RUN apk add make

# set up workdir
RUN cd /go/src
RUN mkdir -p ./github.com/Danil-114195722/GoShortURL
WORKDIR /go/src/github.com/Danil-114195722/GoShortURL

# install dependences
COPY ./go.mod .
COPY ./go.sum .
RUN go mod tidy

# copy project files to container
COPY . .

# compile app
RUN make compile
# run app
CMD ["make", "prod"]
