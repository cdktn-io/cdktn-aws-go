package awsec2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSpotInstanceRequestConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#ami TfSpotInstanceRequest#ami}.
	// Experimental.
	Ami *string `field:"optional" json:"ami" yaml:"ami"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#associate_public_ip_address TfSpotInstanceRequest#associate_public_ip_address}.
	// Experimental.
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#availability_zone TfSpotInstanceRequest#availability_zone}.
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// capacity_reservation_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#capacity_reservation_specification TfSpotInstanceRequest#capacity_reservation_specification}
	// Experimental.
	CapacityReservationSpecification *TfSpotInstanceRequest_CapacityReservationSpecificationProperty `field:"optional" json:"capacityReservationSpecification" yaml:"capacityReservationSpecification"`
	// cpu_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#cpu_options TfSpotInstanceRequest#cpu_options}
	// Experimental.
	CpuOptions *TfSpotInstanceRequest_CpuOptionsProperty `field:"optional" json:"cpuOptions" yaml:"cpuOptions"`
	// credit_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#credit_specification TfSpotInstanceRequest#credit_specification}
	// Experimental.
	CreditSpecification *TfSpotInstanceRequest_CreditSpecificationProperty `field:"optional" json:"creditSpecification" yaml:"creditSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#disable_api_stop TfSpotInstanceRequest#disable_api_stop}.
	// Experimental.
	DisableApiStop interface{} `field:"optional" json:"disableApiStop" yaml:"disableApiStop"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#disable_api_termination TfSpotInstanceRequest#disable_api_termination}.
	// Experimental.
	DisableApiTermination interface{} `field:"optional" json:"disableApiTermination" yaml:"disableApiTermination"`
	// ebs_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#ebs_block_device TfSpotInstanceRequest#ebs_block_device}
	// Experimental.
	EbsBlockDevice interface{} `field:"optional" json:"ebsBlockDevice" yaml:"ebsBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#ebs_optimized TfSpotInstanceRequest#ebs_optimized}.
	// Experimental.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#enable_primary_ipv6 TfSpotInstanceRequest#enable_primary_ipv6}.
	// Experimental.
	EnablePrimaryIpv6 interface{} `field:"optional" json:"enablePrimaryIpv6" yaml:"enablePrimaryIpv6"`
	// enclave_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#enclave_options TfSpotInstanceRequest#enclave_options}
	// Experimental.
	EnclaveOptions *TfSpotInstanceRequest_EnclaveOptionsProperty `field:"optional" json:"enclaveOptions" yaml:"enclaveOptions"`
	// ephemeral_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#ephemeral_block_device TfSpotInstanceRequest#ephemeral_block_device}
	// Experimental.
	EphemeralBlockDevice interface{} `field:"optional" json:"ephemeralBlockDevice" yaml:"ephemeralBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#get_password_data TfSpotInstanceRequest#get_password_data}.
	// Experimental.
	FetchPasswordData interface{} `field:"optional" json:"fetchPasswordData" yaml:"fetchPasswordData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#force_destroy TfSpotInstanceRequest#force_destroy}.
	// Experimental.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#hibernation TfSpotInstanceRequest#hibernation}.
	// Experimental.
	Hibernation interface{} `field:"optional" json:"hibernation" yaml:"hibernation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#host_id TfSpotInstanceRequest#host_id}.
	// Experimental.
	HostId *string `field:"optional" json:"hostId" yaml:"hostId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#host_resource_group_arn TfSpotInstanceRequest#host_resource_group_arn}.
	// Experimental.
	HostResourceGroupArn *string `field:"optional" json:"hostResourceGroupArn" yaml:"hostResourceGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#iam_instance_profile TfSpotInstanceRequest#iam_instance_profile}.
	// Experimental.
	IamInstanceProfile *string `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#id TfSpotInstanceRequest#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#instance_initiated_shutdown_behavior TfSpotInstanceRequest#instance_initiated_shutdown_behavior}.
	// Experimental.
	InstanceInitiatedShutdownBehavior *string `field:"optional" json:"instanceInitiatedShutdownBehavior" yaml:"instanceInitiatedShutdownBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#instance_interruption_behavior TfSpotInstanceRequest#instance_interruption_behavior}.
	// Experimental.
	InstanceInterruptionBehavior *string `field:"optional" json:"instanceInterruptionBehavior" yaml:"instanceInterruptionBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#instance_type TfSpotInstanceRequest#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#ipv6_address_count TfSpotInstanceRequest#ipv6_address_count}.
	// Experimental.
	Ipv6AddressCount *float64 `field:"optional" json:"ipv6AddressCount" yaml:"ipv6AddressCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#ipv6_addresses TfSpotInstanceRequest#ipv6_addresses}.
	// Experimental.
	Ipv6Addresses *[]*string `field:"optional" json:"ipv6Addresses" yaml:"ipv6Addresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#key_name TfSpotInstanceRequest#key_name}.
	// Experimental.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#launch_group TfSpotInstanceRequest#launch_group}.
	// Experimental.
	LaunchGroup *string `field:"optional" json:"launchGroup" yaml:"launchGroup"`
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#launch_template TfSpotInstanceRequest#launch_template}
	// Experimental.
	LaunchTemplate *TfSpotInstanceRequest_LaunchTemplateProperty `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// maintenance_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#maintenance_options TfSpotInstanceRequest#maintenance_options}
	// Experimental.
	MaintenanceOptions *TfSpotInstanceRequest_MaintenanceOptionsProperty `field:"optional" json:"maintenanceOptions" yaml:"maintenanceOptions"`
	// metadata_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#metadata_options TfSpotInstanceRequest#metadata_options}
	// Experimental.
	MetadataOptions *TfSpotInstanceRequest_MetadataOptionsProperty `field:"optional" json:"metadataOptions" yaml:"metadataOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#monitoring TfSpotInstanceRequest#monitoring}.
	// Experimental.
	Monitoring interface{} `field:"optional" json:"monitoring" yaml:"monitoring"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#network_interface TfSpotInstanceRequest#network_interface}
	// Experimental.
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#placement_group TfSpotInstanceRequest#placement_group}.
	// Experimental.
	PlacementGroup *string `field:"optional" json:"placementGroup" yaml:"placementGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#placement_group_id TfSpotInstanceRequest#placement_group_id}.
	// Experimental.
	PlacementGroupId *string `field:"optional" json:"placementGroupId" yaml:"placementGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#placement_partition_number TfSpotInstanceRequest#placement_partition_number}.
	// Experimental.
	PlacementPartitionNumber *float64 `field:"optional" json:"placementPartitionNumber" yaml:"placementPartitionNumber"`
	// private_dns_name_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#private_dns_name_options TfSpotInstanceRequest#private_dns_name_options}
	// Experimental.
	PrivateDnsNameOptions *TfSpotInstanceRequest_PrivateDnsNameOptionsProperty `field:"optional" json:"privateDnsNameOptions" yaml:"privateDnsNameOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#private_ip TfSpotInstanceRequest#private_ip}.
	// Experimental.
	PrivateIp *string `field:"optional" json:"privateIp" yaml:"privateIp"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#region TfSpotInstanceRequest#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// root_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#root_block_device TfSpotInstanceRequest#root_block_device}
	// Experimental.
	RootBlockDevice *TfSpotInstanceRequest_RootBlockDeviceProperty `field:"optional" json:"rootBlockDevice" yaml:"rootBlockDevice"`
	// secondary_network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#secondary_network_interface TfSpotInstanceRequest#secondary_network_interface}
	// Experimental.
	SecondaryNetworkInterface interface{} `field:"optional" json:"secondaryNetworkInterface" yaml:"secondaryNetworkInterface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#secondary_private_ips TfSpotInstanceRequest#secondary_private_ips}.
	// Experimental.
	SecondaryPrivateIps *[]*string `field:"optional" json:"secondaryPrivateIps" yaml:"secondaryPrivateIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#security_groups TfSpotInstanceRequest#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#source_dest_check TfSpotInstanceRequest#source_dest_check}.
	// Experimental.
	SourceDestCheck interface{} `field:"optional" json:"sourceDestCheck" yaml:"sourceDestCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#spot_price TfSpotInstanceRequest#spot_price}.
	// Experimental.
	SpotPrice *string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#spot_type TfSpotInstanceRequest#spot_type}.
	// Experimental.
	SpotType *string `field:"optional" json:"spotType" yaml:"spotType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#subnet_id TfSpotInstanceRequest#subnet_id}.
	// Experimental.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#tags TfSpotInstanceRequest#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#tags_all TfSpotInstanceRequest#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#tenancy TfSpotInstanceRequest#tenancy}.
	// Experimental.
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#timeouts TfSpotInstanceRequest#timeouts}
	// Experimental.
	Timeouts *TfSpotInstanceRequest_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#user_data TfSpotInstanceRequest#user_data}.
	// Experimental.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#user_data_base64 TfSpotInstanceRequest#user_data_base64}.
	// Experimental.
	UserDataBase64 *string `field:"optional" json:"userDataBase64" yaml:"userDataBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#user_data_replace_on_change TfSpotInstanceRequest#user_data_replace_on_change}.
	// Experimental.
	UserDataReplaceOnChange interface{} `field:"optional" json:"userDataReplaceOnChange" yaml:"userDataReplaceOnChange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#valid_from TfSpotInstanceRequest#valid_from}.
	// Experimental.
	ValidFrom *string `field:"optional" json:"validFrom" yaml:"validFrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#valid_until TfSpotInstanceRequest#valid_until}.
	// Experimental.
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#volume_tags TfSpotInstanceRequest#volume_tags}.
	// Experimental.
	VolumeTags *map[string]*string `field:"optional" json:"volumeTags" yaml:"volumeTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#vpc_security_group_ids TfSpotInstanceRequest#vpc_security_group_ids}.
	// Experimental.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_instance_request#wait_for_fulfillment TfSpotInstanceRequest#wait_for_fulfillment}.
	// Experimental.
	WaitForFulfillment interface{} `field:"optional" json:"waitForFulfillment" yaml:"waitForFulfillment"`
}

