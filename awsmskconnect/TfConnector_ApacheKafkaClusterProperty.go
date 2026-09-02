package awsmskconnect


// Experimental.
type TfConnector_ApacheKafkaClusterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#bootstrap_servers TfConnector#bootstrap_servers}.
	// Experimental.
	BootstrapServers *string `field:"required" json:"bootstrapServers" yaml:"bootstrapServers"`
	// vpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mskconnect_connector#vpc TfConnector#vpc}
	// Experimental.
	Vpc *TfConnector_VpcProperty `field:"required" json:"vpc" yaml:"vpc"`
}

