package awsmsk


// Experimental.
type AwsMskServerlessCluster_ClientAuthenticationProperty struct {
	// sasl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_serverless_cluster#sasl AwsMskServerlessCluster#sasl}
	// Experimental.
	Sasl *AwsMskServerlessCluster_SaslProperty `field:"required" json:"sasl" yaml:"sasl"`
}

