package awsbedrockagentcore


// Experimental.
type TfHarness_AgentcoreGatewayProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#gateway_arn TfHarness#gateway_arn}.
	// Experimental.
	GatewayArn *string `field:"required" json:"gatewayArn" yaml:"gatewayArn"`
	// outbound_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness#outbound_auth TfHarness#outbound_auth}
	// Experimental.
	OutboundAuth interface{} `field:"optional" json:"outboundAuth" yaml:"outboundAuth"`
}

