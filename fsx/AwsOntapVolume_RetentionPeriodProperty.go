package fsx


// Experimental.
type AwsOntapVolume_RetentionPeriodProperty struct {
	// default_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#default_retention AwsOntapVolume#default_retention}
	// Experimental.
	DefaultRetention *AwsOntapVolume_DefaultRetentionProperty `field:"optional" json:"defaultRetention" yaml:"defaultRetention"`
	// maximum_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#maximum_retention AwsOntapVolume#maximum_retention}
	// Experimental.
	MaximumRetention *AwsOntapVolume_MaximumRetentionProperty `field:"optional" json:"maximumRetention" yaml:"maximumRetention"`
	// minimum_retention block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#minimum_retention AwsOntapVolume#minimum_retention}
	// Experimental.
	MinimumRetention *AwsOntapVolume_MinimumRetentionProperty `field:"optional" json:"minimumRetention" yaml:"minimumRetention"`
}

