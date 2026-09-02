package awsautoscaling


// Experimental.
type TfLaunchConfiguration_RootBlockDeviceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_configuration#delete_on_termination TfLaunchConfiguration#delete_on_termination}.
	// Experimental.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_configuration#encrypted TfLaunchConfiguration#encrypted}.
	// Experimental.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_configuration#iops TfLaunchConfiguration#iops}.
	// Experimental.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_configuration#throughput TfLaunchConfiguration#throughput}.
	// Experimental.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_configuration#volume_size TfLaunchConfiguration#volume_size}.
	// Experimental.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_configuration#volume_type TfLaunchConfiguration#volume_type}.
	// Experimental.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

