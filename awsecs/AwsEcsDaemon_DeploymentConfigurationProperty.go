package awsecs


// Experimental.
type AwsEcsDaemon_DeploymentConfigurationProperty struct {
	// alarms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon#alarms AwsEcsDaemon#alarms}
	// Experimental.
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon#bake_time_in_minutes AwsEcsDaemon#bake_time_in_minutes}.
	// Experimental.
	BakeTimeInMinutes *float64 `field:"optional" json:"bakeTimeInMinutes" yaml:"bakeTimeInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon#drain_percent AwsEcsDaemon#drain_percent}.
	// Experimental.
	DrainPercent *float64 `field:"optional" json:"drainPercent" yaml:"drainPercent"`
}

