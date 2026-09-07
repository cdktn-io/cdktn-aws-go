package eks


// Experimental.
type AwsCluster_EncryptionConfigProperty struct {
	// provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#provider AwsCluster#provider}
	// Experimental.
	Provider *AwsCluster_ProviderProperty `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#resources AwsCluster#resources}.
	// Experimental.
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
}

