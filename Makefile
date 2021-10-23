PROJECT_NAME=$(shell basename "$(PWD)")
SCRIPT_AUTHOR=Andrey Kapitonov <andrey.kapitonov.96@gmail.com>
SCRIPT_VERSION=0.0.1

api:
	./bin/build_api.sh

monitoring:
	./bin/build_monitoring.sh

repository:
	./bin/build_repository.sh

storage:
	./bin/build_storage.sh

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