

LOCAL_GOPATH=$(shell readlink -f ..)/_local/go
ROOT_PATH=$(shell readlink -f ../src/telegram-meetup)

export GOPATH:=$(LOCAL_GOPATH)

GODEPS=google.golang.org/appengine 		\
	   github.com/cnf/structhash 		\
	   github.com/onsi/gomega 			\
	   gopkg.in/telegram-bot-api.v4		\
	   github.com/glycerine/zebrapack   \
	   github.com/soh335/ical	 	    \
	   github.com/mattn/go-sqlite3      \
	   github.com/pborman/uuid          \
       github.com/aws/aws-sdk-go/aws \
	   github.com/skip2/go-qrcode \
	github.com/karalabe/xgo

env:
	env | egrep ^GO

test: getdeps
	cd $(ROOT_PATH) && go test ./...

getdeps: unmount mount
	GOPATH=$(LOCAL_GOPATH) go get -x -v $(GODEPS)

mount:
	mkdir -p $(LOCAL_GOPATH)/src
	mkdir -p $(LOCAL_GOPATH)/src/telegram-meetup
	bindfs  $(ROOT_PATH) $(LOCAL_GOPATH)/src/telegram-meetup

unmount:
	-fusermount -u  $(LOCAL_GOPATH)/src/telegram-meetup
	-rmdir $(LOCAL_GOPATH)/src/telegram-meetup

updatedeps: deps
	GOPATH=$(LOCAL_GOPATH) go get -u -x $(GODEPS)

generate:
	cd ../types && zebrapack -file dto.go -msgp

gaeserve: 
	cd $(ROOT_PATH)/host/gaeapp && goapp serve 

gaedeployv1: 
	cd $(ROOT_PATH)/host/gaeapp && goapp deploy --version v1 
	curl https://meetup-chatbot.appspot.com/init

arm:
	GOPATH=$(LOCAL_GOPATH) xgo -x -v -targets linux/arm ../_local/go/src/telegram-meetup/host/embed/cmd/bot	

clean:
	rm -vrf $(LOCAL_GOPATH)


profile:
	@echo "alias go=\"GOPATH=$(GOPATH) `which go`\"" > project_set_alias.rc
	@echo "#vim:syn=yes"                            >> project_set_alias.rc
	@echo ok
	@echo now run: 
	@echo "        source ./project_set_alias.rc"
	@echo
	@echo "export GOPATH=$(GOPATH)" > project_set_env.rc
	@echo "#vim:syn=yes"           >> project_set_env.rc
	@echo ok
	@echo now run: 
	@echo "        source ./project_set_env.rc"
