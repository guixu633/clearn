#include <iostream>
#include <limits>
#include <string>

int main() {
  // 整数类型
  std::cout << "===== 整数类型 =====" << std::endl;
  int i = 42;
  short s = 100;
  long l = 123456789L;
  long long ll = 123456789012345LL;

  std::cout << "int 大小: " << sizeof(int) << " 字节" << std::endl;
  std::cout << "short 大小: " << sizeof(short) << " 字节" << std::endl;
  std::cout << "long 大小: " << sizeof(long) << " 字节" << std::endl;
  std::cout << "long long 大小: " << sizeof(long long) << " 字节" << std::endl;

  // 浮点类型
  std::cout << "\n===== 浮点类型 =====" << std::endl;
  float f = 3.14159f;
  double d = 3.14159265359;

  std::cout << "float 大小: " << sizeof(float) << " 字节" << std::endl;
  std::cout << "double 大小: " << sizeof(double) << " 字节" << std::endl;

  // 字符类型
  std::cout << "\n===== 字符类型 =====" << std::endl;
  char c = 'A';
  wchar_t wc = L'世';

  std::cout << "char 大小: " << sizeof(char) << " 字节" << std::endl;
  std::cout << "wchar_t 大小: " << sizeof(wchar_t) << " 字节" << std::endl;

  // 布尔类型
  std::cout << "\n===== 布尔类型 =====" << std::endl;
  bool b1 = true;
  bool b2 = false;

  std::cout << "bool 大小: " << sizeof(bool) << " 字节" << std::endl;
  std::cout << "bool true 值: " << b1 << std::endl;
  std::cout << "bool false 值: " << b2 << std::endl;

  // 类型的最大最小值
  std::cout << "\n===== 类型范围 =====" << std::endl;
  std::cout << "int 最小值: " << std::numeric_limits<int>::min() << std::endl;
  std::cout << "int 最大值: " << std::numeric_limits<int>::max() << std::endl;
  std::cout << "unsigned int 最大值: "
            << std::numeric_limits<unsigned int>::max() << std::endl;

  // 字符串类型
  std::cout << "\n===== 字符串类型 =====" << std::endl;
  std::string str = "Hello, 世界!";
  std::cout << "字符串内容: " << str << std::endl;
  std::cout << "字符串长度: " << str.length() << std::endl;

  return 0;
}
