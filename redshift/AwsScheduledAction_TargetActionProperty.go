package redshift


// Experimental.
type AwsScheduledAction_TargetActionProperty struct {
	// pause_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_scheduled_action#pause_cluster AwsScheduledAction#pause_cluster}
	// Experimental.
	PauseCluster *AwsScheduledAction_PauseClusterProperty `field:"optional" json:"pauseCluster" yaml:"pauseCluster"`
	// resize_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_scheduled_action#resize_cluster AwsScheduledAction#resize_cluster}
	// Experimental.
	ResizeCluster *AwsScheduledAction_ResizeClusterProperty `field:"optional" json:"resizeCluster" yaml:"resizeCluster"`
	// resume_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_scheduled_action#resume_cluster AwsScheduledAction#resume_cluster}
	// Experimental.
	ResumeCluster *AwsScheduledAction_ResumeClusterProperty `field:"optional" json:"resumeCluster" yaml:"resumeCluster"`
}

