FROM golang:1.16-alpine AS build

WORKDIR /src
COPY . .
RUN go build -o /out/hexcloud .
FROM scratch AS bin
COPY --from=build /out/hexcloud /
CMD("/out/hexcloud")