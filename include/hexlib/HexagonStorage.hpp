#pragma once

#include <map>

class HexagonStorage {
public:
    std::map<std::pair<int,int>, int> AllTheHexagons;
    void ReadHexagonDataFromCloudStorage();


};

