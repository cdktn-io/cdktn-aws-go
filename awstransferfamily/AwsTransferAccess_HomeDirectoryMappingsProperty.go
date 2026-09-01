package awstransferfamily


// Experimental.
type AwsTransferAccess_HomeDirectoryMappingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_access#entry AwsTransferAccess#entry}.
	// Experimental.
	Entry *string `field:"required" json:"entry" yaml:"entry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_access#target AwsTransferAccess#target}.
	// Experimental.
	Target *string `field:"required" json:"target" yaml:"target"`
}

