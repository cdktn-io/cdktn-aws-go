package awss3control


// Experimental.
type TfStorageLensConfiguration_IncludeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#buckets TfStorageLensConfiguration#buckets}.
	// Experimental.
	Buckets *[]*string `field:"optional" json:"buckets" yaml:"buckets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#regions TfStorageLensConfiguration#regions}.
	// Experimental.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

