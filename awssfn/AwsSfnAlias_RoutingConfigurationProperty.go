package awssfn


// Experimental.
type AwsSfnAlias_RoutingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sfn_alias#state_machine_version_arn AwsSfnAlias#state_machine_version_arn}.
	// Experimental.
	StateMachineVersionArn *string `field:"required" json:"stateMachineVersionArn" yaml:"stateMachineVersionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sfn_alias#weight AwsSfnAlias#weight}.
	// Experimental.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

