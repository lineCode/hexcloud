FROM alpine:3.11.5 AS build

RUN apk upgrade && \
    apk --update --no-cache add \
        gcc \
        g++ \
        build-base \
        autoconf \
        automake \
        libtool \
        cmake \
        libstdc++ \
        git \
        boost-dev \
        boost-program_options \
        bash \
        vim \
        zlib \
        zlib-dev \
        c-ares \
        c-ares-dev \
        grpc \
        grpc-dev \
        grpc-cli \
        libunwind

RUN mkdir -p ./hexlib
RUN wget https://raw.githubusercontent.com/3vilM33pl3/hexcom/master/lib/protos/hexagon.proto
COPY . /src
WORKDIR /src
RUN cmake .
RUN make -j16
RUN make install

FROM alpine:3.11

RUN apk upgrade && \
    apk --update --no-cache add \
        libstdc++ \
        boost-program_options \
        zlib \
        c-ares \
        grpc

COPY --from=build /src/apps/hexgrpc_server ./
CMD ["./hexgrpc_server"]