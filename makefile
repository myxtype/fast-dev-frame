.PHONY: build clean zip all

CMDs := rest job worker admin pushing

build-linux:
	@for var in $(CMDs); do \
  		echo "building $${var}"; \
  		GOOS=linux go build -o "./build/$${var}" "./cmd/$${var}"; \
    done

build:
	@for var in $(CMDs); do \
  		echo "building $${var}"; \
  		go build -o "./build/$${var}" "./cmd/$${var}"; \
    done

clean:
	@for var in $(CMDs); do \
		if [ -f ./build/$$var ] ; then rm ./build/$$var* ; fi \
    done

	@if [ -f ./build/cmd.zip ] ; then rm ./build/cmd.zip ; fi

zip:
	@for var in $(CMDs); do \
		zip ./build/cmd.zip -jg ./build/$$var; \
	done

all:
	make clean
	make build-linux
	make zip

help:
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make clean - 移除二进制文件"
