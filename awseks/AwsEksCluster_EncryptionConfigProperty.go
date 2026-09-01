package awseks


// Experimental.
type AwsEksCluster_EncryptionConfigProperty struct {
	// provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#provider AwsEksCluster#provider}
	// Experimental.
	Provider *AwsEksCluster_ProviderProperty `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#resources AwsEksCluster#resources}.
	// Experimental.
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
}

