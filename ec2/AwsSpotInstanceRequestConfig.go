package ec2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSpotInstanceRequestConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#ami AwsSpotInstanceRequest#ami}.
	// Experimental.
	Ami *string `field:"optional" json:"ami" yaml:"ami"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#associate_public_ip_address AwsSpotInstanceRequest#associate_public_ip_address}.
	// Experimental.
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#availability_zone AwsSpotInstanceRequest#availability_zone}.
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// capacity_reservation_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#capacity_reservation_specification AwsSpotInstanceRequest#capacity_reservation_specification}
	// Experimental.
	CapacityReservationSpecification *AwsSpotInstanceRequest_CapacityReservationSpecificationProperty `field:"optional" json:"capacityReservationSpecification" yaml:"capacityReservationSpecification"`
	// cpu_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#cpu_options AwsSpotInstanceRequest#cpu_options}
	// Experimental.
	CpuOptions *AwsSpotInstanceRequest_CpuOptionsProperty `field:"optional" json:"cpuOptions" yaml:"cpuOptions"`
	// credit_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#credit_specification AwsSpotInstanceRequest#credit_specification}
	// Experimental.
	CreditSpecification *AwsSpotInstanceRequest_CreditSpecificationProperty `field:"optional" json:"creditSpecification" yaml:"creditSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#disable_api_stop AwsSpotInstanceRequest#disable_api_stop}.
	// Experimental.
	DisableApiStop interface{} `field:"optional" json:"disableApiStop" yaml:"disableApiStop"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#disable_api_termination AwsSpotInstanceRequest#disable_api_termination}.
	// Experimental.
	DisableApiTermination interface{} `field:"optional" json:"disableApiTermination" yaml:"disableApiTermination"`
	// ebs_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#ebs_block_device AwsSpotInstanceRequest#ebs_block_device}
	// Experimental.
	EbsBlockDevice interface{} `field:"optional" json:"ebsBlockDevice" yaml:"ebsBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#ebs_optimized AwsSpotInstanceRequest#ebs_optimized}.
	// Experimental.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#enable_primary_ipv6 AwsSpotInstanceRequest#enable_primary_ipv6}.
	// Experimental.
	EnablePrimaryIpv6 interface{} `field:"optional" json:"enablePrimaryIpv6" yaml:"enablePrimaryIpv6"`
	// enclave_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#enclave_options AwsSpotInstanceRequest#enclave_options}
	// Experimental.
	EnclaveOptions *AwsSpotInstanceRequest_EnclaveOptionsProperty `field:"optional" json:"enclaveOptions" yaml:"enclaveOptions"`
	// ephemeral_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#ephemeral_block_device AwsSpotInstanceRequest#ephemeral_block_device}
	// Experimental.
	EphemeralBlockDevice interface{} `field:"optional" json:"ephemeralBlockDevice" yaml:"ephemeralBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#get_password_data AwsSpotInstanceRequest#get_password_data}.
	// Experimental.
	FetchPasswordData interface{} `field:"optional" json:"fetchPasswordData" yaml:"fetchPasswordData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#force_destroy AwsSpotInstanceRequest#force_destroy}.
	// Experimental.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#hibernation AwsSpotInstanceRequest#hibernation}.
	// Experimental.
	Hibernation interface{} `field:"optional" json:"hibernation" yaml:"hibernation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#host_id AwsSpotInstanceRequest#host_id}.
	// Experimental.
	HostId *string `field:"optional" json:"hostId" yaml:"hostId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#host_resource_group_arn AwsSpotInstanceRequest#host_resource_group_arn}.
	// Experimental.
	HostResourceGroupArn *string `field:"optional" json:"hostResourceGroupArn" yaml:"hostResourceGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#iam_instance_profile AwsSpotInstanceRequest#iam_instance_profile}.
	// Experimental.
	IamInstanceProfile *string `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#id AwsSpotInstanceRequest#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#instance_initiated_shutdown_behavior AwsSpotInstanceRequest#instance_initiated_shutdown_behavior}.
	// Experimental.
	InstanceInitiatedShutdownBehavior *string `field:"optional" json:"instanceInitiatedShutdownBehavior" yaml:"instanceInitiatedShutdownBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#instance_interruption_behavior AwsSpotInstanceRequest#instance_interruption_behavior}.
	// Experimental.
	InstanceInterruptionBehavior *string `field:"optional" json:"instanceInterruptionBehavior" yaml:"instanceInterruptionBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#instance_type AwsSpotInstanceRequest#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#ipv6_address_count AwsSpotInstanceRequest#ipv6_address_count}.
	// Experimental.
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#ipv6_addresses AwsSpotInstanceRequest#ipv6_addresses}.
	// Experimental.
	Ipv6Addresses *[]*string `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#key_name AwsSpotInstanceRequest#key_name}.
	// Experimental.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#launch_group AwsSpotInstanceRequest#launch_group}.
	// Experimental.
	LaunchGroup *string `field:"optional" json:"launchGroup" yaml:"launchGroup"`
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#launch_template AwsSpotInstanceRequest#launch_template}
	// Experimental.
	LaunchTemplate *AwsSpotInstanceRequest_LaunchTemplateProperty `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// maintenance_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#maintenance_options AwsSpotInstanceRequest#maintenance_options}
	// Experimental.
	MaintenanceOptions *AwsSpotInstanceRequest_MaintenanceOptionsProperty `field:"optional" json:"maintenanceOptions" yaml:"maintenanceOptions"`
	// metadata_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#metadata_options AwsSpotInstanceRequest#metadata_options}
	// Experimental.
	MetadataOptions *AwsSpotInstanceRequest_MetadataOptionsProperty `field:"optional" json:"metadataOptions" yaml:"metadataOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#monitoring AwsSpotInstanceRequest#monitoring}.
	// Experimental.
	Monitoring interface{} `field:"optional" json:"monitoring" yaml:"monitoring"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#network_interface AwsSpotInstanceRequest#network_interface}
	// Experimental.
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#placement_group AwsSpotInstanceRequest#placement_group}.
	// Experimental.
	PlacementGroup *string `field:"optional" json:"placementGroup" yaml:"placementGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#placement_group_id AwsSpotInstanceRequest#placement_group_id}.
	// Experimental.
	PlacementGroupId *string `field:"optional" json:"placementGroupId" yaml:"placementGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#placement_partition_number AwsSpotInstanceRequest#placement_partition_number}.
	// Experimental.
	PlacementPartitionNumber *float64 `field:"optional" json:"placementPartitionNumber" yaml:"placementPartitionNumber"`
	// private_dns_name_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#private_dns_name_options AwsSpotInstanceRequest#private_dns_name_options}
	// Experimental.
	PrivateDnsNameOptions *AwsSpotInstanceRequest_PrivateDnsNameOptionsProperty `field:"optional" json:"privateDnsNameOptions" yaml:"privateDnsNameOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#private_ip AwsSpotInstanceRequest#private_ip}.
	// Experimental.
	PrivateIp *string `field:"optional" json:"privateIp" yaml:"privateIp"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#region AwsSpotInstanceRequest#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// root_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#root_block_device AwsSpotInstanceRequest#root_block_device}
	// Experimental.
	RootBlockDevice *AwsSpotInstanceRequest_RootBlockDeviceProperty `field:"optional" json:"rootBlockDevice" yaml:"rootBlockDevice"`
	// secondary_network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#secondary_network_interface AwsSpotInstanceRequest#secondary_network_interface}
	// Experimental.
	SecondaryNetworkInterface interface{} `field:"optional" json:"secondaryNetworkInterface" yaml:"secondaryNetworkInterface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#secondary_private_ips AwsSpotInstanceRequest#secondary_private_ips}.
	// Experimental.
	SecondaryPrivateIps *[]*string `field:"optional" json:"secondaryPrivateIps" yaml:"secondaryPrivateIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#security_groups AwsSpotInstanceRequest#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#source_dest_check AwsSpotInstanceRequest#source_dest_check}.
	// Experimental.
	SourceDestCheck interface{} `field:"optional" json:"sourceDestCheck" yaml:"sourceDestCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#spot_price AwsSpotInstanceRequest#spot_price}.
	// Experimental.
	SpotPrice *string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#spot_type AwsSpotInstanceRequest#spot_type}.
	// Experimental.
	SpotType *string `field:"optional" json:"spotType" yaml:"spotType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#subnet_id AwsSpotInstanceRequest#subnet_id}.
	// Experimental.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#tags AwsSpotInstanceRequest#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#tags_all AwsSpotInstanceRequest#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#tenancy AwsSpotInstanceRequest#tenancy}.
	// Experimental.
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#timeouts AwsSpotInstanceRequest#timeouts}
	// Experimental.
	Timeouts *AwsSpotInstanceRequest_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#user_data AwsSpotInstanceRequest#user_data}.
	// Experimental.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#user_data_base64 AwsSpotInstanceRequest#user_data_base64}.
	// Experimental.
	UserDataBase64 *string `field:"optional" json:"userDataBase64" yaml:"userDataBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#user_data_replace_on_change AwsSpotInstanceRequest#user_data_replace_on_change}.
	// Experimental.
	UserDataReplaceOnChange interface{} `field:"optional" json:"userDataReplaceOnChange" yaml:"userDataReplaceOnChange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#valid_from AwsSpotInstanceRequest#valid_from}.
	// Experimental.
	ValidFrom *string `field:"optional" json:"validFrom" yaml:"validFrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#valid_until AwsSpotInstanceRequest#valid_until}.
	// Experimental.
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#volume_tags AwsSpotInstanceRequest#volume_tags}.
	// Experimental.
	VolumeTags *map[string]*string `field:"optional" json:"volumeTags" yaml:"volumeTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#vpc_security_group_ids AwsSpotInstanceRequest#vpc_security_group_ids}.
	// Experimental.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#wait_for_fulfillment AwsSpotInstanceRequest#wait_for_fulfillment}.
	// Experimental.
	WaitForFulfillment interface{} `field:"optional" json:"waitForFulfillment" yaml:"waitForFulfillment"`
}

