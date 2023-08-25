package tests

type ConfigEnv struct {
	DBUser string `envconfig:"DB_USER" required:"true"`
	DBPass string `envconfig:"DB_PASS" required:"true"`
	DBHost string `envconfig:"DB_HOST" required:"true"`
	DBPort string `envconfig:"DB_PORT" required:"true"`
	DBName string `envconfig:"DB_NAME" required:"true"`

	RabbitUser    string `envconfig:"R_USER" required:"false"`
	RabbitPass    string `envconfig:"R_PASS" required:"false"`
	RabbitHost    string `envconfig:"R_HOST" required:"false"`
	RabbitPort    string `envconfig:"R_PORT" required:"false"`
	RabbiExchange string `envconfig:"E_EXCHANGE" required:"true"`

	ServiceHost string `envconfig:"S_HOST" required:"true"`
	ServicePort string `envconfig:"S_PORT" required:"true"`
}
