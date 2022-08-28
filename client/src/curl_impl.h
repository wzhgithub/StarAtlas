/// Copyleft (c) 2020 Tencent, Inc. All rights reserved.
///
/// @brief access ps by curl
/// @author robertxiong@tencent.com
/// @date 2020/12/24 11:16:36
/// @file curl_ps_impl.h

#ifndef DIVINATION_BASE_CURL_PS_IMPL_H_
#define DIVINATION_BASE_CURL_PS_IMPL_H_

#include <stdio.h>
#include <string>
#include "curl/curl.h"

class CurlPSImpl {
 public:
  CurlPSImpl() : _curl(nullptr) {}
  CurlPSImpl(uint32_t max_timeout, bool verbose) 
    : _curl(nullptr)
    , _max_timeout(max_timeout)
    , _verbose(verbose) {}

  ~CurlPSImpl() {
  }

  int Initialize(const std::string& url) {
    curl_global_init(CURL_GLOBAL_ALL);
    _curl = curl_easy_init();
    if (_curl == nullptr) {
      return -1;
    }

    curl_easy_setopt(_curl, CURLOPT_URL, url.c_str());
    return 0;
  }

  void Clean() {
    curl_easy_cleanup(_curl);
    curl_global_cleanup();
  }

  int Execute(std::string& response) {
    if (_curl == nullptr) {
      return -1;
    }

    if (_verbose) {
      curl_easy_setopt(_curl, CURLOPT_VERBOSE, 1L);                // verbose
    }
    curl_easy_setopt(_curl, CURLOPT_NOSIGNAL, (long)1);            // nosignal
    curl_easy_setopt(_curl, CURLOPT_TIMEOUT_MS, _max_timeout);  // 100ms
    curl_easy_setopt(_curl, CURLOPT_WRITEDATA, &response);
    curl_easy_setopt(_curl, CURLOPT_WRITEFUNCTION, write_data);
    
    CURLcode res;
    res = curl_easy_perform(_curl);
    if (res != CURLE_OK) {
      fprintf(stderr, "curl_easy_perform() failed: %s\n", curl_easy_strerror(res));
      return -1;
    }
    
    return 0;
  }

  static int write_data(void* buffer, size_t sz, size_t nmemb, void* ResInfo) {
    std::string* psResponse = (std::string*)ResInfo;
    psResponse->append((char*)buffer, sz * nmemb);
    return sz * nmemb;
  }

 private:
  CURL* _curl;
  uint32_t _max_timeout = 1000; // ms 
  bool _verbose = true;
};

#endif // DIVINATION_BASE_CURL_PS_IMPL_H_
