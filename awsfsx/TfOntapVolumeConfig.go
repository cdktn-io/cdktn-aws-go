package awsfsx

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfOntapVolumeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#name TfOntapVolume#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#storage_virtual_machine_id TfOntapVolume#storage_virtual_machine_id}.
	// Experimental.
	StorageVirtualMachineId *string `field:"required" json:"storageVirtualMachineId" yaml:"storageVirtualMachineId"`
	// aggregate_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#aggregate_configuration TfOntapVolume#aggregate_configuration}
	// Experimental.
	AggregateConfiguration *TfOntapVolume_AggregateConfigurationProperty `field:"optional" json:"aggregateConfiguration" yaml:"aggregateConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#bypass_snaplock_enterprise_retention TfOntapVolume#bypass_snaplock_enterprise_retention}.
	// Experimental.
	BypassSnaplockEnterpriseRetention interface{} `field:"optional" json:"bypassSnaplockEnterpriseRetention" yaml:"bypassSnaplockEnterpriseRetention"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#copy_tags_to_backups TfOntapVolume#copy_tags_to_backups}.
	// Experimental.
	CopyTagsToBackups interface{} `field:"optional" json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#final_backup_tags TfOntapVolume#final_backup_tags}.
	// Experimental.
	FinalBackupTags *map[string]*string `field:"optional" json:"finalBackupTags" yaml:"finalBackupTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#id TfOntapVolume#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#junction_path TfOntapVolume#junction_path}.
	// Experimental.
	JunctionPath *string `field:"optional" json:"junctionPath" yaml:"junctionPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#ontap_volume_type TfOntapVolume#ontap_volume_type}.
	// Experimental.
	OntapVolumeType *string `field:"optional" json:"ontapVolumeType" yaml:"ontapVolumeType"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#region TfOntapVolume#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#security_style TfOntapVolume#security_style}.
	// Experimental.
	SecurityStyle *string `field:"optional" json:"securityStyle" yaml:"securityStyle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#size_in_bytes TfOntapVolume#size_in_bytes}.
	// Experimental.
	SizeInBytes *string `field:"optional" json:"sizeInBytes" yaml:"sizeInBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#size_in_megabytes TfOntapVolume#size_in_megabytes}.
	// Experimental.
	SizeInMegabytes *float64 `field:"optional" json:"sizeInMegabytes" yaml:"sizeInMegabytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#skip_final_backup TfOntapVolume#skip_final_backup}.
	// Experimental.
	SkipFinalBackup interface{} `field:"optional" json:"skipFinalBackup" yaml:"skipFinalBackup"`
	// snaplock_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#snaplock_configuration TfOntapVolume#snaplock_configuration}
	// Experimental.
	SnaplockConfiguration *TfOntapVolume_SnaplockConfigurationProperty `field:"optional" json:"snaplockConfiguration" yaml:"snaplockConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#snapshot_policy TfOntapVolume#snapshot_policy}.
	// Experimental.
	SnapshotPolicy *string `field:"optional" json:"snapshotPolicy" yaml:"snapshotPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#storage_efficiency_enabled TfOntapVolume#storage_efficiency_enabled}.
	// Experimental.
	StorageEfficiencyEnabled interface{} `field:"optional" json:"storageEfficiencyEnabled" yaml:"storageEfficiencyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#tags TfOntapVolume#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#tags_all TfOntapVolume#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// tiering_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#tiering_policy TfOntapVolume#tiering_policy}
	// Experimental.
	TieringPolicy *TfOntapVolume_TieringPolicyProperty `field:"optional" json:"tieringPolicy" yaml:"tieringPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#timeouts TfOntapVolume#timeouts}
	// Experimental.
	Timeouts *TfOntapVolume_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#volume_style TfOntapVolume#volume_style}.
	// Experimental.
	VolumeStyle *string `field:"optional" json:"volumeStyle" yaml:"volumeStyle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_volume#volume_type TfOntapVolume#volume_type}.
	// Experimental.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

