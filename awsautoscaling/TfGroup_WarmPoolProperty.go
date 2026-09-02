package awsautoscaling


// Experimental.
type TfGroup_WarmPoolProperty struct {
	// instance_reuse_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#instance_reuse_policy TfGroup#instance_reuse_policy}
	// Experimental.
	InstanceReusePolicy *TfGroup_InstanceReusePolicyProperty `field:"optional" json:"instanceReusePolicy" yaml:"instanceReusePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#max_group_prepared_capacity TfGroup#max_group_prepared_capacity}.
	// Experimental.
	MaxGroupPreparedCapacity *float64 `field:"optional" json:"maxGroupPreparedCapacity" yaml:"maxGroupPreparedCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#min_size TfGroup#min_size}.
	// Experimental.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#pool_state TfGroup#pool_state}.
	// Experimental.
	PoolState *string `field:"optional" json:"poolState" yaml:"poolState"`
}

