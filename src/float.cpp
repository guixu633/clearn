#include <iostream>

int main() {
  float num = 0.15625;
  float backup = num;
  bool flag = true;
  if (num >= 0) {
    flag = false;
  }
  int exp = 127;
  int tail[23] = {0};
  for (int i = 0; i < 23; i++) {
    if (num >= 1) {
      break;
    }
    num *= 2;
    exp--;
  }
  num -= 1;
  if (num < 1) {
    for (int i = 0; i < 127; i++) {
      num *= 2;
      if (num >= 1) {
        tail[i] = 1;
        num -= 1;
      }
      if (num == 0) {
        break;
      }
    }
  } else {
    for (int i = 0; i < 127; i++) {
      num /= 2;
      if (num >= 1) {
        tail[i] = 1;
        num -= 1;
      }
    }
  }

  std::cout << backup << std::endl;
  std::cout << flag << std::endl;
  std::cout << exp << std::endl;
  for (int i = 0; i < 23; i++) {
    std::cout << tail[i];
  }
  return 0;
}