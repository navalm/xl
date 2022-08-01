FROM golang:1.18.4-alpine3.15@sha256:665e676bf8e54bc130ec6c45abeb075754db8759adea82587a739ad81e9a6380

WORKDIR /src/xl

COPY api/main.go ./
COPY go.mod ./
# include checksums
# COPY go.sum ./

RUN go build -o ./main && chmod +x ./main

EXPOSE 8080

CMD ["./main"]