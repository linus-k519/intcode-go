# Builing intcode computer
FROM golang:1.15-alpine AS build
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

# Building node server
FROM node:10-alpine as final
WORKDIR /usr/src/app
COPY --from=build /go/bin/intcode .
COPY server .
RUN npm install

EXPOSE 5000
CMD ["node", "server.js"]