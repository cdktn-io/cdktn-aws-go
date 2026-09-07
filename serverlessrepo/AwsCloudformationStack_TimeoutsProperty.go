package serverlessrepo


// Experimental.
type AwsCloudformationStack_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#create AwsCloudformationStack#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#delete AwsCloudformationStack#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/serverlessapplicationrepository_cloudformation_stack#update AwsCloudformationStack#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

