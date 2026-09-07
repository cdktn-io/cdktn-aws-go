package fsx


// Experimental.
type AwsOntapVolume_TieringPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#cooling_period AwsOntapVolume#cooling_period}.
	// Experimental.
	CoolingPeriod *float64 `field:"optional" json:"coolingPeriod" yaml:"coolingPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#name AwsOntapVolume#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

