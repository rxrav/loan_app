# Needed to add mockgen command on the command prompt or terminal

go install github.com/golang/mock/mockgen

# Needed to make gomock library available as project dependency

go get github.com/golang/mock/mockgen

# Needed to unit test loan approval service

mockgen -source .\src\client\credit_score_client.go -destination .\src\mocks\client\credit_score_client_mock.go

# Needed to unit test loan application service

mockgen -source .\src\repo\repo.go -destination .\src\mocks\repo\repo_mock.go

mockgen -source .\src\service\user_service.go -destination .\src\mocks\service\user_service_mock.go

mockgen -source .\src\service\loan_approval_service.go -destination .\src\mocks\service\loan_approval_service_mock.go

# run tests

go test -v ./...