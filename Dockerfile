FROM golang AS build

WORKDIR /build
COPY Makefile Godeps.txt ./
RUN make deps

COPY ./ ./
WORKDIR /build/ecr-get-token
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /bin/ecr-get-token

FROM scratch

ENTRYPOINT [ "/bin/ecr-get-token" ]

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /bin/ecr-get-token /bin/ecr-get-token
