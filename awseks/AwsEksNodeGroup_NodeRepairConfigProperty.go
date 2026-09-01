package awseks


// Experimental.
type AwsEksNodeGroup_NodeRepairConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#enabled AwsEksNodeGroup#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#max_parallel_nodes_repaired_count AwsEksNodeGroup#max_parallel_nodes_repaired_count}.
	// Experimental.
	MaxParallelNodesRepairedCount *float64 `field:"optional" json:"maxParallelNodesRepairedCount" yaml:"maxParallelNodesRepairedCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#max_parallel_nodes_repaired_percentage AwsEksNodeGroup#max_parallel_nodes_repaired_percentage}.
	// Experimental.
	MaxParallelNodesRepairedPercentage *float64 `field:"optional" json:"maxParallelNodesRepairedPercentage" yaml:"maxParallelNodesRepairedPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#max_unhealthy_node_threshold_count AwsEksNodeGroup#max_unhealthy_node_threshold_count}.
	// Experimental.
	MaxUnhealthyNodeThresholdCount *float64 `field:"optional" json:"maxUnhealthyNodeThresholdCount" yaml:"maxUnhealthyNodeThresholdCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#max_unhealthy_node_threshold_percentage AwsEksNodeGroup#max_unhealthy_node_threshold_percentage}.
	// Experimental.
	MaxUnhealthyNodeThresholdPercentage *float64 `field:"optional" json:"maxUnhealthyNodeThresholdPercentage" yaml:"maxUnhealthyNodeThresholdPercentage"`
	// node_repair_config_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#node_repair_config_overrides AwsEksNodeGroup#node_repair_config_overrides}
	// Experimental.
	NodeRepairConfigOverrides interface{} `field:"optional" json:"nodeRepairConfigOverrides" yaml:"nodeRepairConfigOverrides"`
}

