FROM golang:1.19-bullseye
WORKDIR /work

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        wget \
        make \
        unzip \
        git \
        clang-format \
        vim \
    && apt-get clean

# cobraのCLIをインストール
RUN go install github.com/spf13/cobra-cli@latest

RUN pwd
COPY src/go.mod src/go.sum ./
RUN go mod download && go mod verify
