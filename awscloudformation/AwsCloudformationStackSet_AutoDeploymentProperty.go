package awscloudformation


// Experimental.
type AwsCloudformationStackSet_AutoDeploymentProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#depends_on_stack_sets AwsCloudformationStackSet#depends_on_stack_sets}.
	// Experimental.
	DependsOnStackSets *[]*string `field:"optional" json:"dependsOnStackSets" yaml:"dependsOnStackSets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#enabled AwsCloudformationStackSet#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudformation_stack_set#retain_stacks_on_account_removal AwsCloudformationStackSet#retain_stacks_on_account_removal}.
	// Experimental.
	RetainStacksOnAccountRemoval interface{} `field:"optional" json:"retainStacksOnAccountRemoval" yaml:"retainStacksOnAccountRemoval"`
}

