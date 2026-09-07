package workspaces


// Experimental.
type AwsDirectory_SelfServicePermissionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#change_compute_type AwsDirectory#change_compute_type}.
	// Experimental.
	ChangeComputeType interface{} `field:"optional" json:"changeComputeType" yaml:"changeComputeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#increase_volume_size AwsDirectory#increase_volume_size}.
	// Experimental.
	IncreaseVolumeSize interface{} `field:"optional" json:"increaseVolumeSize" yaml:"increaseVolumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#rebuild_workspace AwsDirectory#rebuild_workspace}.
	// Experimental.
	RebuildWorkspace interface{} `field:"optional" json:"rebuildWorkspace" yaml:"rebuildWorkspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#restart_workspace AwsDirectory#restart_workspace}.
	// Experimental.
	RestartWorkspace interface{} `field:"optional" json:"restartWorkspace" yaml:"restartWorkspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#switch_running_mode AwsDirectory#switch_running_mode}.
	// Experimental.
	SwitchRunningMode interface{} `field:"optional" json:"switchRunningMode" yaml:"switchRunningMode"`
}

