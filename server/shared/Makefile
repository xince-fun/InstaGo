errno:
	kitex -module github.com/xince-fun/InstaGo ./idl/base/errno.thrift

generate:
	kitex -module github.com/xince-fun/InstaGo ./../services/$(service)/idl/$(service).thrift

user:
	make service=user generate

blob:
	make service=blob generate

all:
	make user