package awsstoragegateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCachedIscsiVolumeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_cached_iscsi_volume#gateway_arn TfCachedIscsiVolume#gateway_arn}.
	// Experimental.
	GatewayArn *string `field:"required" json:"gatewayArn" yaml:"gatewayArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_cached_iscsi_volume#network_interface_id TfCachedIscsiVolume#network_interface_id}.
	// Experimental.
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_cached_iscsi_volume#target_name TfCachedIscsiVolume#target_name}.
	// Experimental.
	TargetName *string `field:"required" json:"targetName" yaml:"targetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_cached_iscsi_volume#volume_size_in_bytes TfCachedIscsiVolume#volume_size_in_bytes}.
	// Experimental.
	VolumeSizeInBytes *float64 `field:"required" json:"volumeSizeInBytes" yaml:"volumeSizeInBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_cached_iscsi_volume#id TfCachedIscsiVolume#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_cached_iscsi_volume#kms_encrypted TfCachedIscsiVolume#kms_encrypted}.
	// Experimental.
	KmsEncrypted interface{} `field:"optional" json:"kmsEncrypted" yaml:"kmsEncrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_cached_iscsi_volume#kms_key TfCachedIscsiVolume#kms_key}.
	// Experimental.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_cached_iscsi_volume#region TfCachedIscsiVolume#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_cached_iscsi_volume#snapshot_id TfCachedIscsiVolume#snapshot_id}.
	// Experimental.
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_cached_iscsi_volume#source_volume_arn TfCachedIscsiVolume#source_volume_arn}.
	// Experimental.
	SourceVolumeArn *string `field:"optional" json:"sourceVolumeArn" yaml:"sourceVolumeArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_cached_iscsi_volume#tags TfCachedIscsiVolume#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_cached_iscsi_volume#tags_all TfCachedIscsiVolume#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

