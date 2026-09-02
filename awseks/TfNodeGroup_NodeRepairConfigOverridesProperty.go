package awseks


// Experimental.
type TfNodeGroup_NodeRepairConfigOverridesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#min_repair_wait_time_mins TfNodeGroup#min_repair_wait_time_mins}.
	// Experimental.
	MinRepairWaitTimeMins *float64 `field:"required" json:"minRepairWaitTimeMins" yaml:"minRepairWaitTimeMins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#node_monitoring_condition TfNodeGroup#node_monitoring_condition}.
	// Experimental.
	NodeMonitoringCondition *string `field:"required" json:"nodeMonitoringCondition" yaml:"nodeMonitoringCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#node_unhealthy_reason TfNodeGroup#node_unhealthy_reason}.
	// Experimental.
	NodeUnhealthyReason *string `field:"required" json:"nodeUnhealthyReason" yaml:"nodeUnhealthyReason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_node_group#repair_action TfNodeGroup#repair_action}.
	// Experimental.
	RepairAction *string `field:"required" json:"repairAction" yaml:"repairAction"`
}

