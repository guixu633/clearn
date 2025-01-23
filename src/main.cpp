#include <iostream>

#include "include/print.hpp"

int main() {
  std::string input;
  int count = 0;

  while (count < 5) {
    std::cout << "请输入内容 (输入 exit 退出): ";
    std::getline(std::cin, input);

    if (input == "exit") {
      break;
    }

    print(input, 5);

    count++;
  }

  return 0;
}
