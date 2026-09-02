package awseks


// Experimental.
type TfCluster_ControlPlanePlacementProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#group_name TfCluster#group_name}.
	// Experimental.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#spread_level TfCluster#spread_level}.
	// Experimental.
	SpreadLevel *string `field:"optional" json:"spreadLevel" yaml:"spreadLevel"`
}

