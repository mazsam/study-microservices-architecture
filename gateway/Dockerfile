FROM alpine:3.13 AS build

COPY --from=golang:1.18-alpine /usr/local/go/ /usr/local/go/

ENV PATH="/usr/local/go/bin:${PATH}"

COPY --from=quay.io/goswagger/swagger:0.30.4 /usr/bin/swagger /usr/bin/swagger

RUN apk add g++ && apk add make

WORKDIR /app

COPY . .

RUN make all

FROM alpine:3.13
WORKDIR /app
COPY --from=build /app/gateway-server .

EXPOSE 8080

CMD ["./gateway-server", "--port=8080", "--host=0.0.0.0"]

