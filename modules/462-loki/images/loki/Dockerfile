ARG BASE_GOLANG_19_ALPINE
ARG BASE_DISTROLESS

FROM $BASE_GOLANG_19_ALPINE as artifact

ARG SOURCE_REPO
ENV SOURCE_REPO=${SOURCE_REPO}

ARG GOPROXY
ENV GOPROXY=${GOPROXY}

RUN apk add --no-cache make git
RUN git clone --depth 1 --branch v2.7.3 ${SOURCE_REPO}/grafana/loki.git /loki
WORKDIR /loki/
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o loki cmd/loki/main.go && \
    chown -R 64535:64535 /loki/ && \
    chmod 0700 /loki/loki

FROM $BASE_DISTROLESS
COPY --from=artifact /loki/loki /usr/bin
WORKDIR /loki/

EXPOSE 3100

ENTRYPOINT ["/usr/bin/loki"]
