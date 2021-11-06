PROJECT_NAME=$(shell basename "$(PWD)")
SCRIPT_AUTHOR=Andrey Kapitonov <andrey.kapitonov.96@gmail.com>
SCRIPT_VERSION=0.0.1
SERVICES=\
	storage

api:
	./bin/build_api.sh

monitoring:
	./bin/build_monitoring.sh

repository:
	./bin/build_repository.sh

storage:
	./bin/build_storage.sh

common_proto:
	protoc -I . \
    	--go_out=.  \
    	--go_opt=paths=source_relative \
    	--go-grpc_out=. \
    	--go-grpc_opt=paths=source_relative \
    	./src/protobuf/common/*.proto;

$(SERVICES):
	protoc -I ./src/protobuf/$@/ -I . \
		--go_out=./src/protobuf/$@/ \
		--go_opt=paths=source_relative \
		--go-grpc_out=./src/protobuf/$@/ \
		--go-grpc_opt=paths=source_relative \
		$@.proto;

default: common_proto $(SERVICES)

help:
	@echo -e "Usage: make [target] ...\n"
	@echo -e "build-migration-tool 	: Download & create migration tool"
	@echo -e "migrate-up 			: Apply migrations"
	@echo -e "migrate-down 		: Down migrations"
	@echo -e "migrate-status 		: Status of migrations"
	@echo -e "migrate-new 			: Create new migration by name={name_here}"
	@echo -e '\nProject name : '$(PROJECT_NAME)
	@echo -e "Written by $(SCRIPT_AUTHOR), version $(SCRIPT_VERSION)"
	@echo -e "Please report any bug or error to the author."