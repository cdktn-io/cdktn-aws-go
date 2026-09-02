package awsfsx


// Experimental.
type TfLustreFileSystem_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#create TfLustreFileSystem#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#delete TfLustreFileSystem#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#update TfLustreFileSystem#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

