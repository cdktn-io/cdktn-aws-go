package awsstoragegateway

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGatewayConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#gateway_name TfGateway#gateway_name}.
	// Experimental.
	GatewayName *string `field:"required" json:"gatewayName" yaml:"gatewayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#gateway_timezone TfGateway#gateway_timezone}.
	// Experimental.
	GatewayTimezone *string `field:"required" json:"gatewayTimezone" yaml:"gatewayTimezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#activation_key TfGateway#activation_key}.
	// Experimental.
	ActivationKey *string `field:"optional" json:"activationKey" yaml:"activationKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#average_download_rate_limit_in_bits_per_sec TfGateway#average_download_rate_limit_in_bits_per_sec}.
	// Experimental.
	AverageDownloadRateLimitInBitsPerSec *float64 `field:"optional" json:"averageDownloadRateLimitInBitsPerSec" yaml:"averageDownloadRateLimitInBitsPerSec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#average_upload_rate_limit_in_bits_per_sec TfGateway#average_upload_rate_limit_in_bits_per_sec}.
	// Experimental.
	AverageUploadRateLimitInBitsPerSec *float64 `field:"optional" json:"averageUploadRateLimitInBitsPerSec" yaml:"averageUploadRateLimitInBitsPerSec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#cloudwatch_log_group_arn TfGateway#cloudwatch_log_group_arn}.
	// Experimental.
	CloudwatchLogGroupArn *string `field:"optional" json:"cloudwatchLogGroupArn" yaml:"cloudwatchLogGroupArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#gateway_ip_address TfGateway#gateway_ip_address}.
	// Experimental.
	GatewayIpAddress *string `field:"optional" json:"gatewayIpAddress" yaml:"gatewayIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#gateway_type TfGateway#gateway_type}.
	// Experimental.
	GatewayType *string `field:"optional" json:"gatewayType" yaml:"gatewayType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#gateway_vpc_endpoint TfGateway#gateway_vpc_endpoint}.
	// Experimental.
	GatewayVpcEndpoint *string `field:"optional" json:"gatewayVpcEndpoint" yaml:"gatewayVpcEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#id TfGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// maintenance_start_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#maintenance_start_time TfGateway#maintenance_start_time}
	// Experimental.
	MaintenanceStartTime *TfGateway_MaintenanceStartTimeProperty `field:"optional" json:"maintenanceStartTime" yaml:"maintenanceStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#medium_changer_type TfGateway#medium_changer_type}.
	// Experimental.
	MediumChangerType *string `field:"optional" json:"mediumChangerType" yaml:"mediumChangerType"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#region TfGateway#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// smb_active_directory_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#smb_active_directory_settings TfGateway#smb_active_directory_settings}
	// Experimental.
	SmbActiveDirectorySettings *TfGateway_SmbActiveDirectorySettingsProperty `field:"optional" json:"smbActiveDirectorySettings" yaml:"smbActiveDirectorySettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#smb_file_share_visibility TfGateway#smb_file_share_visibility}.
	// Experimental.
	SmbFileShareVisibility interface{} `field:"optional" json:"smbFileShareVisibility" yaml:"smbFileShareVisibility"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#smb_guest_password TfGateway#smb_guest_password}.
	// Experimental.
	SmbGuestPassword *string `field:"optional" json:"smbGuestPassword" yaml:"smbGuestPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#smb_security_strategy TfGateway#smb_security_strategy}.
	// Experimental.
	SmbSecurityStrategy *string `field:"optional" json:"smbSecurityStrategy" yaml:"smbSecurityStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#tags TfGateway#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#tags_all TfGateway#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#tape_drive_type TfGateway#tape_drive_type}.
	// Experimental.
	TapeDriveType *string `field:"optional" json:"tapeDriveType" yaml:"tapeDriveType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_gateway#timeouts TfGateway#timeouts}
	// Experimental.
	Timeouts *TfGateway_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

