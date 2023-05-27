FROM golang:1.19
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /get-weather-api
EXPOSE 8080
CMD ["/get-weather-api"]