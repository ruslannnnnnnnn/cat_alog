package cassandra

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

const migrationsDir = "/app/migrations"

func GetMigrationFileContent(fileName string) string {
	var query []byte
	query, err := os.ReadFile(filepath.Join(migrationsDir, fileName))
	if err != nil {
		log.Fatal(err)
	}
	return string(query)
}

func ApplyMigration(fileName string) error {
	migrationString := GetMigrationFileContent(fileName)
	if migrationString == "" {
		return errors.New("no migration found")
	}

	session, err := GetCassandraSession()
	if err != nil {
		return err
	}
	defer session.Close()

	migrationQuery := session.Query(migrationString)
	err = migrationQuery.Exec()
	if err != nil {
		return err
	}
	return nil
}
