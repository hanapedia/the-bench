package constants

const (
	ServiceNameIndex = 0
	ProtocolIndex    = 1
	ActionIndex      = 2
	AdapterNameIndex = 3
)

type StatelessAdapterVariant string
type BrokerAdapterVariant string
type StatefulAdapterVariant string
const (
	REST StatelessAdapterVariant = "rest"
	GRPC StatelessAdapterVariant = "grpc"

	KAFKA    BrokerAdapterVariant = "kafka"
	RABBITMQ BrokerAdapterVariant = "rabbitmq"
	Pulsar   BrokerAdapterVariant = "pulsar"

	MONGO   StatefulAdapterVariant = "mongo"
	POSTGRE StatefulAdapterVariant = "postgre"
)

type Action string
const (
	READ  Action = "read"
	WRITE Action = "write"
)

const (
	PayloadSize = 16
)

const (
	RestServerAddr  = ":8080"
	KafkaBrokerAddr = "kafka:9092"
	MongoURIAddr    = "mongodb://root:password@stateful-unit-mongo:27017/mongo?authSource=admin"
)

type RepositoryEntryVariant string
type RepositoryEntrySize int

const (
	SMALL      RepositoryEntryVariant = "small"  // 1kb entries
	MEDIUM     RepositoryEntryVariant = "medium" // 4kb entries
	LARGE      RepositoryEntryVariant = "large"  // 16kb entries
	SMALLSIZE  RepositoryEntrySize    = 1        // 1kb entries
	MEDIUMSIZE RepositoryEntrySize    = 4        // 1kb entries
	LARGESIZE  RepositoryEntrySize    = 16       // 1kb entries
)

const (
	NumInitialEntries = 100
)