package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hanapedia/the-bench/the-bench-operator/pkg/object/factory"
	"github.com/hanapedia/the-bench/the-bench-operator/pkg/yaml"
)

func TestKafkaTopicFactory(t *testing.T) {
	args := factory.KafkaTopicArgs{
		Topic:       "test",
		ClusterName: "test",
		Namespace:   "test",
		Partitions:  1,
		Replicas:    1,
	}
	kafkaTopic := factory.KafkaTopicFactory(&args)

	// Generate the YAML
	kafkaTopicYAML := yaml.GenerateManifest(kafkaTopic)
	// Check if some of the fields are correctly present in the generated YAML
	if !strings.Contains(string(kafkaTopicYAML), "name: test") {
		t.Errorf("The 'name' field is incorrect or missing in the generated YAML")
	}

	if !strings.Contains(string(kafkaTopicYAML), "namespace: test") {
		t.Errorf("The 'namespace' field is incorrect or missing in the generated YAML")
	}

	fmt.Printf("%s", string(kafkaTopicYAML))
}
