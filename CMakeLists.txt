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

find_package(GRPC)
IF(NOT GRPC_FOUND)
    FetchContent_Declare(
            grpc
            GIT_REPOSITORY https://github.com/3vilM33pl3/grpc
            GIT_TAG        master)
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