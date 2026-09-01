package awsmsk


// Experimental.
type AwsMskCluster_EncryptionInTransitProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#client_broker AwsMskCluster#client_broker}.
	// Experimental.
	ClientBroker *string `field:"optional" json:"clientBroker" yaml:"clientBroker"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#in_cluster AwsMskCluster#in_cluster}.
	// Experimental.
	InCluster interface{} `field:"optional" json:"inCluster" yaml:"inCluster"`
}

