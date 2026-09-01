package awsfsx


// Experimental.
type AwsFsxOntapVolume_TieringPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#cooling_period AwsFsxOntapVolume#cooling_period}.
	// Experimental.
	CoolingPeriod *float64 `field:"optional" json:"coolingPeriod" yaml:"coolingPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#name AwsFsxOntapVolume#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

