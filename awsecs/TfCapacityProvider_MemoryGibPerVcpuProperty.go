package awsecs


// Experimental.
type TfCapacityProvider_MemoryGibPerVcpuProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#max TfCapacityProvider#max}.
	// Experimental.
	Max *float64 `field:"optional" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#min TfCapacityProvider#min}.
	// Experimental.
	Min *float64 `field:"optional" json:"min" yaml:"min"`
}

