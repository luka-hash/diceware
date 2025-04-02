FROM golang:1.24-alpine as build
ENV CGO_ENABLED=0
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o /bin/diceware ./cmd/server/

FROM scratch
ENV DICEWARE_PORT=8000
COPY --from=build /bin/diceware /bin/diceware
EXPOSE ${DICEWARE_PORT}
CMD ["/bin/diceware"]
