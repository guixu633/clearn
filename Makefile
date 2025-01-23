CXX = clang++          # 编译器：使用 Clang
CXXFLAGS = -std=c++17 -I.  # 添加 -I. 让编译器从当前目录开始查找头文件
TARGET = cmd/hello         # 生成的可执行文件名称
SRCS = src/main.cpp src/print.cpp  # 源文件列表

all: $(TARGET)         # 默认目标

$(TARGET): $(SRCS)
	$(CXX) $(CXXFLAGS) $^ -o $@  # 编译命令

clean:
	rm -f $(TARGET)    # 清理生成的文件