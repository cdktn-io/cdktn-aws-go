package awseks


// Experimental.
type AwsEksNodeGroup_NodeRepairConfigOverridesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#min_repair_wait_time_mins AwsEksNodeGroup#min_repair_wait_time_mins}.
	// Experimental.
	MinRepairWaitTimeMins *float64 `field:"required" json:"minRepairWaitTimeMins" yaml:"minRepairWaitTimeMins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#node_monitoring_condition AwsEksNodeGroup#node_monitoring_condition}.
	// Experimental.
	NodeMonitoringCondition *string `field:"required" json:"nodeMonitoringCondition" yaml:"nodeMonitoringCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#node_unhealthy_reason AwsEksNodeGroup#node_unhealthy_reason}.
	// Experimental.
	NodeUnhealthyReason *string `field:"required" json:"nodeUnhealthyReason" yaml:"nodeUnhealthyReason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#repair_action AwsEksNodeGroup#repair_action}.
	// Experimental.
	RepairAction *string `field:"required" json:"repairAction" yaml:"repairAction"`
}

