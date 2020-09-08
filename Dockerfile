FROM golang:1.15-buster as build
WORKDIR /src
ADD . /src
ENV CGO_ENABLED=0

#RUN go get github.com/markbates/pkger/cmd/pkger && \
#    pkger -include /web && \
#    go build -o /cluster-registry main.go pkged.go
RUN go build -o goproj

FROM gcr.io/distroless/static
COPY --from=build /goproj /
CMD ["/goproj"]
