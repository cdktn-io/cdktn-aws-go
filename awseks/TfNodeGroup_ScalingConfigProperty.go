package awseks


// Experimental.
type TfNodeGroup_ScalingConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#desired_size TfNodeGroup#desired_size}.
	// Experimental.
	DesiredSize *float64 `field:"required" json:"desiredSize" yaml:"desiredSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#max_size TfNodeGroup#max_size}.
	// Experimental.
	MaxSize *float64 `field:"required" json:"maxSize" yaml:"maxSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#min_size TfNodeGroup#min_size}.
	// Experimental.
	MinSize *float64 `field:"required" json:"minSize" yaml:"minSize"`
}

