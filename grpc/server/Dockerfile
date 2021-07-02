FROM golang
ENV GOPATH=
COPY . .
RUN go build -o main main.go
EXPOSE 8080
CMD ./main