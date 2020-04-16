FROM golang
RUN CGO_ENABLED=0 go get -a -ldflags '-s' main.go
COPY hello .
ENTRYPOINT ["./hello"]
