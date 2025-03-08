db:
	dbmate migrate
	sqlc generate

generate:
	npm run build 
	templ generate 

gendb:
	make db 
	make generate


