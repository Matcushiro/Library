package config

import "errors"

func GetPortFromConfig(config map[string]string) (string, error) {
	value, exists := config["PORT"]
	if exists != true {
		return "", errors.New("Ключ PORT в конфигурации отсутствует")
	}
	return value, nil
}
