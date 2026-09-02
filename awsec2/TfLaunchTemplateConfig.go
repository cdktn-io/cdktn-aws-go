package awsec2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfLaunchTemplateConfig struct {
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
	// block_device_mappings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#block_device_mappings TfLaunchTemplate#block_device_mappings}
	// Experimental.
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// capacity_reservation_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#capacity_reservation_specification TfLaunchTemplate#capacity_reservation_specification}
	// Experimental.
	CapacityReservationSpecification *TfLaunchTemplate_CapacityReservationSpecificationProperty `field:"optional" json:"capacityReservationSpecification" yaml:"capacityReservationSpecification"`
	// cpu_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#cpu_options TfLaunchTemplate#cpu_options}
	// Experimental.
	CpuOptions *TfLaunchTemplate_CpuOptionsProperty `field:"optional" json:"cpuOptions" yaml:"cpuOptions"`
	// credit_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#credit_specification TfLaunchTemplate#credit_specification}
	// Experimental.
	CreditSpecification *TfLaunchTemplate_CreditSpecificationProperty `field:"optional" json:"creditSpecification" yaml:"creditSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#default_version TfLaunchTemplate#default_version}.
	// Experimental.
	DefaultVersion *float64 `field:"optional" json:"defaultVersion" yaml:"defaultVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#description TfLaunchTemplate#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#disable_api_stop TfLaunchTemplate#disable_api_stop}.
	// Experimental.
	DisableApiStop interface{} `field:"optional" json:"disableApiStop" yaml:"disableApiStop"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#disable_api_termination TfLaunchTemplate#disable_api_termination}.
	// Experimental.
	DisableApiTermination interface{} `field:"optional" json:"disableApiTermination" yaml:"disableApiTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#ebs_optimized TfLaunchTemplate#ebs_optimized}.
	// Experimental.
	EbsOptimized *string `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// enclave_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#enclave_options TfLaunchTemplate#enclave_options}
	// Experimental.
	EnclaveOptions *TfLaunchTemplate_EnclaveOptionsProperty `field:"optional" json:"enclaveOptions" yaml:"enclaveOptions"`
	// hibernation_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#hibernation_options TfLaunchTemplate#hibernation_options}
	// Experimental.
	HibernationOptions *TfLaunchTemplate_HibernationOptionsProperty `field:"optional" json:"hibernationOptions" yaml:"hibernationOptions"`
	// iam_instance_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#iam_instance_profile TfLaunchTemplate#iam_instance_profile}
	// Experimental.
	IamInstanceProfile *TfLaunchTemplate_IamInstanceProfileProperty `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#id TfLaunchTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#image_id TfLaunchTemplate#image_id}.
	// Experimental.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#instance_initiated_shutdown_behavior TfLaunchTemplate#instance_initiated_shutdown_behavior}.
	// Experimental.
	InstanceInitiatedShutdownBehavior *string `field:"optional" json:"instanceInitiatedShutdownBehavior" yaml:"instanceInitiatedShutdownBehavior"`
	// instance_market_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#instance_market_options TfLaunchTemplate#instance_market_options}
	// Experimental.
	InstanceMarketOptions *TfLaunchTemplate_InstanceMarketOptionsProperty `field:"optional" json:"instanceMarketOptions" yaml:"instanceMarketOptions"`
	// instance_requirements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#instance_requirements TfLaunchTemplate#instance_requirements}
	// Experimental.
	InstanceRequirements *TfLaunchTemplate_InstanceRequirementsProperty `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#instance_type TfLaunchTemplate#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#kernel_id TfLaunchTemplate#kernel_id}.
	// Experimental.
	KernelId *string `field:"optional" json:"kernelId" yaml:"kernelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#key_name TfLaunchTemplate#key_name}.
	// Experimental.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// license_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#license_specification TfLaunchTemplate#license_specification}
	// Experimental.
	LicenseSpecification interface{} `field:"optional" json:"licenseSpecification" yaml:"licenseSpecification"`
	// maintenance_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#maintenance_options TfLaunchTemplate#maintenance_options}
	// Experimental.
	MaintenanceOptions *TfLaunchTemplate_MaintenanceOptionsProperty `field:"optional" json:"maintenanceOptions" yaml:"maintenanceOptions"`
	// metadata_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#metadata_options TfLaunchTemplate#metadata_options}
	// Experimental.
	MetadataOptions *TfLaunchTemplate_MetadataOptionsProperty `field:"optional" json:"metadataOptions" yaml:"metadataOptions"`
	// monitoring block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#monitoring TfLaunchTemplate#monitoring}
	// Experimental.
	Monitoring *TfLaunchTemplate_MonitoringProperty `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#name TfLaunchTemplate#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#name_prefix TfLaunchTemplate#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// network_interfaces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#network_interfaces TfLaunchTemplate#network_interfaces}
	// Experimental.
	NetworkInterfaces interface{} `field:"optional" json:"networkInterfaces" yaml:"networkInterfaces"`
	// network_performance_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#network_performance_options TfLaunchTemplate#network_performance_options}
	// Experimental.
	NetworkPerformanceOptions *TfLaunchTemplate_NetworkPerformanceOptionsProperty `field:"optional" json:"networkPerformanceOptions" yaml:"networkPerformanceOptions"`
	// placement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#placement TfLaunchTemplate#placement}
	// Experimental.
	Placement *TfLaunchTemplate_PlacementProperty `field:"optional" json:"placement" yaml:"placement"`
	// private_dns_name_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#private_dns_name_options TfLaunchTemplate#private_dns_name_options}
	// Experimental.
	PrivateDnsNameOptions *TfLaunchTemplate_PrivateDnsNameOptionsProperty `field:"optional" json:"privateDnsNameOptions" yaml:"privateDnsNameOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#ram_disk_id TfLaunchTemplate#ram_disk_id}.
	// Experimental.
	RamDiskId *string `field:"optional" json:"ramDiskId" yaml:"ramDiskId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#region TfLaunchTemplate#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// secondary_interfaces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#secondary_interfaces TfLaunchTemplate#secondary_interfaces}
	// Experimental.
	SecondaryInterfaces interface{} `field:"optional" json:"secondaryInterfaces" yaml:"secondaryInterfaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#security_group_names TfLaunchTemplate#security_group_names}.
	// Experimental.
	SecurityGroupNames *[]*string `field:"optional" json:"securityGroupNames" yaml:"securityGroupNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#tags TfLaunchTemplate#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#tags_all TfLaunchTemplate#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// tag_specifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#tag_specifications TfLaunchTemplate#tag_specifications}
	// Experimental.
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#update_default_version TfLaunchTemplate#update_default_version}.
	// Experimental.
	UpdateDefaultVersion interface{} `field:"optional" json:"updateDefaultVersion" yaml:"updateDefaultVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#user_data TfLaunchTemplate#user_data}.
	// Experimental.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#vpc_security_group_ids TfLaunchTemplate#vpc_security_group_ids}.
	// Experimental.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

