package eks


// Experimental.
type AwsCluster_ControlPlanePlacementProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#group_name AwsCluster#group_name}.
	// Experimental.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#spread_level AwsCluster#spread_level}.
	// Experimental.
	SpreadLevel *string `field:"optional" json:"spreadLevel" yaml:"spreadLevel"`
}

