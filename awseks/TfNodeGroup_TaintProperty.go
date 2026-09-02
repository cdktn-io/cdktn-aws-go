package awseks


// Experimental.
type TfNodeGroup_TaintProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#effect TfNodeGroup#effect}.
	// Experimental.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#key TfNodeGroup#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#value TfNodeGroup#value}.
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

