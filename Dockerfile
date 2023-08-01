# syntax=docker/dockerfile:1

################################################################################
ARG GO_VERSION=1.20
FROM golang:${GO_VERSION} AS build
WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    CGO_ENABLED=0 go build -o /app/server ./cmd/api

################################################################################
FROM gcr.io/distroless/static-debian11 AS final

WORKDIR /app
COPY --from=build /app/server .

# we could use efs but this way we can leverage the docker layer cache more effectively
COPY ./ui/ ./ui/

ENTRYPOINT [ "/app/server" ]
