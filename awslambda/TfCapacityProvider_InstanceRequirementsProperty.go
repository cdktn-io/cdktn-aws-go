package awslambda


// Experimental.
type TfCapacityProvider_InstanceRequirementsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_capacity_provider#allowed_instance_types TfCapacityProvider#allowed_instance_types}.
	// Experimental.
	AllowedInstanceTypes *[]*string `field:"optional" json:"allowedInstanceTypes" yaml:"allowedInstanceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_capacity_provider#architectures TfCapacityProvider#architectures}.
	// Experimental.
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_capacity_provider#excluded_instance_types TfCapacityProvider#excluded_instance_types}.
	// Experimental.
	ExcludedInstanceTypes *[]*string `field:"optional" json:"excludedInstanceTypes" yaml:"excludedInstanceTypes"`
}

