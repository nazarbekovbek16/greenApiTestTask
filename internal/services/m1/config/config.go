package config

import "github.com/caarlos0/env/v7"

type Config struct {
	URL           string `env:"URL" envDefault:"localhost"`
	Level         string `env:"APP_MODE" envDefault:"dev"`
	Port          string `env:"PORT" envDefault:":8080"`
	RabbitMQ      string `env:"RabbitMQURl" envDefault:"amqp://guest:guest@rabbitmq:5672/"`
	TaskQueueName string `env:"TaskQueueName" envDefault:"task_queue"`
}

func NewConfig() *Config {
	cfg := Config{}
	env.Parse(&cfg)
	return &cfg
}
