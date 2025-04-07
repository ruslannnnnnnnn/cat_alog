package main

import (
	"cat_alog/internal/infrastructure/cassandra"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// example:       <container_name>                        <migration file name>
// docker exec -it cat_alog-server-1 /bin/migrator migrate "00001 create keyspace.cql"
func main() {
	rootCmd := &cobra.Command{
		Use:   "migrator",
		Short: "database migration tool",
	}

	var migrate = &cobra.Command{
		Use:   "migrate",
		Short: "migrate",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			migrationFile := args[0]
			err := cassandra.ApplyMigration(migrationFile)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	rootCmd.AddCommand(migrate)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
