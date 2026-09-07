package elb


// Experimental.
type AwsAlbTargetGroup_TargetHealthStateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_target_group#enable_unhealthy_connection_termination AwsAlbTargetGroup#enable_unhealthy_connection_termination}.
	// Experimental.
	EnableUnhealthyConnectionTermination interface{} `field:"required" json:"enableUnhealthyConnectionTermination" yaml:"enableUnhealthyConnectionTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_target_group#unhealthy_draining_interval AwsAlbTargetGroup#unhealthy_draining_interval}.
	// Experimental.
	UnhealthyDrainingInterval *float64 `field:"optional" json:"unhealthyDrainingInterval" yaml:"unhealthyDrainingInterval"`
}

