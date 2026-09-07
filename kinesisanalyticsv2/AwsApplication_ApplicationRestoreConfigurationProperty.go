package kinesisanalyticsv2


// Experimental.
type AwsApplication_ApplicationRestoreConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#application_restore_type AwsApplication#application_restore_type}.
	// Experimental.
	ApplicationRestoreType *string `field:"optional" json:"applicationRestoreType" yaml:"applicationRestoreType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kinesisanalyticsv2_application#snapshot_name AwsApplication#snapshot_name}.
	// Experimental.
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
}

