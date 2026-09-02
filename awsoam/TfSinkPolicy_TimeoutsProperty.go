package awsoam


// Experimental.
type TfSinkPolicy_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/oam_sink_policy#create TfSinkPolicy#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/oam_sink_policy#delete TfSinkPolicy#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/oam_sink_policy#update TfSinkPolicy#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

