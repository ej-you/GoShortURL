#------
# BUILD
# -----

FROM golang:1.22.4 AS build

# for correct interaction beetwen Go and C. Espeshialy for "go-sqlite3" module
ENV CGO_ENABLED=1

# install make util for Makefile
RUN apt update -y
RUN apt install make
# install required packages for supporting CGo
# RUN apk add --no-cache gcc musl-dev gcompat musl-libc

# set up workdir
RUN cd /go/src
RUN mkdir -p ./github.com/Danil-114195722/GoShortURL
WORKDIR /go/src/github.com/Danil-114195722/GoShortURL

# install dependences
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

# copy project files to container
COPY . .

# compile app
RUN make compile

#----
# RUN
# ---

FROM alpine:3.20

# install make util for Makefile
RUN apk add --no-cache make

# make dir for logs
RUN mkdir /logs

WORKDIR /root
# copy compiled file and Makefile to run app
COPY --from=build /go/src/github.com/Danil-114195722/GoShortURL/Makefile .
COPY --from=build /go/src/github.com/Danil-114195722/GoShortURL/GoShortURL .

# run app
CMD ["make", "prod"]
