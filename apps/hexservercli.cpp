#include <iostream>
#include <boost/program_options.hpp>
#include <hexlib/HexagonService.hpp>

namespace po = boost::program_options;

using namespace std;

void RunServer(const string ServerAddress) {
    HexagonService Service;
    ServerBuilder Builder;

    Builder.AddListeningPort(ServerAddress, grpc::InsecureServerCredentials());
    Builder.RegisterService(&Service);

    std::unique_ptr<Server> server(Builder.BuildAndStart());
    std::cout << "Server listening on port: " << ServerAddress << ":" << std::endl;

    server->Wait();
}


int main(int ac, char** av) {
    string ServerAddress;
    uint16_t ServerPort;
    po::options_description desc("Hexagon server options");
    desc.add_options()
            ("help", "help message")
            ("address", po::value<string>(&ServerAddress)->default_value("0.0.0.0"), "address to listen on [ip]")
            ("port", po::value<uint16_t>(&ServerPort)->default_value(50051), "port listening on [port], can also be set with PORT environment variable");


    ServerAddress = ServerAddress + ":" + to_string(ServerPort);

    po::variables_map vm;
    po::store(po::parse_command_line(ac, av, desc), vm);
    po::notify(vm);

    if (vm.count("help")) {
        std::cout << desc << "\n";
        return 1;
    }

    // Run server
    std::cout << "Running server on " << ServerAddress << ":" << ServerPort << std::endl;
    RunServer(ServerAddress);

    return 0;
}