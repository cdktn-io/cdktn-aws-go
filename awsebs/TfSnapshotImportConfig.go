package awsebs

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSnapshotImportConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// disk_container block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#disk_container TfSnapshotImport#disk_container}
	// Experimental.
	DiskContainer *TfSnapshotImport_DiskContainerProperty `field:"required" json:"diskContainer" yaml:"diskContainer"`
	// client_data block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#client_data TfSnapshotImport#client_data}
	// Experimental.
	ClientData *TfSnapshotImport_ClientDataProperty `field:"optional" json:"clientData" yaml:"clientData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#description TfSnapshotImport#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#encrypted TfSnapshotImport#encrypted}.
	// Experimental.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#id TfSnapshotImport#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#kms_key_id TfSnapshotImport#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#permanent_restore TfSnapshotImport#permanent_restore}.
	// Experimental.
	PermanentRestore interface{} `field:"optional" json:"permanentRestore" yaml:"permanentRestore"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#region TfSnapshotImport#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#role_name TfSnapshotImport#role_name}.
	// Experimental.
	RoleName *string `field:"optional" json:"roleName" yaml:"roleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#storage_tier TfSnapshotImport#storage_tier}.
	// Experimental.
	StorageTier *string `field:"optional" json:"storageTier" yaml:"storageTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#tags TfSnapshotImport#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#tags_all TfSnapshotImport#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#temporary_restore_days TfSnapshotImport#temporary_restore_days}.
	// Experimental.
	TemporaryRestoreDays *float64 `field:"optional" json:"temporaryRestoreDays" yaml:"temporaryRestoreDays"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ebs_snapshot_import#timeouts TfSnapshotImport#timeouts}
	// Experimental.
	Timeouts *TfSnapshotImport_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

