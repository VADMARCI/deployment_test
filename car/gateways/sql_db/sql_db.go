package sql_db

import (
	"car/ent"
	"car/ent/migrate"
	"context"

	sql2 "database/sql"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"

	"log"
	"os"
	"strconv"
)

type SqlDb struct {
	EntClient   *ent.Client
	writeDriver *sql.Driver
	readDriver  *sql.Driver
}

func (sqlDb *SqlDb) getFullHost(dbHost string) string {
	return fmt.Sprint(os.Getenv(dbHost), "/", os.Getenv("DATABASE"), "?charset=utf8mb4&parseTime=True&loc=UTC")
}

func (sqlDb *SqlDb) Connect() (err error) {
	writeDbHost := sqlDb.getFullHost("WRITE_DB_HOST")
	readDbHost := sqlDb.getFullHost("READ_DB_HOST")
	sqlDb.EntClient, err = sqlDb.connect(writeDbHost, readDbHost)
	sqlDb.debugHandler()
	return err
}

func (sqlDb *SqlDb) createDatabase() (err error) {
	db, err := sql.Open("mysql", fmt.Sprint(os.Getenv("WRITE_DB_HOST"), "/"))
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database connected successfully")
	}
	_, err = db.DB().Exec(fmt.Sprintf("CREATE DATABASE  IF NOT EXISTS  %s", os.Getenv("DATABASE")))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created database..")
	}
	return err
}

func (sqlDb *SqlDb) Migrate() (err error) {
	sqlDb.createDatabase()
	writeDbHost := sqlDb.getFullHost("WRITE_DB_HOST")
	readDbHost := sqlDb.getFullHost("READ_DB_HOST")
	sqlDb.EntClient, err = sqlDb.connect(writeDbHost, readDbHost)
	ctx := context.Background()
	err = sqlDb.EntClient.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Println("DB migration error", err)
		panic(fmt.Errorf("migration error", err.Error()))
	}
	return err
}

func (sqlDb *SqlDb) ConnectTest() (err error) {
	dbHost := os.Getenv("DB_TEST_HOST")
	_, err = sqlDb.connect(dbHost, dbHost)
	sqlDb.debugHandler()
	return err
}

func (sqlDb *SqlDb) ExecSql(sql string) {
	sqlDb.writeDriver.DB().Exec(sql)
}

func (sqlDb *SqlDb) DB() *sql2.DB {
	return sqlDb.readDriver.DB()
}

func (sqlDb *SqlDb) Ping() error {
	err := sqlDb.readDriver.DB().Ping()
	if err != nil {
		return err
	}
	err = sqlDb.writeDriver.DB().Ping()
	if err != nil {
		return err
	}
	return nil
}

func (sqlDb *SqlDb) debugHandler() {
	debugLogger, err := strconv.ParseBool(os.Getenv("DB_LOGGER_ENABLED"))
	if err != nil {
		debugLogger = true
	}
	if debugLogger {
		sqlDb.EntClient = sqlDb.EntClient.Debug()
	}
}

func (sqlDb *SqlDb) connect(writeDbHost string, readDbHost string) (client *ent.Client, err error) {
	client, err = sqlDb.getClient(writeDbHost, readDbHost)
	if err != nil {
		return nil, err
	}
	return client, nil

}

func (sqlDb *SqlDb) getClient(writeDbHost string, readDbHost string) (*ent.Client, error) {
	dbIdleConnections := 5
	dbMaxConnection := 10

	if len(os.Getenv("DB_MAX_CONNECTION")) != 0 {
		con, _ := strconv.Atoi(os.Getenv("DB_MAX_CONNECTION"))
		dbMaxConnection = con
	}

	if len(os.Getenv("DB_MAX_IDLE_CONNECTION")) != 0 {
		con, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTION"))
		dbIdleConnections = con
	}

	log.Println("DB_HOST", writeDbHost)

	wd, err := sql.Open("mysql", writeDbHost)
	if err != nil {
		return nil, err
	}
	rd, err := sql.Open("mysql", readDbHost)
	if err != nil {
		return nil, err
	}
	sqlDb.writeDriver = wd
	sqlDb.readDriver = rd
	sqlDb.configureConnection(wd, dbIdleConnections, dbMaxConnection)
	sqlDb.configureConnection(rd, dbIdleConnections, dbMaxConnection)
	client := ent.NewClient(ent.Driver(&multiDriver{w: wd, r: rd}))
	client.Intercept(
		sqlDb.updatedAtNilInterceptor(),
		sqlDb.softDeleteInterceptor(),
	)
	return client, nil
}

func (sqlDb *SqlDb) softDeleteInterceptor() ent.Interceptor {
	return ent.InterceptFunc(func(next ent.Querier) ent.Querier {
		return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {

			values, err := next.Query(ctx, query)
			if err != nil {
				return nil, err
			}
			return values, nil
		})
	})
}

func (sqlDb *SqlDb) updatedAtNilInterceptor() ent.Interceptor {
	return ent.InterceptFunc(func(next ent.Querier) ent.Querier {
		return ent.QuerierFunc(func(ctx context.Context, query ent.Query) (ent.Value, error) {
			values, err := next.Query(ctx, query)
			if err != nil {
				return nil, err
			}
			// These models have updated_at fields  []

			return values, nil
		})
	})
}

func (sqlDb *SqlDb) configureConnection(driver *sql.Driver, dbIdleConnections int, dbMaxConnection int) {
	// Get the underlying sql.DB object of the driver.
	db := driver.DB()
	db.SetMaxIdleConns(dbIdleConnections)
	db.SetMaxOpenConns(dbMaxConnection)
	db.SetConnMaxLifetime(time.Hour)
}

func (sqlDb *SqlDb) Close() {
	sqlDb.EntClient.Close()
}
