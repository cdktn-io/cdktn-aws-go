package awsfsx


// Experimental.
type TfOntapVolume_RetentionPeriodProperty struct {
	// default_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#default_retention TfOntapVolume#default_retention}
	// Experimental.
	DefaultRetention *TfOntapVolume_DefaultRetentionProperty `field:"optional" json:"defaultRetention" yaml:"defaultRetention"`
	// maximum_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#maximum_retention TfOntapVolume#maximum_retention}
	// Experimental.
	MaximumRetention *TfOntapVolume_MaximumRetentionProperty `field:"optional" json:"maximumRetention" yaml:"maximumRetention"`
	// minimum_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#minimum_retention TfOntapVolume#minimum_retention}
	// Experimental.
	MinimumRetention *TfOntapVolume_MinimumRetentionProperty `field:"optional" json:"minimumRetention" yaml:"minimumRetention"`
}

