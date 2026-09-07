package datasync


// Experimental.
type AwsTask_ReportOverridesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#deleted_override AwsTask#deleted_override}.
	// Experimental.
	DeletedOverride *string `field:"optional" json:"deletedOverride" yaml:"deletedOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#skipped_override AwsTask#skipped_override}.
	// Experimental.
	SkippedOverride *string `field:"optional" json:"skippedOverride" yaml:"skippedOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#transferred_override AwsTask#transferred_override}.
	// Experimental.
	TransferredOverride *string `field:"optional" json:"transferredOverride" yaml:"transferredOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/datasync_task#verified_override AwsTask#verified_override}.
	// Experimental.
	VerifiedOverride *string `field:"optional" json:"verifiedOverride" yaml:"verifiedOverride"`
}

