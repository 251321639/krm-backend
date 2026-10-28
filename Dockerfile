FROM alpine:3.18
COPY demo ./
ENTRYPOINT ["./demo"]