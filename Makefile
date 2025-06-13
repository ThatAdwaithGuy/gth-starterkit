db:
	dbmate migrate
	sqlc generate

generate:
	npm run build 
	templ generate 

build:
	make db 
	make generate


