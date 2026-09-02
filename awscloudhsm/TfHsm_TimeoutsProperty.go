package awscloudhsm


// Experimental.
type TfHsm_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudhsm_v2_hsm#create TfHsm#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudhsm_v2_hsm#delete TfHsm#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

