package awsmsk


// Experimental.
type TfServerlessCluster_ClientAuthenticationProperty struct {
	// sasl block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_serverless_cluster#sasl TfServerlessCluster#sasl}
	// Experimental.
	Sasl *TfServerlessCluster_SaslProperty `field:"required" json:"sasl" yaml:"sasl"`
}

