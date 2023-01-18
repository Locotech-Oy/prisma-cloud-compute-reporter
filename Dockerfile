############################
# STEP 1 build application
############################

FROM golang:alpine as builder
# RUN apk update && \
#   apk add --no-cache curl git xz && \
#   curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR $GOPATH/src/pcc-reporter
COPY ./ ./

#RUN dep ensure
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -a -v -installsuffix cgo -ldflags="-s -X github.com/Locotech-Oy/prisma-cloud-compute-reporter/version.version={{.Version}}" -o /go/bin/pcc-reporter

############################
# STEP 2 build a small image
############################

FROM scratch
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/pcc-reporter /go/bin/pcc-reporter
ENTRYPOINT ["/go/bin/pcc-reporter"]