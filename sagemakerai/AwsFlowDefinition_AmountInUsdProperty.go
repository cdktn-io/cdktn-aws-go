package sagemakerai


// Experimental.
type AwsFlowDefinition_AmountInUsdProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_flow_definition#cents AwsFlowDefinition#cents}.
	// Experimental.
	Cents *float64 `field:"optional" json:"cents" yaml:"cents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_flow_definition#dollars AwsFlowDefinition#dollars}.
	// Experimental.
	Dollars *float64 `field:"optional" json:"dollars" yaml:"dollars"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_flow_definition#tenth_fractions_of_a_cent AwsFlowDefinition#tenth_fractions_of_a_cent}.
	// Experimental.
	TenthFractionsOfACent *float64 `field:"optional" json:"tenthFractionsOfACent" yaml:"tenthFractionsOfACent"`
}

