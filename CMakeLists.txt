cmake_minimum_required(VERSION 3.19.0)

project(hexcloud LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

add_subdirectory(hexlib)
add_subdirectory(apps)
