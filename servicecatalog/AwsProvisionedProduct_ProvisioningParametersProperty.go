package servicecatalog


// Experimental.
type AwsProvisionedProduct_ProvisioningParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#key AwsProvisionedProduct#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#use_previous_value AwsProvisionedProduct#use_previous_value}.
	// Experimental.
	UsePreviousValue interface{} `field:"optional" json:"usePreviousValue" yaml:"usePreviousValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_provisioned_product#value AwsProvisionedProduct#value}.
	// Experimental.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

