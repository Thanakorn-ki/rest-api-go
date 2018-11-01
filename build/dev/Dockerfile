FROM golang:1.11-alpine3.8 AS builder 

WORKDIR $GOPATH/src/github.com/salapao2136/rest-api-go

COPY . .
# ====================================================== 
# (1) Install dep dependency tool 
# - debug tool(not in use now) 
# ====================================================== 
RUN apk add --no-cache git
RUN go get github.com/derekparker/delve/cmd/dlv
RUN go get -u github.com/TonPC64/gomon
CMD ["gomon", "main.go"," ", "."]