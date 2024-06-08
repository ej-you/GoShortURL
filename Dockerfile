# -----
# BUILD
# -----

FROM golang:1.22.4-alpine3.20 AS build

ENV PATH_DB="/root/db.sqlite3"

# tells the C compiler to enable support for large files, allowing the use
# of 64-bit versions of file handling functions such as pread64 and pwrite64
ENV CGO_CFLAGS="-D_LARGEFILE64_SOURCE"
ENV CGO_LDFLAGS="-D_LARGEFILE64_SOURCE"

# install make util for Makefile
RUN apk add --no-cache make
# install required packages for supporting CGo
RUN apk add --no-cache build-base sqlite-dev

# set up workdir
RUN cd /go/src
RUN mkdir -p ./github.com/Danil-114195722/GoShortURL
WORKDIR /go/src/github.com/Danil-114195722/GoShortURL

# install dependences
COPY ./go.mod .
COPY ./go.sum .
RUN go mod tidy
RUN go mod download

# copy project files to container
COPY . .

# compile app
RUN make compile

# ---
# RUN
# ---

FROM alpine:3.20 AS run

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
