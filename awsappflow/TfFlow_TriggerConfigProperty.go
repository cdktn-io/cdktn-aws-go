package awsappflow


// Experimental.
type TfFlow_TriggerConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#trigger_type TfFlow#trigger_type}.
	// Experimental.
	TriggerType *string `field:"required" json:"triggerType" yaml:"triggerType"`
	// trigger_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appflow_flow#trigger_properties TfFlow#trigger_properties}
	// Experimental.
	TriggerProperties *TfFlow_TriggerPropertiesProperty `field:"optional" json:"triggerProperties" yaml:"triggerProperties"`
}

