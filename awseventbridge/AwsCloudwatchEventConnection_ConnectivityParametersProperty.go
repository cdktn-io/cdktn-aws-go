package awseventbridge


// Experimental.
type AwsCloudwatchEventConnection_ConnectivityParametersProperty struct {
	// resource_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#resource_parameters AwsCloudwatchEventConnection#resource_parameters}
	// Experimental.
	ResourceParameters *AwsCloudwatchEventConnection_AuthParametersConnectivityParametersResourceParametersProperty `field:"required" json:"resourceParameters" yaml:"resourceParameters"`
}

