package awseventbridge


// Experimental.
type TfEndpoint_RoutingConfigProperty struct {
	// failover_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_endpoint#failover_config TfEndpoint#failover_config}
	// Experimental.
	FailoverConfig *TfEndpoint_FailoverConfigProperty `field:"required" json:"failoverConfig" yaml:"failoverConfig"`
}

