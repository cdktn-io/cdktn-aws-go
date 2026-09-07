package sagemakerai


// Experimental.
type AwsEndpointConfiguration_ShadowProductionVariantsManagedInstanceScalingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#max_instance_count AwsEndpointConfiguration#max_instance_count}.
	// Experimental.
	MaxInstanceCount *float64 `field:"optional" json:"maxInstanceCount" yaml:"maxInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#min_instance_count AwsEndpointConfiguration#min_instance_count}.
	// Experimental.
	MinInstanceCount *float64 `field:"optional" json:"minInstanceCount" yaml:"minInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#status AwsEndpointConfiguration#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

