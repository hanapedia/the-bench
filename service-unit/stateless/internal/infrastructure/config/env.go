package config

import (
	"os"
	"sync"
)

type EnvVars struct {
	DEP_ENV                 string
	HTTP_PORT               string
	GRPC_PORT               string
	KAFKA_PORT              string
	KAFKA_CLUSTER_NAME      string
	KAFKA_CLUSTER_NAMESPACE string
	MONGO_PORT              string
	POSTGRE_PORT            string
}

var envVars *EnvVars
var once sync.Once

func GetEnvs() *EnvVars {
	once.Do(func() {
		envVars = loadEnvVariables()
	})
	return envVars
}

func loadEnvVariables() *EnvVars {
	depEnv := "k8s"
	httpPort := "8080"
	grpcPort := "9090"
	kafkaPort := "9092"
	kafkaClusterName := "my-cluster"
	kafkaClusterNamespace := "kafka"
	mongoPort := "27017"
	postgrePort := "5432"

	if envDepEnv, ok := os.LookupEnv("DEP_ENV"); ok {
		depEnv = envDepEnv
	}

	if envHttpPort, ok := os.LookupEnv("HTTP_PORT"); ok {
		httpPort = envHttpPort
	}

	if envGrpcPort, ok := os.LookupEnv("GRPC_PORT"); ok {
		grpcPort = envGrpcPort
	}

	if envKafkaPort, ok := os.LookupEnv("KAFKA_PORT"); ok {
		kafkaPort = envKafkaPort
	}

	if envKafkaClusterName, ok := os.LookupEnv("KAFKA_PORT"); ok {
		kafkaClusterName = envKafkaClusterName
	}

	if envKafkaClusterNamespace, ok := os.LookupEnv("KAFKA_PORT"); ok {
		kafkaClusterNamespace = envKafkaClusterNamespace
	}

	if envmongoPort, ok := os.LookupEnv("MONGO_PORT"); ok {
		mongoPort = envmongoPort
	}

	if envPostgrePort, ok := os.LookupEnv("POSTGRE_PORT"); ok {
		postgrePort = envPostgrePort
	}

	return &EnvVars{
		DEP_ENV:                 depEnv,
		HTTP_PORT:               httpPort,
		GRPC_PORT:               grpcPort,
		KAFKA_PORT:              kafkaPort,
		KAFKA_CLUSTER_NAME:      kafkaClusterName,
		KAFKA_CLUSTER_NAMESPACE: kafkaClusterNamespace,
		MONGO_PORT:              mongoPort,
		POSTGRE_PORT:            postgrePort,
	}
}
