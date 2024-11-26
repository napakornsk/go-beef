module github.com/napakornsk/go-beef/beef

go 1.23.3

require (
	github.com/golang/protobuf v1.5.4
	github.com/napakornsk/go-beef/proto v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.68.0
)

require (
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241118233622-e639e219e697 // indirect
	google.golang.org/protobuf v1.35.2 // indirect
)

replace github.com/napakornsk/go-beef/proto => ./proto
