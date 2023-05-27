## Generate file .go
protoc --go_out=. calculator/calculatorpb/calculator.proto 

protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
proto/*.proto