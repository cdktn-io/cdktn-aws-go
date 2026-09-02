package awscontroltower


// Experimental.
type TfLandingZone_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/controltower_landing_zone#create TfLandingZone#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/controltower_landing_zone#delete TfLandingZone#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/controltower_landing_zone#update TfLandingZone#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

