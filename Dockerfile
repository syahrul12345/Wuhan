# Stage 1 build the go binary
FROM golang:alpine as stage1
WORKDIR /app
COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download
COPY ./backend .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

#Stage 2 build the website
FROM node:12 as stage2
WORKDIR /app
COPY ./website/package.json ./
RUN npm install
COPY ./website .
RUN npm run build

# Third & Final stage, take on the go binary and react build folder to be served.
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Copy current cache
COPY --from=stage1 /app/db ./db
COPY --from=stage1 /app/main .
COPY --from=stage2 /app/build ./build
EXPOSE 8000
CMD ["./main"]