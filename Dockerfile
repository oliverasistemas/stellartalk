FROM golang:latest

WORKDIR /app
COPY . .

RUN go install ./cmd/stellartalkd

RUN bash ./init.sh

EXPOSE 8080
CMD ["stellartalkd", "start"]