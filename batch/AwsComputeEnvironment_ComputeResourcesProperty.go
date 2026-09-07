package batch


// Experimental.
type AwsComputeEnvironment_ComputeResourcesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#max_vcpus AwsComputeEnvironment#max_vcpus}.
	// Experimental.
	MaxVcpus *float64 `field:"required" json:"maxVcpus" yaml:"maxVcpus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#subnets AwsComputeEnvironment#subnets}.
	// Experimental.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#type AwsComputeEnvironment#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#allocation_strategy AwsComputeEnvironment#allocation_strategy}.
	// Experimental.
	AllocationStrategy *string `field:"optional" json:"allocationStrategy" yaml:"allocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#bid_percentage AwsComputeEnvironment#bid_percentage}.
	// Experimental.
	BidPercentage *float64 `field:"optional" json:"bidPercentage" yaml:"bidPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#desired_vcpus AwsComputeEnvironment#desired_vcpus}.
	// Experimental.
	DesiredVcpus *float64 `field:"optional" json:"desiredVcpus" yaml:"desiredVcpus"`
	// ec2_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#ec2_configuration AwsComputeEnvironment#ec2_configuration}
	// Experimental.
	Ec2Configuration interface{} `field:"optional" json:"ec2Configuration" yaml:"ec2Configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#ec2_key_pair AwsComputeEnvironment#ec2_key_pair}.
	// Experimental.
	Ec2KeyPair *string `field:"optional" json:"ec2KeyPair" yaml:"ec2KeyPair"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#image_id AwsComputeEnvironment#image_id}.
	// Experimental.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#instance_role AwsComputeEnvironment#instance_role}.
	// Experimental.
	InstanceRole *string `field:"optional" json:"instanceRole" yaml:"instanceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#instance_type AwsComputeEnvironment#instance_type}.
	// Experimental.
	InstanceType *[]*string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#launch_template AwsComputeEnvironment#launch_template}
	// Experimental.
	LaunchTemplate *AwsComputeEnvironment_LaunchTemplateProperty `field:"optional" json:"launchTemplate" yaml:"launchTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#min_vcpus AwsComputeEnvironment#min_vcpus}.
	// Experimental.
	MinVcpus *float64 `field:"optional" json:"minVcpus" yaml:"minVcpus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#placement_group AwsComputeEnvironment#placement_group}.
	// Experimental.
	PlacementGroup *string `field:"optional" json:"placementGroup" yaml:"placementGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#security_group_ids AwsComputeEnvironment#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#spot_iam_fleet_role AwsComputeEnvironment#spot_iam_fleet_role}.
	// Experimental.
	SpotIamFleetRole *string `field:"optional" json:"spotIamFleetRole" yaml:"spotIamFleetRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#tags AwsComputeEnvironment#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

