package conf

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Конфиги приложения
type Config struct {
	Server struct {
		Port   string `yaml:"port"`
		Host   string `yaml:"host"`
		Prefix string `yaml:"prefix"`
	} `yaml:"server"`
	TemplatesDatabase struct {
		Host                string `yaml:"host"`
		Port                string `yaml:"port"`
		DBName              string `yaml:"db_name"`
		Username            string `yaml:"user"`
		Password            string `yaml:"password"`
		TemplatesCollection string `yaml:"templates_collection"`
	} `yaml:"templates_database"`
	TasksQueueDatabase struct {
		Host                 string `yaml:"host"`
		Port                 string `yaml:"port"`
		DBName               string `yaml:"db_name"`
		Username             string `yaml:"user"`
		Password             string `yaml:"password"`
		TasksQueueCollection string `yaml:"tasks_queue_collection"`
	} `yaml:"tasks_queue_database"`
}

func New() *Config {
	f, err := os.Open("./config.yaml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var cfg Config

	decoder := yaml.NewDecoder(f)

	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}

func (c *Config) ServerAddr() string {
	return fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port)
}

func (c *Config) TemplatesDatabaseURI() string {
	return fmt.Sprintf("%s:%s", c.TemplatesDatabase.Host, c.TemplatesDatabase.Port)
}

func (c *Config) TemplatesCollection() string {
	return c.TemplatesDatabase.TemplatesCollection
}

func (c *Config) TemplatesDBName() string {
	return c.TemplatesDatabase.DBName
}

func (c *Config) TasksQueueDatabaseURI() string {
	return fmt.Sprintf("%s:%s", c.TasksQueueDatabase.Host, c.TasksQueueDatabase.Port)
}

func (c *Config) TasksQueueCollection() string {
	return c.TasksQueueDatabase.TasksQueueCollection
}

func (c *Config) TasksQueueDBName() string {
	return c.TemplatesDatabase.DBName
}
