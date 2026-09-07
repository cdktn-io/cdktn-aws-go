package sagemakerai


// Experimental.
type AwsEndpoint_BlueGreenUpdatePolicyProperty struct {
	// traffic_routing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#traffic_routing_configuration AwsEndpoint#traffic_routing_configuration}
	// Experimental.
	TrafficRoutingConfiguration *AwsEndpoint_TrafficRoutingConfigurationProperty `field:"required" json:"trafficRoutingConfiguration" yaml:"trafficRoutingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#maximum_execution_timeout_in_seconds AwsEndpoint#maximum_execution_timeout_in_seconds}.
	// Experimental.
	MaximumExecutionTimeoutInSeconds *float64 `field:"optional" json:"maximumExecutionTimeoutInSeconds" yaml:"maximumExecutionTimeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#termination_wait_in_seconds AwsEndpoint#termination_wait_in_seconds}.
	// Experimental.
	TerminationWaitInSeconds *float64 `field:"optional" json:"terminationWaitInSeconds" yaml:"terminationWaitInSeconds"`
}

