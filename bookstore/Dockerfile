# FROM golang:1.19.3
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download 
# COPY *.go ./
# RUN go build -o /main
# CMD [ "/app/main" ]
# # ENTRYPOINT ["./main"]


FROM golang:1.19.3

RUN go version
ENV GOPATH=/

WORKDIR /project

COPY ./ ./

RUN go mod download
RUN go build -o main /project/app

CMD ["./main"]