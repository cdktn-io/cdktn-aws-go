package awsconnect


// Experimental.
type TfQuickConnect_QueueConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_quick_connect#contact_flow_id TfQuickConnect#contact_flow_id}.
	// Experimental.
	ContactFlowId *string `field:"required" json:"contactFlowId" yaml:"contactFlowId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_quick_connect#queue_id TfQuickConnect#queue_id}.
	// Experimental.
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
}

