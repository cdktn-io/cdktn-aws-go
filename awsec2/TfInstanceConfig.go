package awsec2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#ami TfInstance#ami}.
	// Experimental.
	Ami *string `field:"optional" json:"ami" yaml:"ami"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#associate_public_ip_address TfInstance#associate_public_ip_address}.
	// Experimental.
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#availability_zone TfInstance#availability_zone}.
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// capacity_reservation_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#capacity_reservation_specification TfInstance#capacity_reservation_specification}
	// Experimental.
	CapacityReservationSpecification *TfInstance_CapacityReservationSpecificationProperty `field:"optional" json:"capacityReservationSpecification" yaml:"capacityReservationSpecification"`
	// cpu_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#cpu_options TfInstance#cpu_options}
	// Experimental.
	CpuOptions *TfInstance_CpuOptionsProperty `field:"optional" json:"cpuOptions" yaml:"cpuOptions"`
	// credit_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#credit_specification TfInstance#credit_specification}
	// Experimental.
	CreditSpecification *TfInstance_CreditSpecificationProperty `field:"optional" json:"creditSpecification" yaml:"creditSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#disable_api_stop TfInstance#disable_api_stop}.
	// Experimental.
	DisableApiStop interface{} `field:"optional" json:"disableApiStop" yaml:"disableApiStop"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#disable_api_termination TfInstance#disable_api_termination}.
	// Experimental.
	DisableApiTermination interface{} `field:"optional" json:"disableApiTermination" yaml:"disableApiTermination"`
	// ebs_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#ebs_block_device TfInstance#ebs_block_device}
	// Experimental.
	EbsBlockDevice interface{} `field:"optional" json:"ebsBlockDevice" yaml:"ebsBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#ebs_optimized TfInstance#ebs_optimized}.
	// Experimental.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#enable_primary_ipv6 TfInstance#enable_primary_ipv6}.
	// Experimental.
	EnablePrimaryIpv6 interface{} `field:"optional" json:"enablePrimaryIpv6" yaml:"enablePrimaryIpv6"`
	// enclave_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#enclave_options TfInstance#enclave_options}
	// Experimental.
	EnclaveOptions *TfInstance_EnclaveOptionsProperty `field:"optional" json:"enclaveOptions" yaml:"enclaveOptions"`
	// ephemeral_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#ephemeral_block_device TfInstance#ephemeral_block_device}
	// Experimental.
	EphemeralBlockDevice interface{} `field:"optional" json:"ephemeralBlockDevice" yaml:"ephemeralBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#get_password_data TfInstance#get_password_data}.
	// Experimental.
	FetchPasswordData interface{} `field:"optional" json:"fetchPasswordData" yaml:"fetchPasswordData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#force_destroy TfInstance#force_destroy}.
	// Experimental.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#hibernation TfInstance#hibernation}.
	// Experimental.
	Hibernation interface{} `field:"optional" json:"hibernation" yaml:"hibernation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#host_id TfInstance#host_id}.
	// Experimental.
	HostId *string `field:"optional" json:"hostId" yaml:"hostId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#host_resource_group_arn TfInstance#host_resource_group_arn}.
	// Experimental.
	HostResourceGroupArn *string `field:"optional" json:"hostResourceGroupArn" yaml:"hostResourceGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#iam_instance_profile TfInstance#iam_instance_profile}.
	// Experimental.
	IamInstanceProfile *string `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#id TfInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#instance_initiated_shutdown_behavior TfInstance#instance_initiated_shutdown_behavior}.
	// Experimental.
	InstanceInitiatedShutdownBehavior *string `field:"optional" json:"instanceInitiatedShutdownBehavior" yaml:"instanceInitiatedShutdownBehavior"`
	// instance_market_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#instance_market_options TfInstance#instance_market_options}
	// Experimental.
	InstanceMarketOptions *TfInstance_InstanceMarketOptionsProperty `field:"optional" json:"instanceMarketOptions" yaml:"instanceMarketOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#instance_type TfInstance#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#ipv6_address_count TfInstance#ipv6_address_count}.
	// Experimental.
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#ipv6_addresses TfInstance#ipv6_addresses}.
	// Experimental.
	Ipv6Addresses *[]*string `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#key_name TfInstance#key_name}.
	// Experimental.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#launch_template TfInstance#launch_template}
	// Experimental.
	LaunchTemplate *TfInstance_LaunchTemplateProperty `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// maintenance_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#maintenance_options TfInstance#maintenance_options}
	// Experimental.
	MaintenanceOptions *TfInstance_MaintenanceOptionsProperty `field:"optional" json:"maintenanceOptions" yaml:"maintenanceOptions"`
	// metadata_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#metadata_options TfInstance#metadata_options}
	// Experimental.
	MetadataOptions *TfInstance_MetadataOptionsProperty `field:"optional" json:"metadataOptions" yaml:"metadataOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#monitoring TfInstance#monitoring}.
	// Experimental.
	Monitoring interface{} `field:"optional" json:"monitoring" yaml:"monitoring"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#network_interface TfInstance#network_interface}
	// Experimental.
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#placement_group TfInstance#placement_group}.
	// Experimental.
	PlacementGroup *string `field:"optional" json:"placementGroup" yaml:"placementGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#placement_group_id TfInstance#placement_group_id}.
	// Experimental.
	PlacementGroupId *string `field:"optional" json:"placementGroupId" yaml:"placementGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#placement_partition_number TfInstance#placement_partition_number}.
	// Experimental.
	PlacementPartitionNumber *float64 `field:"optional" json:"placementPartitionNumber" yaml:"placementPartitionNumber"`
	// primary_network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#primary_network_interface TfInstance#primary_network_interface}
	// Experimental.
	PrimaryNetworkInterface *TfInstance_PrimaryNetworkInterfaceProperty `field:"optional" json:"primaryNetworkInterface" yaml:"primaryNetworkInterface"`
	// private_dns_name_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#private_dns_name_options TfInstance#private_dns_name_options}
	// Experimental.
	PrivateDnsNameOptions *TfInstance_PrivateDnsNameOptionsProperty `field:"optional" json:"privateDnsNameOptions" yaml:"privateDnsNameOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#private_ip TfInstance#private_ip}.
	// Experimental.
	PrivateIp *string `field:"optional" json:"privateIp" yaml:"privateIp"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#region TfInstance#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// root_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#root_block_device TfInstance#root_block_device}
	// Experimental.
	RootBlockDevice *TfInstance_RootBlockDeviceProperty `field:"optional" json:"rootBlockDevice" yaml:"rootBlockDevice"`
	// secondary_network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#secondary_network_interface TfInstance#secondary_network_interface}
	// Experimental.
	SecondaryNetworkInterface interface{} `field:"optional" json:"secondaryNetworkInterface" yaml:"secondaryNetworkInterface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#secondary_private_ips TfInstance#secondary_private_ips}.
	// Experimental.
	SecondaryPrivateIps *[]*string `field:"optional" json:"secondaryPrivateIps" yaml:"secondaryPrivateIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#security_groups TfInstance#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#source_dest_check TfInstance#source_dest_check}.
	// Experimental.
	SourceDestCheck interface{} `field:"optional" json:"sourceDestCheck" yaml:"sourceDestCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#subnet_id TfInstance#subnet_id}.
	// Experimental.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#tags TfInstance#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#tags_all TfInstance#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#tenancy TfInstance#tenancy}.
	// Experimental.
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#timeouts TfInstance#timeouts}
	// Experimental.
	Timeouts *TfInstance_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#user_data TfInstance#user_data}.
	// Experimental.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#user_data_base64 TfInstance#user_data_base64}.
	// Experimental.
	UserDataBase64 *string `field:"optional" json:"userDataBase64" yaml:"userDataBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#user_data_replace_on_change TfInstance#user_data_replace_on_change}.
	// Experimental.
	UserDataReplaceOnChange interface{} `field:"optional" json:"userDataReplaceOnChange" yaml:"userDataReplaceOnChange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#volume_tags TfInstance#volume_tags}.
	// Experimental.
	VolumeTags *map[string]*string `field:"optional" json:"volumeTags" yaml:"volumeTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/instance#vpc_security_group_ids TfInstance#vpc_security_group_ids}.
	// Experimental.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

