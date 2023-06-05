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

# sqlboilerコマンドをインストール
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

RUN pwd
COPY src/go.mod src/go.sum ./
RUN go mod download && go mod verify
