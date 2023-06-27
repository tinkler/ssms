module github.com/tinkler/ssms/backend/service/store_house

go 1.20

replace github.com/tinkler/mqttadmin => ../../../../mqttadmin

replace github.com/tinkler/ssms/backend/service/proxy => ../proxy

require (
	github.com/go-chi/chi/v5 v5.0.8
	github.com/google/wire v0.5.0
	github.com/joho/godotenv v1.5.1
	github.com/tinkler/mqttadmin v0.0.0-00010101000000-000000000000
	gorm.io/gorm v1.25.1
)

require (
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.3.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/microsoft/go-mssqldb v1.1.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4 // indirect
	google.golang.org/grpc v1.55.0 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
	gorm.io/driver/postgres v1.5.0 // indirect
	gorm.io/driver/sqlserver v1.5.1 // indirect
)
