package awsec2


// Experimental.
type TfSpotFleetRequest_RootBlockDeviceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#delete_on_termination TfSpotFleetRequest#delete_on_termination}.
	// Experimental.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#encrypted TfSpotFleetRequest#encrypted}.
	// Experimental.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#iops TfSpotFleetRequest#iops}.
	// Experimental.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#kms_key_id TfSpotFleetRequest#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#throughput TfSpotFleetRequest#throughput}.
	// Experimental.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#volume_size TfSpotFleetRequest#volume_size}.
	// Experimental.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#volume_type TfSpotFleetRequest#volume_type}.
	// Experimental.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

