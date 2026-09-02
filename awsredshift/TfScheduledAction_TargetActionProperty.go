package awsredshift


// Experimental.
type TfScheduledAction_TargetActionProperty struct {
	// pause_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_scheduled_action#pause_cluster TfScheduledAction#pause_cluster}
	// Experimental.
	PauseCluster *TfScheduledAction_PauseClusterProperty `field:"optional" json:"pauseCluster" yaml:"pauseCluster"`
	// resize_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_scheduled_action#resize_cluster TfScheduledAction#resize_cluster}
	// Experimental.
	ResizeCluster *TfScheduledAction_ResizeClusterProperty `field:"optional" json:"resizeCluster" yaml:"resizeCluster"`
	// resume_cluster block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_scheduled_action#resume_cluster TfScheduledAction#resume_cluster}
	// Experimental.
	ResumeCluster *TfScheduledAction_ResumeClusterProperty `field:"optional" json:"resumeCluster" yaml:"resumeCluster"`
}

