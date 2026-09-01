package awscodebuild


// Experimental.
type AwsCodebuildProject_SecondarySourceVersionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#source_identifier AwsCodebuildProject#source_identifier}.
	// Experimental.
	SourceIdentifier *string `field:"required" json:"sourceIdentifier" yaml:"sourceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#source_version AwsCodebuildProject#source_version}.
	// Experimental.
	SourceVersion *string `field:"required" json:"sourceVersion" yaml:"sourceVersion"`
}

