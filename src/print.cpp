#include "include/print.hpp"
#include <iostream>
#include <string>

int print(const std::string &text, int times) {
  int maxTimes = 3;
  if (times <= 0) {
    return 0;
  }

  if (times > maxTimes) {
    times = maxTimes;
  }

  for (int i = 0; i < times; i++) {
    std::cout << text << std::endl;
  }

  return times;
}
