package awsservicecatalog


// Experimental.
type TfTagOptionResourceAssociation_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_tag_option_resource_association#create TfTagOptionResourceAssociation#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_tag_option_resource_association#delete TfTagOptionResourceAssociation#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/servicecatalog_tag_option_resource_association#read TfTagOptionResourceAssociation#read}.
	// Experimental.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

