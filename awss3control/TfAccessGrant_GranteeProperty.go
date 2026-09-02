package awss3control


// Experimental.
type TfAccessGrant_GranteeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_access_grant#grantee_identifier TfAccessGrant#grantee_identifier}.
	// Experimental.
	GranteeIdentifier *string `field:"required" json:"granteeIdentifier" yaml:"granteeIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_access_grant#grantee_type TfAccessGrant#grantee_type}.
	// Experimental.
	GranteeType *string `field:"required" json:"granteeType" yaml:"granteeType"`
}

