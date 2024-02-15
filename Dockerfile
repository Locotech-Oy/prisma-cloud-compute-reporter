############################
# STEP 1 build application
############################

FROM golang:alpine as builder

ARG USERNAME=pccrep
ARG USER_UID=1000
ARG USER_GID=$USER_UID
ARG ORG_NAME=Locotech-Oy
ARG REPOSITORY_NAME=prisma-cloud-compute-reporter
ARG RELEASE_VERSION

# Verify that RELEASE_VERSION has been passed to build
RUN test -n ${RELEASE_VERSION:?}

WORKDIR $GOPATH/src/pcc-reporter

RUN wget -O prisma-cloud-compute-reporter_${RELEASE_VERSION}_linux_arm64.tar.gz https://github.com/${ORG_NAME}/${REPOSITORY_NAME}/releases/download/v${RELEASE_VERSION}/prisma-cloud-compute-reporter_${RELEASE_VERSION}_linux_arm64.tar.gz \
  && tar -xzf prisma-cloud-compute-reporter_${RELEASE_VERSION}_linux_arm64.tar.gz \
  && rm prisma-cloud-compute-reporter_${RELEASE_VERSION}_linux_arm64.tar.gz \
  && cp prisma-cloud-compute-reporter /go/bin/pcc-reporter

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