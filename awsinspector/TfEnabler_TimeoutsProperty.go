package awsinspector


// Experimental.
type TfEnabler_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_enabler#create TfEnabler#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_enabler#delete TfEnabler#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_enabler#update TfEnabler#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

