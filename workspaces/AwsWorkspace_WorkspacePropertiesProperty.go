package workspaces


// Experimental.
type AwsWorkspace_WorkspacePropertiesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_workspace#compute_type_name AwsWorkspace#compute_type_name}.
	// Experimental.
	ComputeTypeName *string `field:"optional" json:"computeTypeName" yaml:"computeTypeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_workspace#root_volume_size_gib AwsWorkspace#root_volume_size_gib}.
	// Experimental.
	RootVolumeSizeGib *float64 `field:"optional" json:"rootVolumeSizeGib" yaml:"rootVolumeSizeGib"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_workspace#running_mode AwsWorkspace#running_mode}.
	// Experimental.
	RunningMode *string `field:"optional" json:"runningMode" yaml:"runningMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_workspace#running_mode_auto_stop_timeout_in_minutes AwsWorkspace#running_mode_auto_stop_timeout_in_minutes}.
	// Experimental.
	RunningModeAutoStopTimeoutInMinutes *float64 `field:"optional" json:"runningModeAutoStopTimeoutInMinutes" yaml:"runningModeAutoStopTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_workspace#user_volume_size_gib AwsWorkspace#user_volume_size_gib}.
	// Experimental.
	UserVolumeSizeGib *float64 `field:"optional" json:"userVolumeSizeGib" yaml:"userVolumeSizeGib"`
}

