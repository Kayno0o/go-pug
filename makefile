jade-build: jade-clean
	jade -writer -pkg=jade -d=app/jade client/pages/*.pug

jade-clean:
	rm -rf app/jade/*

go-build:
	go build -o ./dist/server main.go

go-prod: go-build
	./dist/server

go-dev: jade
	go run main.go

tailwind:
	npx tailwindcss -i ./client/css/main.css -o ./public/style/main.css

unocss:
	pnpm unocss

unocss-minify:
	pnpm unocss -m

prod: jade-build unocss-minify go-prod

watch:
	./watcher.sh
