############################
# STEP 1 build application
############################

FROM golang:alpine as builder

ARG USERNAME=pccrep
ARG USER_UID=1000
ARG USER_GID=$USER_UID

WORKDIR $GOPATH/src/pcc-reporter
COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -a -v -installsuffix cgo -ldflags="-s -X github.com/Locotech-Oy/prisma-cloud-compute-reporter/version.version={{.Version}}" -o /go/bin/pcc-reporter

# Create non-root user
RUN addgroup -g $USER_GID -S $USERNAME && \
  adduser -u $USER_UID -S $USERNAME -G $USERNAME

############################
# STEP 2 build a small image
############################

FROM scratch

ARG USERNAME=pccrep
ARG USER_UID=1000
ARG USER_GID=$USER_UID

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/pcc-reporter /go/bin/pcc-reporter

USER ${USER_UID}:${USER_GID}
ENTRYPOINT ["/go/bin/pcc-reporter"]