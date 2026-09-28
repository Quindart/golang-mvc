# Tên ứng dụng và thư mục build
APP_NAME = app
BUILD_DIR = tmp

.PHONY: all build run watch clean tidy test

# Mặc định khi gõ lệnh 'make'
all: tidy build

# Cập nhật các gói phụ thuộc
tidy:
	go mod tidy

# Build ứng dụng
build:
	go build -o $(BUILD_DIR)/$(APP_NAME) .

# Chạy ứng dụng thông thường
run:
	air

# Xóa các file build tạm thời
clean:
	rm -rf $(BUILD_DIR)

# Chạy unit test
test:
	go test -v ./...