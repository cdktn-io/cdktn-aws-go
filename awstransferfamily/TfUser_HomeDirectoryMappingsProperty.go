package awstransferfamily


// Experimental.
type TfUser_HomeDirectoryMappingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_user#entry TfUser#entry}.
	// Experimental.
	Entry *string `field:"required" json:"entry" yaml:"entry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_user#target TfUser#target}.
	// Experimental.
	Target *string `field:"required" json:"target" yaml:"target"`
}

