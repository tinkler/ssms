$backendLocation = Get-Location
$protoInclude = "F:\Program\protoc\include"
Set-Location ../../../frontend/flutter/smart_inventory

protoc.exe --dart_out=grpc:lib/src/generated -I $protoInclude -I $backendLocation"/api/proto" smart/v1/smart.proto
protoc.exe --dart_out=grpc:lib/src/generated -I $protoInclude google/protobuf/any.proto
Set-Location $backendLocation