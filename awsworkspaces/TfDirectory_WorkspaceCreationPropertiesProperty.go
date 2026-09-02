package awsworkspaces


// Experimental.
type TfDirectory_WorkspaceCreationPropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#custom_security_group_id TfDirectory#custom_security_group_id}.
	// Experimental.
	CustomSecurityGroupId *string `field:"optional" json:"customSecurityGroupId" yaml:"customSecurityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#default_ou TfDirectory#default_ou}.
	// Experimental.
	DefaultOu *string `field:"optional" json:"defaultOu" yaml:"defaultOu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#enable_internet_access TfDirectory#enable_internet_access}.
	// Experimental.
	EnableInternetAccess interface{} `field:"optional" json:"enableInternetAccess" yaml:"enableInternetAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#enable_maintenance_mode TfDirectory#enable_maintenance_mode}.
	// Experimental.
	EnableMaintenanceMode interface{} `field:"optional" json:"enableMaintenanceMode" yaml:"enableMaintenanceMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#user_enabled_as_local_administrator TfDirectory#user_enabled_as_local_administrator}.
	// Experimental.
	UserEnabledAsLocalAdministrator interface{} `field:"optional" json:"userEnabledAsLocalAdministrator" yaml:"userEnabledAsLocalAdministrator"`
}

