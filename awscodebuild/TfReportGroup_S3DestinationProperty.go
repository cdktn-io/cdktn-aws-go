package awscodebuild


// Experimental.
type TfReportGroup_S3DestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_report_group#bucket TfReportGroup#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_report_group#encryption_key TfReportGroup#encryption_key}.
	// Experimental.
	EncryptionKey *string `field:"required" json:"encryptionKey" yaml:"encryptionKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_report_group#encryption_disabled TfReportGroup#encryption_disabled}.
	// Experimental.
	EncryptionDisabled interface{} `field:"optional" json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_report_group#packaging TfReportGroup#packaging}.
	// Experimental.
	Packaging *string `field:"optional" json:"packaging" yaml:"packaging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_report_group#path TfReportGroup#path}.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

