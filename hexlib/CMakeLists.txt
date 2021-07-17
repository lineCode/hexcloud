set(HEADER_LIST "${hexgrpc_SOURCE_DIR}/include/hexlib/hexgrpc.hpp")


IF(NOT hexagon.proto)
    # Proto file
    # Get file from hexcom project (single source of truth
    set(HEXCOM_PROTO_URL https://raw.githubusercontent.com/3vilM33pl3/hexcom/master/lib/protos/hexagon.proto SHOW_PROGRESS)
    set(PROTO_FILE_LOCAL ${CMAKE_CURRENT_BINARY_DIR}/hexagon.proto)
    file(DOWNLOAD ${HEXCOM_PROTO_URL} ${PROTO_FILE_LOCAL})
ELSEIF()
    set(PROTO_FILE_LOCAL ${CMAKE_CURRENT_SOURCE_DIR}/hexagon.proto)
ENDIF()

get_filename_component(PROTO_FILE ${PROTO_FILE_LOCAL} ABSOLUTE)
get_filename_component(PROTO_PATH "${PROTO_FILE_LOCAL}" PATH)

IF(GRPC_FOUND)
    set(PROTOC_EXECUTABLE /usr/bin/protoc)
    set(GRPC_CPP_PLUGIN_EXECUTABLE /usr/bin/grpc_cpp_plugin)
ELSE()
    set(PROTOBUF_LIBRARY libprotobuf)
    set(GRPC_GRPCPP_UNSECURE_LIBRARY grpc++_unsecure)

    set(PROTOC_EXECUTABLE $<TARGET_FILE:protoc>)
    set(GRPC_CPP_PLUGIN_EXECUTABLE $<TARGET_FILE:grpc_cpp_plugin>)
ENDIF()

# Generate sources
# Generated sources
set(PROTOBUF_SOURCE "${CMAKE_CURRENT_BINARY_DIR}/hexagon.pb.cc")
set(PROTOBUF_HEADER "${CMAKE_CURRENT_BINARY_DIR}/hexagon.pb.h")
set(GRPC_SOURCE "${CMAKE_CURRENT_BINARY_DIR}/hexagon.grpc.pb.cc")
set(GRPC_HEADER "${CMAKE_CURRENT_BINARY_DIR}/hexagon.grpc.pb.h")
add_custom_command(
        OUTPUT "${PROTOBUF_SOURCE}" "${PROTOBUF_HEADER}" "${GRPC_SOURCE}" "${GRPC_HEADER}"
        COMMAND ${PROTOC_EXECUTABLE}
        ARGS --grpc_out="${CMAKE_CURRENT_BINARY_DIR}"
        --cpp_out="${CMAKE_CURRENT_BINARY_DIR}"
        -I "${PROTO_PATH}"
        --plugin=protoc-gen-grpc="${GRPC_CPP_PLUGIN_EXECUTABLE}"
        "${PROTO_FILE}"
        DEPENDS "${PROTO_FILE}")

include_directories("${CMAKE_CURRENT_BINARY_DIR}")

add_library(hex_library STATIC HexagonService.cpp HexagonClient.cpp ${PROTOBUF_SOURCE} ${GRPC_SOURCE})
target_compile_features(hex_library PRIVATE cxx_std_17)

target_include_directories(hex_library PUBLIC ../include)
target_include_directories(hex_library PUBLIC "${CMAKE_CURRENT_BINARY_DIR}")


IF(GRPC_FOUND)
    target_link_libraries(hex_library
            gRPC::grpc
            gRPC::grpc++
            gRPC::grpc++_reflection
            protobuf::libprotobuf)
ELSE()

    target_link_libraries(hex_library
            ${GRPC_GRPCPP_UNSECURE_LIBRARY}
            ${PROTOBUF_LIBRARY})
ENDIF()




install(TARGETS hex_library LIBRARY)