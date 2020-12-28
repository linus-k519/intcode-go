# Building intcode interpreter
FROM golang:1.15-alpine AS build
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

# Building node server
FROM node:15-slim as final
WORKDIR /usr/src/app
ENV PORT=6789
COPY --from=build /go/bin/intcode .
COPY server .
RUN npm install

# Run server
EXPOSE $PORT
CMD ["node", "server.js"]