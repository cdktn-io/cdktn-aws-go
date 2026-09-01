package awsredshift


// Experimental.
type AwsRedshiftScheduledAction_TargetActionProperty struct {
	// pause_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_scheduled_action#pause_cluster AwsRedshiftScheduledAction#pause_cluster}
	// Experimental.
	PauseCluster *AwsRedshiftScheduledAction_PauseClusterProperty `field:"optional" json:"pauseCluster" yaml:"pauseCluster"`
	// resize_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_scheduled_action#resize_cluster AwsRedshiftScheduledAction#resize_cluster}
	// Experimental.
	ResizeCluster *AwsRedshiftScheduledAction_ResizeClusterProperty `field:"optional" json:"resizeCluster" yaml:"resizeCluster"`
	// resume_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_scheduled_action#resume_cluster AwsRedshiftScheduledAction#resume_cluster}
	// Experimental.
	ResumeCluster *AwsRedshiftScheduledAction_ResumeClusterProperty `field:"optional" json:"resumeCluster" yaml:"resumeCluster"`
}

