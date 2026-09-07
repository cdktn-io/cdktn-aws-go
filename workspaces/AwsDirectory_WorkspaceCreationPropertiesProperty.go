package workspaces


// Experimental.
type AwsDirectory_WorkspaceCreationPropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#custom_security_group_id AwsDirectory#custom_security_group_id}.
	// Experimental.
	CustomSecurityGroupId *string `field:"optional" json:"customSecurityGroupId" yaml:"customSecurityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#default_ou AwsDirectory#default_ou}.
	// Experimental.
	DefaultOu *string `field:"optional" json:"defaultOu" yaml:"defaultOu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#enable_internet_access AwsDirectory#enable_internet_access}.
	// Experimental.
	EnableInternetAccess interface{} `field:"optional" json:"enableInternetAccess" yaml:"enableInternetAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#enable_maintenance_mode AwsDirectory#enable_maintenance_mode}.
	// Experimental.
	EnableMaintenanceMode interface{} `field:"optional" json:"enableMaintenanceMode" yaml:"enableMaintenanceMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#user_enabled_as_local_administrator AwsDirectory#user_enabled_as_local_administrator}.
	// Experimental.
	UserEnabledAsLocalAdministrator interface{} `field:"optional" json:"userEnabledAsLocalAdministrator" yaml:"userEnabledAsLocalAdministrator"`
}

