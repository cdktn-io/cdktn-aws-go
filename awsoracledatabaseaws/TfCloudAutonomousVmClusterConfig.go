package awsoracledatabaseaws

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCloudAutonomousVmClusterConfig struct {
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
	// The data storage size allocated for Autonomous Databases in the Autonomous VM cluster, in TB.
	//
	// Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#autonomous_data_storage_size_in_tbs TfCloudAutonomousVmCluster#autonomous_data_storage_size_in_tbs}
	// Experimental.
	AutonomousDataStorageSizeInTbs *float64 `field:"required" json:"autonomousDataStorageSizeInTbs" yaml:"autonomousDataStorageSizeInTbs"`
	// The number of CPU cores enabled per node in the Autonomous VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#cpu_core_count_per_node TfCloudAutonomousVmCluster#cpu_core_count_per_node}
	// Experimental.
	CpuCoreCountPerNode *float64 `field:"required" json:"cpuCoreCountPerNode" yaml:"cpuCoreCountPerNode"`
	// The database servers in the Autonomous VM cluster. Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#db_servers TfCloudAutonomousVmCluster#db_servers}
	// Experimental.
	DbServers *[]*string `field:"required" json:"dbServers" yaml:"dbServers"`
	// The display name of the Autonomous VM cluster. Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#display_name TfCloudAutonomousVmCluster#display_name}
	// Experimental.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The amount of memory allocated per Oracle Compute Unit, in GB.
	//
	// Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#memory_per_oracle_compute_unit_in_gbs TfCloudAutonomousVmCluster#memory_per_oracle_compute_unit_in_gbs}
	// Experimental.
	MemoryPerOracleComputeUnitInGbs *float64 `field:"required" json:"memoryPerOracleComputeUnitInGbs" yaml:"memoryPerOracleComputeUnitInGbs"`
	// The SCAN listener port for non-TLS (TCP) protocol.
	//
	// The default is 1521. Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#scan_listener_port_non_tls TfCloudAutonomousVmCluster#scan_listener_port_non_tls}
	// Experimental.
	ScanListenerPortNonTls *float64 `field:"required" json:"scanListenerPortNonTls" yaml:"scanListenerPortNonTls"`
	// The SCAN listener port for TLS (TCP) protocol.
	//
	// The default is 2484. Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#scan_listener_port_tls TfCloudAutonomousVmCluster#scan_listener_port_tls}
	// Experimental.
	ScanListenerPortTls *float64 `field:"required" json:"scanListenerPortTls" yaml:"scanListenerPortTls"`
	// The total number of Autonomous Container Databases that can be created with the allocated local storage.
	//
	// Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#total_container_databases TfCloudAutonomousVmCluster#total_container_databases}
	// Experimental.
	TotalContainerDatabases *float64 `field:"required" json:"totalContainerDatabases" yaml:"totalContainerDatabases"`
	// The unique identifier of the Exadata infrastructure for this VM cluster. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#cloud_exadata_infrastructure_arn TfCloudAutonomousVmCluster#cloud_exadata_infrastructure_arn}
	// Experimental.
	CloudExadataInfrastructureArn *string `field:"optional" json:"cloudExadataInfrastructureArn" yaml:"cloudExadataInfrastructureArn"`
	// Exadata infrastructure id. Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#cloud_exadata_infrastructure_id TfCloudAutonomousVmCluster#cloud_exadata_infrastructure_id}
	// Experimental.
	CloudExadataInfrastructureId *string `field:"optional" json:"cloudExadataInfrastructureId" yaml:"cloudExadataInfrastructureId"`
	// The description of the Autonomous VM cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#description TfCloudAutonomousVmCluster#description}
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Indicates whether mutual TLS (mTLS) authentication is enabled for the Autonomous VM cluster.
	//
	// Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#is_mtls_enabled_vm_cluster TfCloudAutonomousVmCluster#is_mtls_enabled_vm_cluster}
	// Experimental.
	IsMtlsEnabledVmCluster interface{} `field:"optional" json:"isMtlsEnabledVmCluster" yaml:"isMtlsEnabledVmCluster"`
	// The license model for the Autonomous VM cluster.
	//
	// Valid values are LICENSE_INCLUDED or BRING_YOUR_OWN_LICENSE . Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#license_model TfCloudAutonomousVmCluster#license_model}
	// Experimental.
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#maintenance_window TfCloudAutonomousVmCluster#maintenance_window}
	// Experimental.
	MaintenanceWindow interface{} `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// The unique identifier of the ODB network for the VM cluster.
	//
	// This member is required. Changing this will create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#odb_network_arn TfCloudAutonomousVmCluster#odb_network_arn}
	// Experimental.
	OdbNetworkArn *string `field:"optional" json:"odbNetworkArn" yaml:"odbNetworkArn"`
	// The unique identifier of the ODB network associated with this Autonomous VM Cluster.
	//
	// Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#odb_network_id TfCloudAutonomousVmCluster#odb_network_id}
	// Experimental.
	OdbNetworkId *string `field:"optional" json:"odbNetworkId" yaml:"odbNetworkId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#region TfCloudAutonomousVmCluster#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#tags TfCloudAutonomousVmCluster#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#timeouts TfCloudAutonomousVmCluster#timeouts}
	// Experimental.
	Timeouts *TfCloudAutonomousVmCluster_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The time zone of the Autonomous VM cluster. Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_autonomous_vm_cluster#time_zone TfCloudAutonomousVmCluster#time_zone}
	// Experimental.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

