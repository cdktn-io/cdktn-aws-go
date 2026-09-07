package fsx

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsOpenzfsVolumeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#name AwsOpenzfsVolume#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#parent_volume_id AwsOpenzfsVolume#parent_volume_id}.
	// Experimental.
	ParentVolumeId *string `field:"required" json:"parentVolumeId" yaml:"parentVolumeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#copy_tags_to_snapshots AwsOpenzfsVolume#copy_tags_to_snapshots}.
	// Experimental.
	CopyTagsToSnapshots interface{} `field:"optional" json:"copyTagsToSnapshots" yaml:"copyTagsToSnapshots"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#data_compression_type AwsOpenzfsVolume#data_compression_type}.
	// Experimental.
	DataCompressionType *string `field:"optional" json:"dataCompressionType" yaml:"dataCompressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#delete_volume_options AwsOpenzfsVolume#delete_volume_options}.
	// Experimental.
	DeleteVolumeOptions *[]*string `field:"optional" json:"deleteVolumeOptions" yaml:"deleteVolumeOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#id AwsOpenzfsVolume#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// nfs_exports block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#nfs_exports AwsOpenzfsVolume#nfs_exports}
	// Experimental.
	NfsExports *AwsOpenzfsVolume_NfsExportsProperty `field:"optional" json:"nfsExports" yaml:"nfsExports"`
	// origin_snapshot block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#origin_snapshot AwsOpenzfsVolume#origin_snapshot}
	// Experimental.
	OriginSnapshot *AwsOpenzfsVolume_OriginSnapshotProperty `field:"optional" json:"originSnapshot" yaml:"originSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#read_only AwsOpenzfsVolume#read_only}.
	// Experimental.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#record_size_kib AwsOpenzfsVolume#record_size_kib}.
	// Experimental.
	RecordSizeKib *float64 `field:"optional" json:"recordSizeKib" yaml:"recordSizeKib"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#region AwsOpenzfsVolume#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#storage_capacity_quota_gib AwsOpenzfsVolume#storage_capacity_quota_gib}.
	// Experimental.
	StorageCapacityQuotaGib *float64 `field:"optional" json:"storageCapacityQuotaGib" yaml:"storageCapacityQuotaGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#storage_capacity_reservation_gib AwsOpenzfsVolume#storage_capacity_reservation_gib}.
	// Experimental.
	StorageCapacityReservationGib *float64 `field:"optional" json:"storageCapacityReservationGib" yaml:"storageCapacityReservationGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#tags AwsOpenzfsVolume#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#tags_all AwsOpenzfsVolume#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#timeouts AwsOpenzfsVolume#timeouts}
	// Experimental.
	Timeouts *AwsOpenzfsVolume_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// user_and_group_quotas block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#user_and_group_quotas AwsOpenzfsVolume#user_and_group_quotas}
	// Experimental.
	UserAndGroupQuotas interface{} `field:"optional" json:"userAndGroupQuotas" yaml:"userAndGroupQuotas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#volume_type AwsOpenzfsVolume#volume_type}.
	// Experimental.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

