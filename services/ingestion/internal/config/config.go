package config

type Config struct {
	Binance BinanceConfig
	Kafka   KafkaConfig
}

type BinanceConfig struct {
	WSURL   string
	Symbols []string
}

type KafkaConfig struct {
	Brokers []string
	Topic   string
}

func Load() (*Config, error) {
	// TODO: load from env
	return nil, nil
}
