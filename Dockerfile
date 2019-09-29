FROM golang:1.12.7

# add the source
COPY . /go/src/mysmartfarm_api
WORKDIR /go/src/mysmartfarm_api/

# Install go dependencies
RUN go get github.com/gin-gonic/gin && \
    go get github.com/lib/pq && \
    go get github.com/jinzhu/gorm

#build the go app
RUN go build -o app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./app"]