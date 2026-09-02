package awsoracledatabaseaws

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCloudExadataInfrastructureConfig struct {
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
	// The AZ ID of the AZ where the Exadata infrastructure is located.
	//
	// Changing this will force terraform to create new resource
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#availability_zone_id TfCloudExadataInfrastructure#availability_zone_id}
	// Experimental.
	AvailabilityZoneId *string `field:"required" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// The user-friendly name for the Exadata infrastructure. Changing this will force terraform to create a new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#display_name TfCloudExadataInfrastructure#display_name}
	// Experimental.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The model name of the Exadata infrastructure. Changing this will force terraform to create new resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#shape TfCloudExadataInfrastructure#shape}
	// Experimental.
	Shape *string `field:"required" json:"shape" yaml:"shape"`
	// The name of the Availability Zone (AZ) where the Exadata infrastructure is located.
	//
	// Changing this will force terraform to create new resource
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#availability_zone TfCloudExadataInfrastructure#availability_zone}
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// The number of compute instances that the Exadata infrastructure is located.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#compute_count TfCloudExadataInfrastructure#compute_count}
	// Experimental.
	ComputeCount *float64 `field:"optional" json:"computeCount" yaml:"computeCount"`
	// The email addresses of contacts to receive notification from Oracle about maintenance updates for the Exadata infrastructure.
	//
	// Changing this will force terraform to create new resource
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#customer_contacts_to_send_to_oci TfCloudExadataInfrastructure#customer_contacts_to_send_to_oci}
	// Experimental.
	CustomerContactsToSendToOci interface{} `field:"optional" json:"customerContactsToSendToOci" yaml:"customerContactsToSendToOci"`
	// The database server model type of the Exadata infrastructure.
	//
	// For the list of valid model names, use the ListDbSystemShapes operation
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#database_server_type TfCloudExadataInfrastructure#database_server_type}
	// Experimental.
	DatabaseServerType *string `field:"optional" json:"databaseServerType" yaml:"databaseServerType"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#maintenance_window TfCloudExadataInfrastructure#maintenance_window}
	// Experimental.
	MaintenanceWindow interface{} `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#region TfCloudExadataInfrastructure#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// TThe number of storage servers that are activated for the Exadata infrastructure.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#storage_count TfCloudExadataInfrastructure#storage_count}
	// Experimental.
	StorageCount *float64 `field:"optional" json:"storageCount" yaml:"storageCount"`
	// The storage server model type of the Exadata infrastructure.
	//
	// For the list of valid model names, use the ListDbSystemShapes operation
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#storage_server_type TfCloudExadataInfrastructure#storage_server_type}
	// Experimental.
	StorageServerType *string `field:"optional" json:"storageServerType" yaml:"storageServerType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#tags TfCloudExadataInfrastructure#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_cloud_exadata_infrastructure#timeouts TfCloudExadataInfrastructure#timeouts}
	// Experimental.
	Timeouts *TfCloudExadataInfrastructure_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

