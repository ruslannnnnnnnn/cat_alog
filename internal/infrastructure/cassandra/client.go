package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/spf13/viper" // viper for .yaml files
)

const configPath = "/app/cassandra_config.yaml"
const keyspace = "catalog"

type Config struct {
	Hosts       []string
	ClusterName string
	Keyspace    string
	Datacenter  string
	Rack        string
	Consistency gocql.Consistency
	NumTokens   int
	Snitch      string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(configPath)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %v", err)
	}

	cfg := &Config{
		Hosts:       viper.GetStringSlice("cassandra.hosts"),
		ClusterName: viper.GetString("cassandra.cluster_name"),
		Keyspace:    viper.GetString("cassandra.keyspace"),
		Datacenter:  viper.GetString("cassandra.datacenter"),
		Rack:        viper.GetString("cassandra.rack"),
		Consistency: gocql.ParseConsistency(viper.GetString("cassandra.consistency")),
		NumTokens:   viper.GetInt("cassandra.num_tokens"),
		Snitch:      viper.GetString("cassandra.snitch"),
	}

	return cfg, nil
}

func GetCassandraSession() (*gocql.Session, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	cluster := gocql.NewCluster(cfg.Hosts...)
	cluster.Keyspace = cfg.Keyspace
	cluster.Consistency = cfg.Consistency
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy(cfg.Datacenter)

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("connection failed: %v", err)
	}
	return session, nil
}
