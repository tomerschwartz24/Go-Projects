###########BUILD THE APPLICATION########
FROM golang:latest as webapp-builder
WORKDIR /build-golang-webapp
COPY go.mod httpg.go  ./
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -o httpg
RUN mkdir /webapp
COPY webapp /webapp
########################################

#############RUN APPLICATION ON MINIMAL IMAGE#############
FROM scratch
COPY --from=webapp-builder  /build-golang-webapp/httpg /
COPY --from=webapp-builder /webapp /webapp
CMD ["./httpg"]
##########################################################
