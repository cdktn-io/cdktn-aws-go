package awsworkspaces


// Experimental.
type AwsWorkspacesDirectory_WorkspaceAccessPropertiesProperty struct {
	// access_endpoint_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#access_endpoint_config AwsWorkspacesDirectory#access_endpoint_config}
	// Experimental.
	AccessEndpointConfig *AwsWorkspacesDirectory_AccessEndpointConfigProperty `field:"optional" json:"accessEndpointConfig" yaml:"accessEndpointConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#device_type_android AwsWorkspacesDirectory#device_type_android}.
	// Experimental.
	DeviceTypeAndroid *string `field:"optional" json:"deviceTypeAndroid" yaml:"deviceTypeAndroid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#device_type_chromeos AwsWorkspacesDirectory#device_type_chromeos}.
	// Experimental.
	DeviceTypeChromeos *string `field:"optional" json:"deviceTypeChromeos" yaml:"deviceTypeChromeos"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#device_type_ios AwsWorkspacesDirectory#device_type_ios}.
	// Experimental.
	DeviceTypeIos *string `field:"optional" json:"deviceTypeIos" yaml:"deviceTypeIos"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#device_type_linux AwsWorkspacesDirectory#device_type_linux}.
	// Experimental.
	DeviceTypeLinux *string `field:"optional" json:"deviceTypeLinux" yaml:"deviceTypeLinux"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#device_type_osx AwsWorkspacesDirectory#device_type_osx}.
	// Experimental.
	DeviceTypeOsx *string `field:"optional" json:"deviceTypeOsx" yaml:"deviceTypeOsx"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#device_type_web AwsWorkspacesDirectory#device_type_web}.
	// Experimental.
	DeviceTypeWeb *string `field:"optional" json:"deviceTypeWeb" yaml:"deviceTypeWeb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#device_type_windows AwsWorkspacesDirectory#device_type_windows}.
	// Experimental.
	DeviceTypeWindows *string `field:"optional" json:"deviceTypeWindows" yaml:"deviceTypeWindows"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#device_type_zeroclient AwsWorkspacesDirectory#device_type_zeroclient}.
	// Experimental.
	DeviceTypeZeroclient *string `field:"optional" json:"deviceTypeZeroclient" yaml:"deviceTypeZeroclient"`
}

