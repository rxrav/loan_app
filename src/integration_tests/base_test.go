package integrationtests

import (
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"os"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/rxrav/loan_app/src/repo"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestMain(m *testing.M) {
	globalSetup()
	code := m.Run()
	globalTeardown()
	os.Exit(code)
}

var tcContext context.Context
var postgresContainer *postgres.PostgresContainer
var cs testcontainers.Container
var TestDb *sqlx.DB
var tcConnStr string
var csPort nat.Port
var err error

func globalSetup() {
	fmt.Println("global setup integration tests")
	tcContext = context.Background()

	dbName := "testloandb"
	dbUser := "pgtestadmin"
	dbPassword := "testadmin"

	log.Info().Msg("starting postgres")
	postgresContainer, err = postgres.RunContainer(
		tcContext,
		testcontainers.WithImage("docker.io/postgres:14-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		panic(err)
	}
	log.Info().Msg("postgres started")
	if tcConnStr, err = postgresContainer.ConnectionString(tcContext, "sslmode=disable"); err != nil {
		panic(err)
	} else {
		log.Info().Msg(fmt.Sprintf("conn str is %s", tcConnStr))
		TestDb = repo.GetDbInstance(tcConnStr)
		schemaScript, _ := os.ReadFile("../migration/schema.sql")
		_, err = TestDb.Exec(string(schemaScript))
		if err != nil {
			return
		}
		dataScript, _ := os.ReadFile("../migration/data.sql")
		_, err = TestDb.Exec(string(dataScript))
		if err != nil {
			return
		}
	}

	log.Info().Msg("starting credit score service")
	csReq := testcontainers.ContainerRequest{
		Image:        "di-course/credit-score-service",
		ExposedPorts: []string{"3000"},
		WaitingFor:   wait.ForLog("Credit score service listening on port 3000"),
	}

	cs, err = testcontainers.GenericContainer(tcContext, testcontainers.GenericContainerRequest{
		ContainerRequest: csReq,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}
	csPort, err = cs.MappedPort(tcContext, "3000")
	log.Info().Msg(fmt.Sprintf("credit score service available on port %v", csPort.Port()))
}

func globalTeardown() {
	fmt.Println("global teardown integration tests")
	if err := postgresContainer.Terminate(tcContext); err != nil {
		panic(err)
	}
	if err = cs.Terminate(tcContext); err != nil {
		panic(err)
	}
}
