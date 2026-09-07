package connect


// Experimental.
type AwsQuickConnect_UserConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_quick_connect#contact_flow_id AwsQuickConnect#contact_flow_id}.
	// Experimental.
	ContactFlowId *string `field:"required" json:"contactFlowId" yaml:"contactFlowId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_quick_connect#user_id AwsQuickConnect#user_id}.
	// Experimental.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}

