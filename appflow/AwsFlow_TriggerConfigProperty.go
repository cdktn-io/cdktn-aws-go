package appflow


// Experimental.
type AwsFlow_TriggerConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#trigger_type AwsFlow#trigger_type}.
	// Experimental.
	TriggerType *string `field:"required" json:"triggerType" yaml:"triggerType"`
	// trigger_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#trigger_properties AwsFlow#trigger_properties}
	// Experimental.
	TriggerProperties *AwsFlow_TriggerPropertiesProperty `field:"optional" json:"triggerProperties" yaml:"triggerProperties"`
}

