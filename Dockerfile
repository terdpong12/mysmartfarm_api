FROM golang:1.13.1

# add the source
COPY . /go/src/github.com/MySmartFarm/mysmartfarm_api
WORKDIR /go/src/github.com/MySmartFarm/mysmartfarm_api/

# Install go dependencies
RUN go get github.com/gin-gonic/gin && \
    go get github.com/lib/pq && \
    go get github.com/influxdata/influxdb1-client && \
    go get github.com/fatih/structs && \
    go get github.com/dgrijalva/jwt-go

#build the go app
RUN go build -o app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./app"]