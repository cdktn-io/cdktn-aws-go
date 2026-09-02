package awslambda


// Experimental.
type TfCapacityProvider_CapacityProviderScalingConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_capacity_provider#max_vcpu_count TfCapacityProvider#max_vcpu_count}.
	// Experimental.
	MaxVcpuCount *float64 `field:"optional" json:"maxVcpuCount" yaml:"maxVcpuCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_capacity_provider#scaling_mode TfCapacityProvider#scaling_mode}.
	// Experimental.
	ScalingMode *string `field:"optional" json:"scalingMode" yaml:"scalingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_capacity_provider#scaling_policies TfCapacityProvider#scaling_policies}.
	// Experimental.
	ScalingPolicies interface{} `field:"optional" json:"scalingPolicies" yaml:"scalingPolicies"`
}

