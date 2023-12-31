package constant

import "net/url"

var BaseUrl, _ = url.Parse("/")

const GetAllLoansRoute = "/loan/all"
const GetLoanRoute = "/loan/{loanID}"
const ApplyRoute = "/apply"
const GetUserRoute = "/user/{username}"
const CreateUserRoute = "/register"

const CredScoreClientBaseUrl = "http://localhost:3000/"
const PgConnectionString = "postgresql://pgadmin:admin@127.0.0.1:5432/loandb?sslmode=disable"
