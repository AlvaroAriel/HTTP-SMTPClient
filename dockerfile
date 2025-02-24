FROM golang:latest as builder  

ARG CGO_ENABLED=0            
WORKDIR /app                 
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN go build -o server .

# Minimal image which runs binary
FROM scratch                  
COPY --from=builder /app/server /server
EXPOSE 8000
#Execute it
ENTRYPOINT ["/server"]
