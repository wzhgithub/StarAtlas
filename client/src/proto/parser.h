#ifndef _PROTO_PARSER_INCLUDED_
#define _PROTO_PARSER_INCLUDED_

#include <memory>

#include "task.h"
#include "device.h"

#include "rapidjson/istreamwrapper.h"
#include "rapidjson/schema.h"
#include "rapidjson/stringbuffer.h"

inline int parseDefault(rapidjson::Document& _document, vector<Device>& _device) {return 0;}

bool parse(const char* filename, rapidjson::Document& _document);
int parseSwitch(rapidjson::Document& _document, vector<Device>& _device);
int parseRemote(rapidjson::Document& _document, vector<Device>& _device);
int parseXpu(rapidjson::Value& _document, vector<Device>& _device, uint8_t typ, uint8_t baseIndex, uint8_t globalDev, bool _use_idx = false);
int parsePartition(rapidjson::Document& _document, std::shared_ptr< vector<Partition> > _parts);

#endif
