package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DSN string `json:"dsn"`
}

func Setup() (string, error) {
	configFile := "config.json"
	var config Config

	// Verificar si el archivo de configuración existe
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		// Si el archivo no existe, créalo con el DSN
		config = Config{
			DSN: "root:mysql@tcp(127.0.0.1:3306)/ferry",
		}

		// Convertir el struct de configuración a JSON
		configData, err := json.MarshalIndent(config, "", "    ")
		if err != nil {
			return "", fmt.Errorf("error al convertir la configuración a JSON: %v", err)
		}

		// Escribir JSON en el archivo usando os.WriteFile
		err = os.WriteFile(configFile, configData, 0644)
		if err != nil {
			return "", fmt.Errorf("error al escribir el archivo de configuración: %v", err)
		}

		fmt.Println("Archivo de configuración creado:", configFile)
	} else {
		// Si el archivo existe, léelo y analízalo
		configData, err := os.ReadFile(configFile)
		if err != nil {
			return "", fmt.Errorf("error al leer el archivo de configuración: %v", err)
		}

		err = json.Unmarshal(configData, &config)
		if err != nil {
			return "", fmt.Errorf("error al analizar el JSON de configuración: %v", err)
		}

		fmt.Println("Archivo de configuración encontrado:", configFile)
	}

	return config.DSN, nil
}
