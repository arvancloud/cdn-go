##################################### Build #####################################

FROM --platform=$BUILDPLATFORM golang:1.20-alpine as builder

ARG APP_VERSION="undefined@docker"

WORKDIR /src

COPY cmd cmd
COPY pkg pkg
COPY go.* .

ENV LDFLAGS="-s -w -X github.com/arvancloud/cdn-go/internal/pkg/version.version=$APP_VERSION"
ENV GO111MODULE=on
ARG TARGETOS TARGETARCH
ENV GOOS $TARGETOS
ENV GOARCH $TARGETARCH

RUN set -x \
    && go version \
    && CGO_ENABLED=0 go build -trimpath -ldflags "$LDFLAGS" -o /src/cdn-uncompress cmd/*.go

##################################### Compression #####################################

FROM hatamiarash7/upx:latest as upx

COPY --from=builder /src /

RUN upx --best --lzma -o /cdn /cdn-uncompress

######################################## Final ########################################

FROM scratch

ARG APP_VERSION="undefined@docker"
ARG DATE_CREATED

LABEL \
    org.opencontainers.image.title="cdn" \
    org.opencontainers.image.description="Docker image for ArvanCloud CDN API " \
    org.opencontainers.image.url="https://github.com/arvancloud/cdn-go" \
    org.opencontainers.image.source="https://github.com/arvancloud/cdn-go" \
    org.opencontainers.image.vendor="arvancloud" \
    org.opencontainers.image.author="arvancloud" \
    org.opencontainers.version="$APP_VERSION" \
    org.opencontainers.image.created="$DATE_CREATED" \
    org.opencontainers.image.licenses="MIT"

COPY --from=upx /cdn /cdn

ENTRYPOINT ["/cdn"]
