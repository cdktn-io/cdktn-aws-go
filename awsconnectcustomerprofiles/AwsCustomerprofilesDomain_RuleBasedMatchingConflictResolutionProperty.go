package awsconnectcustomerprofiles


// Experimental.
type AwsCustomerprofilesDomain_RuleBasedMatchingConflictResolutionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#conflict_resolving_model AwsCustomerprofilesDomain#conflict_resolving_model}.
	// Experimental.
	ConflictResolvingModel *string `field:"required" json:"conflictResolvingModel" yaml:"conflictResolvingModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#source_name AwsCustomerprofilesDomain#source_name}.
	// Experimental.
	SourceName *string `field:"optional" json:"sourceName" yaml:"sourceName"`
}

