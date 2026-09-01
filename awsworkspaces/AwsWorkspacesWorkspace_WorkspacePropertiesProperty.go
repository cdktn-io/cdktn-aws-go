package awsworkspaces


// Experimental.
type AwsWorkspacesWorkspace_WorkspacePropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_workspace#compute_type_name AwsWorkspacesWorkspace#compute_type_name}.
	// Experimental.
	ComputeTypeName *string `field:"optional" json:"computeTypeName" yaml:"computeTypeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_workspace#root_volume_size_gib AwsWorkspacesWorkspace#root_volume_size_gib}.
	// Experimental.
	RootVolumeSizeGib *float64 `field:"optional" json:"rootVolumeSizeGib" yaml:"rootVolumeSizeGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_workspace#running_mode AwsWorkspacesWorkspace#running_mode}.
	// Experimental.
	RunningMode *string `field:"optional" json:"runningMode" yaml:"runningMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_workspace#running_mode_auto_stop_timeout_in_minutes AwsWorkspacesWorkspace#running_mode_auto_stop_timeout_in_minutes}.
	// Experimental.
	RunningModeAutoStopTimeoutInMinutes *float64 `field:"optional" json:"runningModeAutoStopTimeoutInMinutes" yaml:"runningModeAutoStopTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_workspace#user_volume_size_gib AwsWorkspacesWorkspace#user_volume_size_gib}.
	// Experimental.
	UserVolumeSizeGib *float64 `field:"optional" json:"userVolumeSizeGib" yaml:"userVolumeSizeGib"`
}

