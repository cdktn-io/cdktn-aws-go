package msk


// Experimental.
type AwsCluster_EncryptionInTransitProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#client_broker AwsCluster#client_broker}.
	// Experimental.
	ClientBroker *string `field:"optional" json:"clientBroker" yaml:"clientBroker"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_cluster#in_cluster AwsCluster#in_cluster}.
	// Experimental.
	InCluster interface{} `field:"optional" json:"inCluster" yaml:"inCluster"`
}

