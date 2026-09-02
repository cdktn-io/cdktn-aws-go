package awseks


// Experimental.
type TfNodeGroup_UpdateConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#max_unavailable TfNodeGroup#max_unavailable}.
	// Experimental.
	MaxUnavailable *float64 `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#max_unavailable_percentage TfNodeGroup#max_unavailable_percentage}.
	// Experimental.
	MaxUnavailablePercentage *float64 `field:"optional" json:"maxUnavailablePercentage" yaml:"maxUnavailablePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#update_strategy TfNodeGroup#update_strategy}.
	// Experimental.
	UpdateStrategy *string `field:"optional" json:"updateStrategy" yaml:"updateStrategy"`
}

