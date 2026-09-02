package awsfsx


// Experimental.
type TfOntapVolume_TieringPolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#cooling_period TfOntapVolume#cooling_period}.
	// Experimental.
	CoolingPeriod *float64 `field:"optional" json:"coolingPeriod" yaml:"coolingPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#name TfOntapVolume#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

