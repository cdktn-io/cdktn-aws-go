package awsram


// Experimental.
type TfResourceShare_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ram_resource_share#create TfResourceShare#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ram_resource_share#delete TfResourceShare#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

