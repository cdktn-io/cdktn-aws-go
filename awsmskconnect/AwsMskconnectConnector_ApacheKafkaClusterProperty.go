package awsmskconnect


// Experimental.
type AwsMskconnectConnector_ApacheKafkaClusterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#bootstrap_servers AwsMskconnectConnector#bootstrap_servers}.
	// Experimental.
	BootstrapServers *string `field:"required" json:"bootstrapServers" yaml:"bootstrapServers"`
	// vpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#vpc AwsMskconnectConnector#vpc}
	// Experimental.
	Vpc *AwsMskconnectConnector_VpcProperty `field:"required" json:"vpc" yaml:"vpc"`
}

