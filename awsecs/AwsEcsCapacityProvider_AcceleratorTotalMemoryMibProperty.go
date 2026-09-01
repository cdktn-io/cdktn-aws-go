package awsecs


// Experimental.
type AwsEcsCapacityProvider_AcceleratorTotalMemoryMibProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#max AwsEcsCapacityProvider#max}.
	// Experimental.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#min AwsEcsCapacityProvider#min}.
	// Experimental.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

