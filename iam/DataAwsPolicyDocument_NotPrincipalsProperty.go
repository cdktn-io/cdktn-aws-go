package iam


// Experimental.
type DataAwsPolicyDocument_NotPrincipalsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#identifiers DataAwsPolicyDocument#identifiers}.
	// Experimental.
	Identifiers *[]*string `field:"required" json:"identifiers" yaml:"identifiers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/iam_policy_document#type DataAwsPolicyDocument#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

