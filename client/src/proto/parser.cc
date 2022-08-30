#include "parser.h"

#include <fstream>
#include <iostream>

bool parse(const char* filename, rapidjson::Document& _document) {
  std::ifstream ifs(filename);
  rapidjson::IStreamWrapper isw(ifs);
  _document.ParseStream(isw);
  if (_document.HasParseError()) {
    std::cerr << "ParseStream json error: "<<filename<<std::endl;
    return false;
  }
  return true;
}

int parseSwitch(rapidjson::Document& _document, vector<Device>& _device) {
  if (!_document.HasMember("central")) {
    std::cerr << "Invalid switch, do not have central switch." << std::endl;
    return 0;
  }

  const char* _key[] = {
    "central",
    "access",
  };
  uint8_t _type[] = {
    0,
    1,
  };

  int n_exch = 0;
  for (size_t h=0; h<sizeof(_key)/sizeof(const char*); h++) {
    const auto& _switches = _document[_key[h]].GetArray();
    for (const auto& _switch: _switches) {
      if (!_switch.IsObject()) {
        std::cerr << "Invalid switch: !_switch.IsObject(). " << std::endl;
        return 0;
      }

      if (!_switch.HasMember("name") ||
          !_switch.HasMember("index") ||
          !_switch.HasMember("connect_to")) {
        std::cerr << "Invalid switch: <name|index|connect_to> needed." << std::endl;
        return 0;
      }
      Device _dev;
      _dev.init(eEXCHNAGE, 
        uint8_t(_switch["index"].GetInt()),
        _switch["name"].GetString(),
        _type[h], 
        uint8_t(_switch["connect_to"].GetInt())
      );
      _device.emplace_back(std::move(_dev));
      n_exch++;
    }
  }
  return n_exch;
}

int parseRemote(rapidjson::Document& _document, vector<Device>& _device) {
  if (!_document.IsArray()) {
    std::cerr << "Invalid remote." << std::endl;
    return 0;
  }
  
  const auto& _clusters = _document.GetArray();
  int n_remote = 0;
  for (const auto& _cluster: _clusters) {
    if (!_cluster.IsObject()) {
      std::cerr << "Invalid remote: !_cluster.IsObject(). " << std::endl;
      return 0;
    }

    if (!_cluster.HasMember("connect_to") ||
        !_cluster.HasMember("list") ||
        !_cluster["list"].IsArray()) {
      std::cerr << "Invalid remote: <connect_to|list> needed." << std::endl;
      return 0;
    }
    
    uint8_t _connect_to = uint8_t(_cluster["connect_to"].GetInt());
    const auto& _items = _cluster["list"].GetArray();
    for (const auto& _item: _items) {
      if (!_item.HasMember("name")||
          !_item.HasMember("index")||
          !_item.HasMember("type")) {
        std::cerr << "Invalid remote: <name|index|type> need." << std::endl;
        return 0;
      }
      
      Device _dev;
      _dev.init(eREMOTE, uint8_t(_item["index"].GetInt()),
        _item["name"].GetString(),
        uint8_t(_item["type"].GetInt()), _connect_to);
      _device.emplace_back(std::move(_dev));
      n_remote++;
    }
  }
  return n_remote;
}

int parseXpu(rapidjson::Value& _document, vector<Device>& _device, uint8_t typ, uint8_t baseIndex, uint8_t globalDev, bool _use_idx) {
  if (!_document.IsArray()) {
    std::cerr << "Invalid xpu." << std::endl;
    return 0;
  }
  
  const auto& _items = _document.GetArray();
  int n_xpu = 0;
  for (const auto& _item: _items) {
    if (!_item.IsObject()) {
      std::cerr << "Invalid xpu: !_item.IsObject(). " << std::endl;
      return 0;
    }

    if (!_item.HasMember("name") ||
        // !_item.HasMember("index") ||
        !_item.HasMember("type")) {
      std::cerr << "Invalid xpu: <name|type> needed." << std::endl;
      return 0;
    }

    int _dev_index = baseIndex + (_item.HasMember("index")?
      uint8_t(_item["index"].GetInt()):
    (_device.size()-globalDev+1));
    if (_use_idx && _item.HasMember("index")) {
      _dev_index = _item["index"].GetInt();
    }
    const char* _dev_name = _item["name"].GetString();
    const char* _auto_name = "@auto@";
    if (strcmp(_dev_name, _auto_name)==0) {
      _dev_name = nullptr;
    }
     
    Device _dev;
    _dev.init(typ, _dev_index, _dev_name, 
      uint8_t(_item["type"].GetInt()), 0);

    if (typ == eFPGA) {
      _device.emplace_back(std::move(_dev));
      n_xpu++;
      continue;
    }

    uint8_t core = _item.HasMember("core")?_item["core"].GetInt():1;
    uint16_t inops = _item.HasMember("inops")?_item["inops"].GetInt():32768;
    uint16_t flops = _item.HasMember("flops")?_item["flops"].GetInt():32768;
    uint16_t mem = _item.HasMember("mem")?_item["mem"].GetInt():32768;
    //std::cout << "core: "<< core <<std::endl;
    //std::cout << "inops: "<< inops <<std::endl;
    //std::cout << "flops: "<< flops <<std::endl;
    //std::cout << "mem: "<< mem <<std::endl;
    _dev.setBasic(core, inops, flops, mem);
    _device.emplace_back(std::move(_dev));
    n_xpu++;
  }
  return n_xpu;
}

int parsePartition(rapidjson::Document& _document, std::shared_ptr< vector<Partition> > _parts) {
 if (!_document.IsArray()) {
    std::cerr << "Invalid partition." << std::endl;
    return 0;
  }
  
  const auto& _clusters = _document.GetArray();
  int n_part = 0;
  for (const auto& _cluster: _clusters) {
    if (!_cluster.IsObject()) {
      std::cerr << "Invalid partition: !_cluster.IsObject()" << std::endl;
      return 0;
    }

    if (!_cluster.HasMember("name") ||
        !_cluster.HasMember("index") ||
        !_cluster.HasMember("vmc") ||
        !_cluster.HasMember("device") ||
        !_cluster.HasMember("tasks")) {
      std::cerr << "Invalid partition: <name|index|vmc|tasks> needed." << std::endl;
      return 0;
    }

    Partition _part;
    _part.init(_cluster["index"].GetInt(),
      _cluster["vmc"].GetInt(),
      _cluster["device"].GetInt(),
      _cluster["name"].GetString());

    if (!_part.parseTask(_cluster)) {
      std::cerr << "Invalid partition: invalied task." << std::endl;
      return 0;
    }

    _parts->emplace_back(std::move(_part));
    n_part++;
  }
  return n_part;
}
