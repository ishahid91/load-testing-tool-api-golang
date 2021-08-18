module load-test-tool/services

go 1.16

replace load-test-tool/models => ../models

require (
	github.com/stretchr/testify v1.7.0
	load-test-tool/models v0.0.0-00010101000000-000000000000
	load-test-tool/utilities v0.0.0-00010101000000-000000000000
)

replace load-test-tool/utilities => ../utilities
