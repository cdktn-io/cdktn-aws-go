package eventbridge


// Experimental.
type AwsConnection_InvocationConnectivityParametersProperty struct {
	// resource_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_connection#resource_parameters AwsConnection#resource_parameters}
	// Experimental.
	ResourceParameters *AwsConnection_InvocationConnectivityParametersResourceParametersProperty `field:"required" json:"resourceParameters" yaml:"resourceParameters"`
}

