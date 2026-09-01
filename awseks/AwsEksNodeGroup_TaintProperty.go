package awseks


// Experimental.
type AwsEksNodeGroup_TaintProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#effect AwsEksNodeGroup#effect}.
	// Experimental.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#key AwsEksNodeGroup#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#value AwsEksNodeGroup#value}.
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

