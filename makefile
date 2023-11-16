.PHONY: build clean zip all

CMDs := rest job worker

build:
	@for var in $(CMDs); do \
  		echo "building $${var}"; \
  		GOOS=linux go build -o "./build/$${var}" "./cmd/$${var}"; \
    done

buildlocal:
	@for var in $(CMDs); do \
  		echo "building $${var}"; \
  		go build -o "./build/$${var}" "./cmd/$${var}"; \
    done

clean:
	@for var in $(CMDs); do \
		if [ -f ./build/$$var ] ; then rm ./build/$$var ; fi \
    done

zip:
	@for var in $(CMDs); do \
		zip ./build/cmd.zip -jg ./build/$$var; \
	done

all:
	make clean
	make build
	make zip

help:
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make clean - 移除二进制文件"
