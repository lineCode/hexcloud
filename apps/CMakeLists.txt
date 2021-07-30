

find_package(Boost REQUIRED COMPONENTS program_options )

set(SOURCES hexservercli.cpp)

set(BOOST_INCLUDE_DIR F:/vcpkg/installed/x86-windows/include)

add_executable(hexgrpc_server ${SOURCES})
target_link_libraries(hexgrpc_server PRIVATE hexlib Boost::program_options)

