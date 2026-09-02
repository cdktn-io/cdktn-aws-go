package awstransferfamily


// Experimental.
type TfAccess_HomeDirectoryMappingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_access#entry TfAccess#entry}.
	// Experimental.
	Entry *string `field:"required" json:"entry" yaml:"entry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_access#target TfAccess#target}.
	// Experimental.
	Target *string `field:"required" json:"target" yaml:"target"`
}

