package sfn


// Experimental.
type AwsStateMachine_LoggingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sfn_state_machine#include_execution_data AwsStateMachine#include_execution_data}.
	// Experimental.
	IncludeExecutionData interface{} `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sfn_state_machine#level AwsStateMachine#level}.
	// Experimental.
	Level *string `field:"optional" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sfn_state_machine#log_destination AwsStateMachine#log_destination}.
	// Experimental.
	LogDestination *string `field:"optional" json:"logDestination" yaml:"logDestination"`
}

