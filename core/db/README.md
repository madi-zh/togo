### DB migrations

Initially, i wanted to reuse golang-migrate/ pkg to write my own migration tool, but then I realized that i'm steering from my original tasks backend app.
So, for now, all the migrations will be run using the same [golang-migrate](https://github.com/golang-migrate/migrate) pkg.

#### create migration
```bash $ migrate create -ext sql -dir ./migrations -seq create_tasks_table```

### migrate
```bash $ migrate -source file://./migrations -database "postgres://test:test@localhost:5430/database?sslmode=disable" up```
the sslmode=disable due to local postgres instance (it's simpler for now, but will change)
