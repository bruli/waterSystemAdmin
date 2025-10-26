FROM golang:1.25.2

RUN mkdir /app
WORKDIR /app
RUN apt-get update && apt-get upgrade -y && apt-get install -y make git
RUN go install -v github.com/cespare/reflex@latest
EXPOSE 3000
ENTRYPOINT ["reflex", "-c", "./reflex.conf"]
