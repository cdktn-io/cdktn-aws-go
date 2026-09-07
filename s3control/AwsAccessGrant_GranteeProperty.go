package s3control


// Experimental.
type AwsAccessGrant_GranteeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_access_grant#grantee_identifier AwsAccessGrant#grantee_identifier}.
	// Experimental.
	GranteeIdentifier *string `field:"required" json:"granteeIdentifier" yaml:"granteeIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_access_grant#grantee_type AwsAccessGrant#grantee_type}.
	// Experimental.
	GranteeType *string `field:"required" json:"granteeType" yaml:"granteeType"`
}

