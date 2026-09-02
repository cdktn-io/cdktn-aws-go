package awsautoscaling


// Experimental.
type TfGroup_PreferencesProperty struct {
	// alarm_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#alarm_specification TfGroup#alarm_specification}
	// Experimental.
	AlarmSpecification *TfGroup_AlarmSpecificationProperty `field:"optional" json:"alarmSpecification" yaml:"alarmSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#auto_rollback TfGroup#auto_rollback}.
	// Experimental.
	AutoRollback interface{} `field:"optional" json:"autoRollback" yaml:"autoRollback"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#checkpoint_delay TfGroup#checkpoint_delay}.
	// Experimental.
	CheckpointDelay *string `field:"optional" json:"checkpointDelay" yaml:"checkpointDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#checkpoint_percentages TfGroup#checkpoint_percentages}.
	// Experimental.
	CheckpointPercentages *[]*float64 `field:"optional" json:"checkpointPercentages" yaml:"checkpointPercentages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#instance_warmup TfGroup#instance_warmup}.
	// Experimental.
	InstanceWarmup *string `field:"optional" json:"instanceWarmup" yaml:"instanceWarmup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#max_healthy_percentage TfGroup#max_healthy_percentage}.
	// Experimental.
	MaxHealthyPercentage *float64 `field:"optional" json:"maxHealthyPercentage" yaml:"maxHealthyPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#min_healthy_percentage TfGroup#min_healthy_percentage}.
	// Experimental.
	MinHealthyPercentage *float64 `field:"optional" json:"minHealthyPercentage" yaml:"minHealthyPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#scale_in_protected_instances TfGroup#scale_in_protected_instances}.
	// Experimental.
	ScaleInProtectedInstances *string `field:"optional" json:"scaleInProtectedInstances" yaml:"scaleInProtectedInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#skip_matching TfGroup#skip_matching}.
	// Experimental.
	SkipMatching interface{} `field:"optional" json:"skipMatching" yaml:"skipMatching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_group#standby_instances TfGroup#standby_instances}.
	// Experimental.
	StandbyInstances *string `field:"optional" json:"standbyInstances" yaml:"standbyInstances"`
}

