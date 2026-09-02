package awslexv2models


// Experimental.
type TfIntent_FulfillmentCodeHookProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#enabled TfIntent#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#active TfIntent#active}.
	// Experimental.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// fulfillment_updates_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#fulfillment_updates_specification TfIntent#fulfillment_updates_specification}
	// Experimental.
	FulfillmentUpdatesSpecification interface{} `field:"optional" json:"fulfillmentUpdatesSpecification" yaml:"fulfillmentUpdatesSpecification"`
	// post_fulfillment_status_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#post_fulfillment_status_specification TfIntent#post_fulfillment_status_specification}
	// Experimental.
	PostFulfillmentStatusSpecification interface{} `field:"optional" json:"postFulfillmentStatusSpecification" yaml:"postFulfillmentStatusSpecification"`
}

