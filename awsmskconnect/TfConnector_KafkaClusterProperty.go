package awsmskconnect


// Experimental.
type TfConnector_KafkaClusterProperty struct {
	// apache_kafka_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#apache_kafka_cluster TfConnector#apache_kafka_cluster}
	// Experimental.
	ApacheKafkaCluster *TfConnector_ApacheKafkaClusterProperty `field:"required" json:"apacheKafkaCluster" yaml:"apacheKafkaCluster"`
}

