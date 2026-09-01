package awseventbridge


// Experimental.
type AwsCloudwatchEventEndpoint_RoutingConfigProperty struct {
	// failover_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_endpoint#failover_config AwsCloudwatchEventEndpoint#failover_config}
	// Experimental.
	FailoverConfig *AwsCloudwatchEventEndpoint_FailoverConfigProperty `field:"required" json:"failoverConfig" yaml:"failoverConfig"`
}

