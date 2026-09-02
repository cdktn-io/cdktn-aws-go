package awsconnect


// Experimental.
type TfInstance_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance#create TfInstance#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_instance#delete TfInstance#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

