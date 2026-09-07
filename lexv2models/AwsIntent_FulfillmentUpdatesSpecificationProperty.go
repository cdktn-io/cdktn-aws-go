package lexv2models


// Experimental.
type AwsIntent_FulfillmentUpdatesSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#active AwsIntent#active}.
	// Experimental.
	Active interface{} `field:"required" json:"active" yaml:"active"`
	// start_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#start_response AwsIntent#start_response}
	// Experimental.
	StartResponse interface{} `field:"optional" json:"startResponse" yaml:"startResponse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#timeout_in_seconds AwsIntent#timeout_in_seconds}.
	// Experimental.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
	// update_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#update_response AwsIntent#update_response}
	// Experimental.
	UpdateResponse interface{} `field:"optional" json:"updateResponse" yaml:"updateResponse"`
}

