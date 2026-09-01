package awstransferfamily


// Experimental.
type AwsTransferUser_HomeDirectoryMappingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_user#entry AwsTransferUser#entry}.
	// Experimental.
	Entry *string `field:"required" json:"entry" yaml:"entry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_user#target AwsTransferUser#target}.
	// Experimental.
	Target *string `field:"required" json:"target" yaml:"target"`
}

