FROM golang:1.23
ENV PORT $PORT
WORKDIR weather-predictor/
COPY . .
RUN go build -o bin/server main.go
EXPOSE $PORT
CMD ["./bin/server"]