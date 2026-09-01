package awsappflow


// Experimental.
type AwsAppflowFlow_TriggerConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#trigger_type AwsAppflowFlow#trigger_type}.
	// Experimental.
	TriggerType *string `field:"required" json:"triggerType" yaml:"triggerType"`
	// trigger_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#trigger_properties AwsAppflowFlow#trigger_properties}
	// Experimental.
	TriggerProperties *AwsAppflowFlow_TriggerPropertiesProperty `field:"optional" json:"triggerProperties" yaml:"triggerProperties"`
}

