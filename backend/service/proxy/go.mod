module github.com/tinkler/ssms/backend/service/proxy

go 1.20

replace github.com/tinkler/mqttadmin => ../../../../mqttadmin

replace github.com/sashabaranov/go-openai => ../../../../go-openai

require (
	github.com/joho/godotenv v1.5.1
	github.com/pkg/errors v0.9.1
	github.com/sashabaranov/go-openai v1.10.2-0.20230613203226-3f4e3bb0ca25
	github.com/tinkler/mqttadmin v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.55.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/go-chi/chi/v5 v5.0.8 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4 // indirect
)
