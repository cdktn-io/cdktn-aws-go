package storagegateway


// Experimental.
type AwsSmbFileShare_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#create AwsSmbFileShare#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#delete AwsSmbFileShare#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#update AwsSmbFileShare#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

