package awsec2


// Experimental.
type TfFleet_OverrideProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#availability_zone TfFleet#availability_zone}.
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// instance_requirements block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#instance_requirements TfFleet#instance_requirements}
	// Experimental.
	InstanceRequirements *TfFleet_InstanceRequirementsProperty `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#instance_type TfFleet#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#max_price TfFleet#max_price}.
	// Experimental.
	MaxPrice *string `field:"optional" json:"maxPrice" yaml:"maxPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#priority TfFleet#priority}.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#subnet_id TfFleet#subnet_id}.
	// Experimental.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#weighted_capacity TfFleet#weighted_capacity}.
	// Experimental.
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}

