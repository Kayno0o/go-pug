jade-clean:
	rm -rf src/jade/*

jade-build: jade-clean
	jade -writer -pkg=jade -d=src/jade /app/client/pages/*.pug

go-build:
	go build -o ./dist/server ./main.go
	./dist/server

prod: jade-build go-build

watch:
	./watcher.sh

docker:
ifeq ($(ENV),prod)
	make prod
else ifeq ($(ENV),dev)
	make watch
else
	@echo "Invalid environment specified. Use 'env=prod' or 'env=dev'."
	@exit 1
endif
