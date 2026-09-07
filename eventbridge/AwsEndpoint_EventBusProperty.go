package eventbridge


// Experimental.
type AwsEndpoint_EventBusProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_endpoint#event_bus_arn AwsEndpoint#event_bus_arn}.
	// Experimental.
	EventBusArn *string `field:"required" json:"eventBusArn" yaml:"eventBusArn"`
}

