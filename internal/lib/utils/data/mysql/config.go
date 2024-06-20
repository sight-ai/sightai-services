package mysql

// Config defines a cluster of mysql dbs, master and slave
type Config struct {
	Debug  bool
	Master *ConnectionConfig
	Slave  *ConnectionConfig
}

// ConnectionConfig defines configuration of a db connection
type ConnectionConfig struct {
	Dsn         string
	MaxIdle     int
	MaxOpen     int
	MaxLifeTime int
	Name        string // we use this name for logging and reporting, no space allowed
}
