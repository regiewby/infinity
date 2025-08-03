package config

type AppConfig struct {
	Name        string       `json:"name"`
	ServiceName string       `json:"serviceName"`
	Env         string       `json:"env"`
	Host        string       `json:"host"`
	Port        int          `json:"port"`
	OwnerInfo   OwnerInfo    `json:"ownerInfo"`
	Data        DataConfig   `json:"data"`
	RedisConfig RedisConfig  `json:"redisConfig"`
	Trace       TraceConfig  `json:"trace"`
	Logger      LoggerConfig `json:"logger"`
}

type OwnerInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	URL   string `json:"url"`
}

type DataConfig struct {
	MySQL MySQLConfig `json:"mysql"`
}

type MySQLConfig struct {
	Master               DBConfig       `json:"master"`
	Slave                DBConfig       `json:"slave"`
	MasterCircuitBreaker CircuitBreaker `json:"masterCircuitBreaker"`
	SlaveCircuitBreaker  CircuitBreaker `json:"slaveCircuitBreaker"`
}

type DBConfig struct {
	DSN            string `json:"dsn"`
	MaxIdle        int    `json:"maxIdle"`
	MaxOpen        int    `json:"maxOpen"`
	ConMaxLifeTime int    `json:"conMaxLifeTime"`
}

type CircuitBreaker struct {
	TimeoutsInMs           int `json:"timeoutsInMs"`
	MaxConcurrentReq       int `json:"maxConcurrentReq"`
	VolumePercentThreshold int `json:"volumePercentThreshold"`
}

type RedisConfig struct {
	Addr               string `json:"addr"`
	IdleTimeoutInSec   int    `json:"idleTimeoutInSec"`
	PoolSize           int    `json:"poolSize"`
	ReadOnlyFromSlaves bool   `json:"readOnlyFromSlaves"`
	ReadTimeoutInSec   int    `json:"readTimeoutInSec"`
	WriteTimeoutInSec  int    `json:"writeTimeoutInSec"`
	TLSEnabled         bool   `json:"tlsEnabled"`
}

type TraceConfig struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Disable bool   `json:"disable"`
}

type LoggerConfig struct {
	WorkerCount     int    `json:"workerCount"`
	BufferSize      int    `json:"bufferSize"`
	LogLevel        int    `json:"logLevel"`
	StacktraceLevel int    `json:"stacktraceLevel"`
	LogFormat       string `json:"logFormat"`
}

type KafkaConfig struct {
	Brokers             []string `json:"brokers"`
	GroupID             string   `json:"groupID"`
	TopicName           string   `json:"topicName"`
	ClusterType         string   `json:"clusterType"`
	EnableTLS           bool     `json:"enableTLS"`
	InitOffset          string   `json:"initOffset"`
	ClientID            string   `json:"clientID"`
	EnableAsyncAcks     bool     `json:"enableAsyncAcks"`
	SyncProd            bool     `json:"syncProd"`
	RequiredAcks        int16    `json:"requiredAcks"`
	MaxAttempts         int      `json:"maxAttempts"`
	RetryIntervalMillis int      `json:"retryIntervalMillis"`
}
