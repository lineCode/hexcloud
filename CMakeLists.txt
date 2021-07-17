cmake_minimum_required(VERSION 3.15)
project(hexcloud)

set(CMAKE_CXX_STANDARD 20)

set(CMAKE_SUPPRESS_DEVELOPER_WARNINGS 1 CACHE INTERNAL "No dev warnings")

# Cmake find modules
list(APPEND CMAKE_MODULE_PATH "${CMAKE_CURRENT_LIST_DIR}/cmake")

include(FetchContent)

set(Boost_DEBUG FALSE)
set(BUILD_TESTING OFF)
set(Boost_USE_MULTITHREADED ON)

IF (WIN32)
    set(BOOST_ROOT C:/local/boost_1_69_0)
    set(Boost_NO_SYSTEM_PATHS ON)
    set(Boost_USE_STATIC_LIBS ON)
    set(Boost_LIB_VERSION vc120)
    # ASM compile fails for 64 bit builds
    set(OPENSSL_NO_ASM ON)
ENDIF()

IF(MSVC)
    SET(CMAKE_CXX_FLAGS "/Zc:__cplusplus /EHsc")
ENDIF()

find_package(Boost REQUIRED COMPONENTS program_options)

find_package(Protobuf)
IF(NOT Protobuf_FOUND)
    FetchContent_Declare(
            protobuf
            GIT_REPOSITORY https://github.com/protocolbuffers/protobuf
            GIT_TAG        v3.10.0
    )
ENDIF()

set(gRPC_BUILD_GRPC_CPP_PLUGIN ON)
set(gRPC_BUILD_GRPC_CSHARP_PLUGIN OFF)
set(gRPC_BUILD_GRPC_NODE_PLUGIN OFF)
set(gRPC_BUILD_GRPC_OBJECTIVE_C_PLUGIN OFF)
set(gRPC_BUILD_GRPC_PHP_PLUGIN OFF)
set(gRPC_BUILD_GRPC_PYTHON_PLUGIN OFF)
set(gRPC_BUILD_GRPC_RUBY_PLUGIN  OFF)
set(gRPC)

set(gRPC_USE_PROTO_LITE OFF)

find_package(GRPC)
IF(NOT GRPC_FOUND)
    FetchContent_Declare(
            grpc
            GIT_REPOSITORY https://github.com/grpc/grpc.git
            GIT_TAG        v1.38.1)
    FetchContent_MakeAvailable(grpc)
ENDIF()

FetchContent_Declare(
        cli
        GIT_REPOSITORY  https://github.com/daniele77/cli.git
        GIT_TAG         v1.1.1
)
FetchContent_MakeAvailable(cli)

add_subdirectory(hexlib)
add_subdirectory(apps)
