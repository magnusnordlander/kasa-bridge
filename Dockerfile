FROM golang:1.13 
RUN mkdir /app 
ADD go.mod main.go /app/
WORKDIR /app 
RUN go mod download
RUN go build -o main . 

FROM nginx:1.17
RUN mkdir /app 
ADD nginx.conf run.sh /app/
WORKDIR /app
COPY --from=0 /app/main .
RUN apt-get update &&\
    apt-get install -y ffmpeg

CMD ["/app/run.sh"]
