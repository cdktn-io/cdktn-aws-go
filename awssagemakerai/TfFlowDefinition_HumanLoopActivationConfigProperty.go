package awssagemakerai


// Experimental.
type TfFlowDefinition_HumanLoopActivationConfigProperty struct {
	// human_loop_activation_conditions_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_flow_definition#human_loop_activation_conditions_config TfFlowDefinition#human_loop_activation_conditions_config}
	// Experimental.
	HumanLoopActivationConditionsConfig *TfFlowDefinition_HumanLoopActivationConditionsConfigProperty `field:"optional" json:"humanLoopActivationConditionsConfig" yaml:"humanLoopActivationConditionsConfig"`
}

