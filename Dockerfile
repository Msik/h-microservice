FROM golang:1.21

WORKDIR /app

COPY . ./

# CMD ["/app"]

# RUN apt update

# RUN apt install -y git
# RUN apt install -y ca-certificates

# RUN apt install -y protobuf-compiler

# COPY . /path-to-project

# RUN make build
# RUN make run

RUN curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
    sh

EXPOSE 8080
EXPOSE 8082
CMD ["make", "run"]
