package awsoracledatabaseaws

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCloudVmClusterConfig struct {
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
	// The number of CPU cores to enable on the VM cluster. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#cpu_core_count TfCloudVmCluster#cpu_core_count}
	// Experimental.
	CpuCoreCount *float64 `field:"required" json:"cpuCoreCount" yaml:"cpuCoreCount"`
	// The size of the data disk group, in terabytes (TBs), to allocate for the VM cluster.
	//
	// Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#data_storage_size_in_tbs TfCloudVmCluster#data_storage_size_in_tbs}
	// Experimental.
	DataStorageSizeInTbs *float64 `field:"required" json:"dataStorageSizeInTbs" yaml:"dataStorageSizeInTbs"`
	// The list of database servers for the VM cluster. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#db_servers TfCloudVmCluster#db_servers}
	// Experimental.
	DbServers *[]*string `field:"required" json:"dbServers" yaml:"dbServers"`
	// A user-friendly name for the VM cluster. This member is required. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#display_name TfCloudVmCluster#display_name}
	// Experimental.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// A valid software version of Oracle Grid Infrastructure (GI).
	//
	// To get the list of valid values, use the ListGiVersions operation and specify the shape of the Exadata infrastructure. Example: 19.0.0.0 This member is required. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#gi_version TfCloudVmCluster#gi_version}
	// Experimental.
	GiVersion *string `field:"required" json:"giVersion" yaml:"giVersion"`
	// The host name prefix for the VM cluster.
	//
	// Constraints: - Can't be "localhost" or "hostname". - Can't contain "-version". - The maximum length of the combined hostname and domain is 63 characters. - The hostname must be unique within the subnet. This member is required. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#hostname_prefix TfCloudVmCluster#hostname_prefix}
	// Experimental.
	HostnamePrefix *string `field:"required" json:"hostnamePrefix" yaml:"hostnamePrefix"`
	// The public key portion of one or more key pairs used for SSH access to the VM cluster.
	//
	// This member is required. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#ssh_public_keys TfCloudVmCluster#ssh_public_keys}
	// Experimental.
	SshPublicKeys *[]*string `field:"required" json:"sshPublicKeys" yaml:"sshPublicKeys"`
	// The unique identifier of the Exadata infrastructure for this VM cluster. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#cloud_exadata_infrastructure_arn TfCloudVmCluster#cloud_exadata_infrastructure_arn}
	// Experimental.
	CloudExadataInfrastructureArn *string `field:"optional" json:"cloudExadataInfrastructureArn" yaml:"cloudExadataInfrastructureArn"`
	// The unique identifier of the Exadata infrastructure for this VM cluster. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#cloud_exadata_infrastructure_id TfCloudVmCluster#cloud_exadata_infrastructure_id}
	// Experimental.
	CloudExadataInfrastructureId *string `field:"optional" json:"cloudExadataInfrastructureId" yaml:"cloudExadataInfrastructureId"`
	// The name of the Grid Infrastructure (GI) cluster. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#cluster_name TfCloudVmCluster#cluster_name}
	// Experimental.
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
	// data_collection_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#data_collection_options TfCloudVmCluster#data_collection_options}
	// Experimental.
	DataCollectionOptions interface{} `field:"optional" json:"dataCollectionOptions" yaml:"dataCollectionOptions"`
	// The amount of local node storage, in gigabytes (GBs), to allocate for the VM cluster.
	//
	// Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#db_node_storage_size_in_gbs TfCloudVmCluster#db_node_storage_size_in_gbs}
	// Experimental.
	DbNodeStorageSizeInGbs *float64 `field:"optional" json:"dbNodeStorageSizeInGbs" yaml:"dbNodeStorageSizeInGbs"`
	// Specifies whether to enable database backups to local Exadata storage for the VM cluster.
	//
	// Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#is_local_backup_enabled TfCloudVmCluster#is_local_backup_enabled}
	// Experimental.
	IsLocalBackupEnabled interface{} `field:"optional" json:"isLocalBackupEnabled" yaml:"isLocalBackupEnabled"`
	// Specifies whether to create a sparse disk group for the VM cluster. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#is_sparse_diskgroup_enabled TfCloudVmCluster#is_sparse_diskgroup_enabled}
	// Experimental.
	IsSparseDiskgroupEnabled interface{} `field:"optional" json:"isSparseDiskgroupEnabled" yaml:"isSparseDiskgroupEnabled"`
	// The Oracle license model to apply to the VM cluster. Default: LICENSE_INCLUDED. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#license_model TfCloudVmCluster#license_model}
	// Experimental.
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// The amount of memory, in gigabytes (GBs), to allocate for the VM cluster.
	//
	// Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#memory_size_in_gbs TfCloudVmCluster#memory_size_in_gbs}
	// Experimental.
	MemorySizeInGbs *float64 `field:"optional" json:"memorySizeInGbs" yaml:"memorySizeInGbs"`
	// The unique identifier of the ODB network for the VM cluster.
	//
	// This member is required. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#odb_network_arn TfCloudVmCluster#odb_network_arn}
	// Experimental.
	OdbNetworkArn *string `field:"optional" json:"odbNetworkArn" yaml:"odbNetworkArn"`
	// The unique identifier of the ODB network for the VM cluster.
	//
	// This member is required. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#odb_network_id TfCloudVmCluster#odb_network_id}
	// Experimental.
	OdbNetworkId *string `field:"optional" json:"odbNetworkId" yaml:"odbNetworkId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#region TfCloudVmCluster#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// The port number for TCP connections to the single client access name (SCAN) listener.
	//
	// Valid values: 1024–8999 with the following exceptions: 2484 , 6100 , 6200 , 7060, 7070 , 7085 , and 7879Default: 1521. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#scan_listener_port_tcp TfCloudVmCluster#scan_listener_port_tcp}
	// Experimental.
	ScanListenerPortTcp *float64 `field:"optional" json:"scanListenerPortTcp" yaml:"scanListenerPortTcp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#tags TfCloudVmCluster#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#timeouts TfCloudVmCluster#timeouts}
	// Experimental.
	Timeouts *TfCloudVmCluster_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The configured time zone of the VM cluster. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_vm_cluster#timezone TfCloudVmCluster#timezone}
	// Experimental.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}

