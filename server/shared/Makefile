errno:
	kitex -module github.com/xince-fun/InstaGo ./idl/base/errno.thrift

generate:
	kitex -module github.com/xince-fun/InstaGo ./../services/$(service)/idl/$(service).thrift

user:
	make service=user generate

blob:
	make service=blob generate

relation:
	make service=relation generate

post:
	make service=post generate

all:
	make user
	make blob
	make follow
	make relation
	make post