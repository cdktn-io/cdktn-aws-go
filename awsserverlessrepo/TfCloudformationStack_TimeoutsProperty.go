package awsserverlessrepo


// Experimental.
type TfCloudformationStack_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#create TfCloudformationStack#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#delete TfCloudformationStack#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#update TfCloudformationStack#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

