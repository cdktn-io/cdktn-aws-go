package awseks


// Experimental.
type TfCluster_EncryptionConfigProperty struct {
	// provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#provider TfCluster#provider}
	// Experimental.
	Provider *TfCluster_ProviderProperty `field:"required" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#resources TfCluster#resources}.
	// Experimental.
	Resources *[]*string `field:"required" json:"resources" yaml:"resources"`
}

