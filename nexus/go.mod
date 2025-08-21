module github.com/HierIsPierre/nexus

go 1.24.6

require (
	github.com/HierIsPierre/nexus/docker v0.0.0-20250821130254-411e5ac3d3ca
	github.com/HierIsPierre/nexus/opsportal v0.0.0-20250821130254-411e5ac3d3ca
)

require (
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
	google.golang.org/grpc v1.75.0 // indirect
	google.golang.org/protobuf v1.36.8 // indirect
)

replace github.com/HierIsPierre/nexus/docker => ../docker

replace github.com/HierIsPierre/nexus/opsportal => ../opsportal
