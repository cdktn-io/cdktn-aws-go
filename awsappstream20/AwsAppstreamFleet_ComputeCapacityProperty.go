package awsappstream20


// Experimental.
type AwsAppstreamFleet_ComputeCapacityProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#desired_instances AwsAppstreamFleet#desired_instances}.
	// Experimental.
	DesiredInstances *float64 `field:"optional" json:"desiredInstances" yaml:"desiredInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_fleet#desired_sessions AwsAppstreamFleet#desired_sessions}.
	// Experimental.
	DesiredSessions *float64 `field:"optional" json:"desiredSessions" yaml:"desiredSessions"`
}

