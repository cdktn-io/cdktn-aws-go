package eks


// Experimental.
type AwsNodeGroup_TaintProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#effect AwsNodeGroup#effect}.
	// Experimental.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#key AwsNodeGroup#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#value AwsNodeGroup#value}.
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

