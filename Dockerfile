# build stage
FROM golang:1.11.1-stretch AS build-env
ENV GO111MODULE=on
COPY . /work
WORKDIR /work
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/pipi8 ./src
 
# final stage
FROM scratch
COPY --from=build-env /work/bin/pipi8 /work/bin/pipi8
ENV PATH="$PATH:/work/bin"