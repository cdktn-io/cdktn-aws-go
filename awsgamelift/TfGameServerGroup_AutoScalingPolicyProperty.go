package awsgamelift


// Experimental.
type TfGameServerGroup_AutoScalingPolicyProperty struct {
	// target_tracking_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#target_tracking_configuration TfGameServerGroup#target_tracking_configuration}
	// Experimental.
	TargetTrackingConfiguration *TfGameServerGroup_TargetTrackingConfigurationProperty `field:"required" json:"targetTrackingConfiguration" yaml:"targetTrackingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/gamelift_game_server_group#estimated_instance_warmup TfGameServerGroup#estimated_instance_warmup}.
	// Experimental.
	EstimatedInstanceWarmup *float64 `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
}

