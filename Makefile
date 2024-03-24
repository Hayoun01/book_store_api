db_name = book_store
createdb:
	docker exec -it modest_raman createdb --username=postgres --owner=postgres $(db_name)
dropdb:
	docker exec -it modest_raman dropdb --username=postgres $(db_name)
migrateup:
	goose -dir='.\pkg\db\migration' postgres postgres://postgres:toor@localhost:5432/$(db_name) up
migratedown:
	goose -dir='.\pkg\db\migration' postgres postgres://postgres:toor@localhost:5432/$(db_name) down
sqlc:
	sqlc generate
.PHONY: createdb dropdb migrateup migratedown sqlc