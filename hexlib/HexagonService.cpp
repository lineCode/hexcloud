#include <hexlib/HexagonService.hpp>
#include <hexlib/HexagonLibrary.hpp>

Status HexagonService::GetHexagonRing(::grpc::ServerContext *context, const ::hexagon::HexagonRingRequest *request,
                                      ::hexagon::HexCubeResponse *response) {
    auto hexpb = request->ha();
    int maxStep = 1;
    if(request->fill()) {
        maxStep = request->radius();
    }

    for(int step = 0; step <= maxStep; step++) {
        auto result = HexagonLibrary::Ring(Hexagon(hexpb.x(), hexpb.y(), hexpb.z()), request->radius() - step);
        std::cout << "Request: X: " << hexpb.x() << " Y: " << hexpb.y() << " Z: " << hexpb.z() << std::endl;

        google::protobuf::RepeatedPtrField <Hex> hexpbv;
        for (auto hex : result) {
            std::cout << "Response: Q: " << hex.Q << " R: " << hex.R << " S: " << hex.S << std::endl;
            auto hexpb = response->mutable_hc()->Add();
            hexpb->set_x(hex.Q);
            hexpb->set_y(hex.R);
            hexpb->set_z(hex.S);
        }
    }


    return Status::OK;
}
