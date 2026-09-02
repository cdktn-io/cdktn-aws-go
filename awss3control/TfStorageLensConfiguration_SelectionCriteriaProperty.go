package awss3control


// Experimental.
type TfStorageLensConfiguration_SelectionCriteriaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#delimiter TfStorageLensConfiguration#delimiter}.
	// Experimental.
	Delimiter *string `field:"optional" json:"delimiter" yaml:"delimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#max_depth TfStorageLensConfiguration#max_depth}.
	// Experimental.
	MaxDepth *float64 `field:"optional" json:"maxDepth" yaml:"maxDepth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_storage_lens_configuration#min_storage_bytes_percentage TfStorageLensConfiguration#min_storage_bytes_percentage}.
	// Experimental.
	MinStorageBytesPercentage *float64 `field:"optional" json:"minStorageBytesPercentage" yaml:"minStorageBytesPercentage"`
}

