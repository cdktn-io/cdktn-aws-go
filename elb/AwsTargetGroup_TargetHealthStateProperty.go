package elb


// Experimental.
type AwsTargetGroup_TargetHealthStateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#enable_unhealthy_connection_termination AwsTargetGroup#enable_unhealthy_connection_termination}.
	// Experimental.
	EnableUnhealthyConnectionTermination interface{} `field:"required" json:"enableUnhealthyConnectionTermination" yaml:"enableUnhealthyConnectionTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#unhealthy_draining_interval AwsTargetGroup#unhealthy_draining_interval}.
	// Experimental.
	UnhealthyDrainingInterval *float64 `field:"optional" json:"unhealthyDrainingInterval" yaml:"unhealthyDrainingInterval"`
}

