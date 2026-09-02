package awsec2


// Experimental.
type TfSpotFleetRequest_LaunchSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#ami TfSpotFleetRequest#ami}.
	// Experimental.
	Ami *string `field:"required" json:"ami" yaml:"ami"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#instance_type TfSpotFleetRequest#instance_type}.
	// Experimental.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#associate_public_ip_address TfSpotFleetRequest#associate_public_ip_address}.
	// Experimental.
	AssociatePublicIpAddress interface{} `field:"optional" json:"associatePublicIpAddress" yaml:"associatePublicIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#availability_zone TfSpotFleetRequest#availability_zone}.
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// ebs_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#ebs_block_device TfSpotFleetRequest#ebs_block_device}
	// Experimental.
	EbsBlockDevice interface{} `field:"optional" json:"ebsBlockDevice" yaml:"ebsBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#ebs_optimized TfSpotFleetRequest#ebs_optimized}.
	// Experimental.
	EbsOptimized interface{} `field:"optional" json:"ebsOptimized" yaml:"ebsOptimized"`
	// ephemeral_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#ephemeral_block_device TfSpotFleetRequest#ephemeral_block_device}
	// Experimental.
	EphemeralBlockDevice interface{} `field:"optional" json:"ephemeralBlockDevice" yaml:"ephemeralBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#iam_instance_profile TfSpotFleetRequest#iam_instance_profile}.
	// Experimental.
	IamInstanceProfile *string `field:"optional" json:"iamInstanceProfile" yaml:"iamInstanceProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#iam_instance_profile_arn TfSpotFleetRequest#iam_instance_profile_arn}.
	// Experimental.
	IamInstanceProfileArn *string `field:"optional" json:"iamInstanceProfileArn" yaml:"iamInstanceProfileArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#key_name TfSpotFleetRequest#key_name}.
	// Experimental.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#monitoring TfSpotFleetRequest#monitoring}.
	// Experimental.
	Monitoring interface{} `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#placement_group TfSpotFleetRequest#placement_group}.
	// Experimental.
	PlacementGroup *string `field:"optional" json:"placementGroup" yaml:"placementGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#placement_tenancy TfSpotFleetRequest#placement_tenancy}.
	// Experimental.
	PlacementTenancy *string `field:"optional" json:"placementTenancy" yaml:"placementTenancy"`
	// root_block_device block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#root_block_device TfSpotFleetRequest#root_block_device}
	// Experimental.
	RootBlockDevice interface{} `field:"optional" json:"rootBlockDevice" yaml:"rootBlockDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#spot_price TfSpotFleetRequest#spot_price}.
	// Experimental.
	SpotPrice *string `field:"optional" json:"spotPrice" yaml:"spotPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#subnet_id TfSpotFleetRequest#subnet_id}.
	// Experimental.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#tags TfSpotFleetRequest#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#user_data TfSpotFleetRequest#user_data}.
	// Experimental.
	UserData *string `field:"optional" json:"userData" yaml:"userData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#vpc_security_group_ids TfSpotFleetRequest#vpc_security_group_ids}.
	// Experimental.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/spot_fleet_request#weighted_capacity TfSpotFleetRequest#weighted_capacity}.
	// Experimental.
	WeightedCapacity *string `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}

