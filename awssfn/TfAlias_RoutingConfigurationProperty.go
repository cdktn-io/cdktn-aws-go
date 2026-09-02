package awssfn


// Experimental.
type TfAlias_RoutingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sfn_alias#state_machine_version_arn TfAlias#state_machine_version_arn}.
	// Experimental.
	StateMachineVersionArn *string `field:"required" json:"stateMachineVersionArn" yaml:"stateMachineVersionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sfn_alias#weight TfAlias#weight}.
	// Experimental.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

