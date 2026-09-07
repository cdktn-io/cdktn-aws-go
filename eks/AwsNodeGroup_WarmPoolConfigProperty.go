package eks


// Experimental.
type AwsNodeGroup_WarmPoolConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#max_group_prepared_capacity AwsNodeGroup#max_group_prepared_capacity}.
	// Experimental.
	MaxGroupPreparedCapacity *float64 `field:"optional" json:"maxGroupPreparedCapacity" yaml:"maxGroupPreparedCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#min_size AwsNodeGroup#min_size}.
	// Experimental.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#pool_state AwsNodeGroup#pool_state}.
	// Experimental.
	PoolState *string `field:"optional" json:"poolState" yaml:"poolState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#reuse_on_scale_in AwsNodeGroup#reuse_on_scale_in}.
	// Experimental.
	ReuseOnScaleIn interface{} `field:"optional" json:"reuseOnScaleIn" yaml:"reuseOnScaleIn"`
}

