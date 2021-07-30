find_package(Protobuf CONFIG REQUIRED)
find_package(gRPC CONFIG REQUIRED)

set(PROTO_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}/proto)
set(PROTO_FILE ${PROTO_DIRECTORY}/hexagon.proto)
file(MAKE_DIRECTORY ${PROTO_DIRECTORY})

message("Downloading hexagon.proto file from 3vilM33pl3/hexcom")
set(HEXCOM_PROTO_URL https://raw.githubusercontent.com/3vilM33pl3/hexcom/master/lib/protos/hexagon.proto)
file(DOWNLOAD ${HEXCOM_PROTO_URL} ${PROTO_FILE} SHOW_PROGRESS)

add_library(hexlib
        ${PROTO_FILE}
        HexagonService.cpp
        HexagonClient.cpp)

target_link_libraries(hexlib
        PUBLIC
        protobuf::libprotobuf
        gRPC::grpc
        gRPC::grpc++
        )
target_include_directories(hexlib PUBLIC ${CMAKE_CURRENT_BINARY_DIR} ../include)

get_target_property(grpc_cpp_plugin_location gRPC::grpc_cpp_plugin LOCATION)
protobuf_generate(TARGET hexlib LANGUAGE cpp)
protobuf_generate(TARGET hexlib LANGUAGE grpc GENERATE_EXTENSIONS .grpc.pb.h .grpc.pb.cc PLUGIN "protoc-gen-grpc=${grpc_cpp_plugin_location}")
