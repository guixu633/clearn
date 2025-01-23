#include <iostream>
#include <unordered_map>

struct Student {
  int id;
  std::string name;
};

enum Color { RED, GREEN, BLUE };

std::string colorToString(Color color) {
  switch (color) {
  case RED:
    return "Red";
  case GREEN:
    return "Green";
  case BLUE:
    return "Blue";
  default:
    return "Unknown";
  }
}

int main() {
    int arr[5] = {1, 2, 3, 4, 5};
    for (int i = 0; i < 5; i++) {
      std::cout << arr[i] << std::endl;
    }

  std::cout << RED << std::endl;
  std::cout << colorToString(RED) << std::endl;
  std::cout << GREEN << std::endl;
  std::cout << colorToString(GREEN) << std::endl;
  std::cout << BLUE << std::endl;
  std::cout << colorToString(BLUE) << std::endl;

  Student stu = {22, "张三"};
  std::cout << stu.id << std::endl;
  std::cout << stu.name << std::endl;
  return 0;
}