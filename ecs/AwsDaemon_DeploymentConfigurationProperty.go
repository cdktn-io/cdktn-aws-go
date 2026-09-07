package ecs


// Experimental.
type AwsDaemon_DeploymentConfigurationProperty struct {
	// alarms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon#alarms AwsDaemon#alarms}
	// Experimental.
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon#bake_time_in_minutes AwsDaemon#bake_time_in_minutes}.
	// Experimental.
	BakeTimeInMinutes *float64 `field:"optional" json:"bakeTimeInMinutes" yaml:"bakeTimeInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_daemon#drain_percent AwsDaemon#drain_percent}.
	// Experimental.
	DrainPercent *float64 `field:"optional" json:"drainPercent" yaml:"drainPercent"`
}

