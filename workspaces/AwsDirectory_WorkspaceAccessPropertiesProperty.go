package workspaces


// Experimental.
type AwsDirectory_WorkspaceAccessPropertiesProperty struct {
	// access_endpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#access_endpoint_config AwsDirectory#access_endpoint_config}
	// Experimental.
	AccessEndpointConfig *AwsDirectory_AccessEndpointConfigProperty `field:"optional" json:"accessEndpointConfig" yaml:"accessEndpointConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#device_type_android AwsDirectory#device_type_android}.
	// Experimental.
	DeviceTypeAndroid *string `field:"optional" json:"deviceTypeAndroid" yaml:"deviceTypeAndroid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#device_type_chromeos AwsDirectory#device_type_chromeos}.
	// Experimental.
	DeviceTypeChromeos *string `field:"optional" json:"deviceTypeChromeos" yaml:"deviceTypeChromeos"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#device_type_ios AwsDirectory#device_type_ios}.
	// Experimental.
	DeviceTypeIos *string `field:"optional" json:"deviceTypeIos" yaml:"deviceTypeIos"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#device_type_linux AwsDirectory#device_type_linux}.
	// Experimental.
	DeviceTypeLinux *string `field:"optional" json:"deviceTypeLinux" yaml:"deviceTypeLinux"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#device_type_osx AwsDirectory#device_type_osx}.
	// Experimental.
	DeviceTypeOsx *string `field:"optional" json:"deviceTypeOsx" yaml:"deviceTypeOsx"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#device_type_web AwsDirectory#device_type_web}.
	// Experimental.
	DeviceTypeWeb *string `field:"optional" json:"deviceTypeWeb" yaml:"deviceTypeWeb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#device_type_windows AwsDirectory#device_type_windows}.
	// Experimental.
	DeviceTypeWindows *string `field:"optional" json:"deviceTypeWindows" yaml:"deviceTypeWindows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#device_type_zeroclient AwsDirectory#device_type_zeroclient}.
	// Experimental.
	DeviceTypeZeroclient *string `field:"optional" json:"deviceTypeZeroclient" yaml:"deviceTypeZeroclient"`
}

