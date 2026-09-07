package codebuild


// Experimental.
type AwsProject_RestrictionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#compute_types_allowed AwsProject#compute_types_allowed}.
	// Experimental.
	ComputeTypesAllowed *[]*string `field:"optional" json:"computeTypesAllowed" yaml:"computeTypesAllowed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#maximum_builds_allowed AwsProject#maximum_builds_allowed}.
	// Experimental.
	MaximumBuildsAllowed *float64 `field:"optional" json:"maximumBuildsAllowed" yaml:"maximumBuildsAllowed"`
}

