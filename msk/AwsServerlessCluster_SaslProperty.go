package msk


// Experimental.
type AwsServerlessCluster_SaslProperty struct {
	// iam block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/msk_serverless_cluster#iam AwsServerlessCluster#iam}
	// Experimental.
	Iam *AwsServerlessCluster_IamProperty `field:"required" json:"iam" yaml:"iam"`
}

