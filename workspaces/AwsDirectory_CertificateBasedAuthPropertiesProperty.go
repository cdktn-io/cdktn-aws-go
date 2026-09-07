package workspaces


// Experimental.
type AwsDirectory_CertificateBasedAuthPropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#certificate_authority_arn AwsDirectory#certificate_authority_arn}.
	// Experimental.
	CertificateAuthorityArn *string `field:"optional" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#status AwsDirectory#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

