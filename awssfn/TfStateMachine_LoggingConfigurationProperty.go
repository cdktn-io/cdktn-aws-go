package awssfn


// Experimental.
type TfStateMachine_LoggingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sfn_state_machine#include_execution_data TfStateMachine#include_execution_data}.
	// Experimental.
	IncludeExecutionData interface{} `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sfn_state_machine#level TfStateMachine#level}.
	// Experimental.
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sfn_state_machine#log_destination TfStateMachine#log_destination}.
	// Experimental.
	LogDestination *string `field:"optional" json:"logDestination" yaml:"logDestination"`
}

