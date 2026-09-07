package ecs


// Experimental.
type AwsTaskDefinition_VolumeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#name AwsTaskDefinition#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#configure_at_launch AwsTaskDefinition#configure_at_launch}.
	// Experimental.
	ConfigureAtLaunch interface{} `field:"optional" json:"configureAtLaunch" yaml:"configureAtLaunch"`
	// docker_volume_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#docker_volume_configuration AwsTaskDefinition#docker_volume_configuration}
	// Experimental.
	DockerVolumeConfiguration *AwsTaskDefinition_DockerVolumeConfigurationProperty `field:"optional" json:"dockerVolumeConfiguration" yaml:"dockerVolumeConfiguration"`
	// efs_volume_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#efs_volume_configuration AwsTaskDefinition#efs_volume_configuration}
	// Experimental.
	EfsVolumeConfiguration *AwsTaskDefinition_EfsVolumeConfigurationProperty `field:"optional" json:"efsVolumeConfiguration" yaml:"efsVolumeConfiguration"`
	// fsx_windows_file_server_volume_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#fsx_windows_file_server_volume_configuration AwsTaskDefinition#fsx_windows_file_server_volume_configuration}
	// Experimental.
	FsxWindowsFileServerVolumeConfiguration *AwsTaskDefinition_FsxWindowsFileServerVolumeConfigurationProperty `field:"optional" json:"fsxWindowsFileServerVolumeConfiguration" yaml:"fsxWindowsFileServerVolumeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#host_path AwsTaskDefinition#host_path}.
	// Experimental.
	HostPath *string `field:"optional" json:"hostPath" yaml:"hostPath"`
	// s3files_volume_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_task_definition#s3files_volume_configuration AwsTaskDefinition#s3files_volume_configuration}
	// Experimental.
	S3FilesVolumeConfiguration *AwsTaskDefinition_S3filesVolumeConfigurationProperty `field:"optional" json:"s3FilesVolumeConfiguration" yaml:"s3FilesVolumeConfiguration"`
}

