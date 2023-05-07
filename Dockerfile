FROM golang:1.19-bullseye as builder

ADD . /go/urakil
WORKDIR /go/urakil
RUN make clean && make && adduser --disabled-login --disabled-password nonroot

FROM scratch

COPY --from=builder /go/urakil/urakil /usr/bin/urakil
COPY --from=builder /etc/passwd /etc/passwd
USER nonroot

ENTRYPOINT [ "/usr/bin/urakil" ]