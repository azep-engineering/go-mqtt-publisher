FROM golang AS BUILD

WORKDIR /src
ADD go.mod go.sum main.go /src/

RUN go get
RUN go build -ldflags "-linkmode external -extldflags -static" -o go-mqtt-publisher

FROM scratch

COPY --from=build /src/go-mqtt-publisher /
ENV SERIAL_DEVICE /dev/ttyUSB0

CMD ["/go-mqtt-publisher"]