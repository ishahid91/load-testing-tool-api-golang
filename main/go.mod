module load-test-tool/main

go 1.16

replace load-test-tool/services => ../services

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.1
	github.com/swaggo/swag v1.7.1
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20210813211128-0a44fdfbc16e // indirect
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d // indirect
	golang.org/x/sys v0.0.0-20210809222454-d867a43fc93e // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.5 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	load-test-tool/models v0.0.0-00010101000000-000000000000
	load-test-tool/services v0.0.0-00010101000000-000000000000
	load-test-tool/utilities v0.0.0-00010101000000-000000000000
)

replace load-test-tool/models => ../models

replace load-test-tool/utilities => ../utilities
