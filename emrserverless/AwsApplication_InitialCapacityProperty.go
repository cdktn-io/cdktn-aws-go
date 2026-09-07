package emrserverless


// Experimental.
type AwsApplication_InitialCapacityProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#initial_capacity_type AwsApplication#initial_capacity_type}.
	// Experimental.
	InitialCapacityType *string `field:"required" json:"initialCapacityType" yaml:"initialCapacityType"`
	// initial_capacity_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emrserverless_application#initial_capacity_config AwsApplication#initial_capacity_config}
	// Experimental.
	InitialCapacityConfig *AwsApplication_InitialCapacityConfigProperty `field:"optional" json:"initialCapacityConfig" yaml:"initialCapacityConfig"`
}

