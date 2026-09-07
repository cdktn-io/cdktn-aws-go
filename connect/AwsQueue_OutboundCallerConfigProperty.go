package connect


// Experimental.
type AwsQueue_OutboundCallerConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_queue#outbound_caller_id_name AwsQueue#outbound_caller_id_name}.
	// Experimental.
	OutboundCallerIdName *string `field:"optional" json:"outboundCallerIdName" yaml:"outboundCallerIdName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_queue#outbound_caller_id_number_id AwsQueue#outbound_caller_id_number_id}.
	// Experimental.
	OutboundCallerIdNumberId *string `field:"optional" json:"outboundCallerIdNumberId" yaml:"outboundCallerIdNumberId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_queue#outbound_flow_id AwsQueue#outbound_flow_id}.
	// Experimental.
	OutboundFlowId *string `field:"optional" json:"outboundFlowId" yaml:"outboundFlowId"`
}

