package awscodebuild


// Experimental.
type TfProject_RegistryCredentialProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#credential TfProject#credential}.
	// Experimental.
	Credential *string `field:"required" json:"credential" yaml:"credential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#credential_provider TfProject#credential_provider}.
	// Experimental.
	CredentialProvider *string `field:"required" json:"credentialProvider" yaml:"credentialProvider"`
}

