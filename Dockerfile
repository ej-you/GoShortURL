FROM golang:alpine3.20

# install make util for Makefile
RUN apk add make

# set up workdir
RUN mkdir -p ./github.com/Danil-114195722/GoShortURL
WORKDIR /go/github.com/Danil-114195722/GoShortURL

# copy project files to container
COPY . .

RUN make compile

CMD ["make", "prod"]
