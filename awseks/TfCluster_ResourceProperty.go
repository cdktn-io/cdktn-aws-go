package awseks


// Experimental.
type TfCluster_ResourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#name TfCluster#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#weight TfCluster#weight}.
	// Experimental.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

