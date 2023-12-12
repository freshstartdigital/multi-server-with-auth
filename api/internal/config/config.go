package config

type Config struct {
    Port string
}

func Load() *Config {
    // For simplicity, hardcoding values. In real scenarios, use environment variables or a config file.
    return &Config{
        Port: "8080",
    }
}
