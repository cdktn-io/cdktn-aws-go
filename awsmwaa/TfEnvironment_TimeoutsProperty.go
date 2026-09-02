package awsmwaa


// Experimental.
type TfEnvironment_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#create TfEnvironment#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#delete TfEnvironment#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mwaa_environment#update TfEnvironment#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

