In order to run migrations you need to install golang migrate CLI:

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

Running migrations:
migrate -path ./migrations -database mysql://{{user}}:{{password}}@tcp/{{dbname}} {{up/down}} {{n}}
