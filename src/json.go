package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Database config
type DBConfig struct {
	Adapter string `json:adapter`
	Dsn     string `json:dsn`
}

// Server config
type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	TLS  bool   `json:"tls"`
}

// Application config
type Config struct {
	Domain string       `json:"domain"`
	Server ServerConfig `json:"server"`
	DB     DBConfig     `json:"database"`
}

var devconfig = `{
    "domain" : "localhost",
    "server" : { "host":"127.0.0.1", "port":3333, "tls":false},
    "database" : {
        "adapter":"postgres",
        "dsn":"pgsql:host=127.0.0.1;port=5432;dbname=mydb;user=myuser;password=mypass"
        }
    }`

var AppConfig Config

func init() {
	fmt.Println("**** Initializing ****")
	if err := json.Unmarshal([]byte(devconfig), &AppConfig); err != nil {
		log.Fatalf("An error occured while unmarshalling json", err.Error())
	}
}

// This is a mini demonstration for cross JSON <> Struct
// conversions, indention and tagging structs.
func main() {
	fmt.Println("**** Application configuration struct ****")
	fmt.Printf("%v\n", AppConfig)
	fmt.Println("**** Alter struct and convert back to INDENTED JSON ****")
	// Alter the configuration
	AppConfig.Server.Port = 4444
	config, _ := json.MarshalIndent(AppConfig, "", "  ")
	fmt.Println(string(config))
}
