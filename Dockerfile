FROM golang:1.19
WORKDIR /app
ARG ARCH=
ARG GOOS="linux"
ARG GOARCH="${ARCH}"
ARG CGO_ENABLED=0
ENV USER=limited
ENV UID=10001
RUN adduser \    
    --disabled-password \    
    --gecos "" \    
    --home "/nonexistent" \    
    --shell "/sbin/nologin" \    
    --no-create-home \    
    --uid "${UID}" \    
    "${USER}"
COPY . .
RUN go build -ldflags="-w -s" .

FROM scratch
WORKDIR /app
COPY --from=0 /etc/passwd /etc/passwd
COPY --from=0 /etc/group /etc/group
COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /app/dynamic-do-dns .
USER limited:limited
ENTRYPOINT ["/app/dynamic-do-dns"]
