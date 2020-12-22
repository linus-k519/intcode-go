FROM golang:1.15 AS build
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

FROM python:3.8 as final
WORKDIR /usr/src/app
COPY --from=build /go/src/app/intcode .
COPY server .
RUN pip install -r requirements.txt

EXPOSE 5000
CMD ["python3.8", "server.py"]