FROM alpine:latest as builder

WORKDIR /openserver

RUN apk add --no-cache libc6-compat
RUN apk add --no-cache tzdata

FROM scratch

COPY --from=builder /lib/ld-musl-x86_64.so.1 /lib/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY output/ .

ENV TZ=Asia/Shanghai
