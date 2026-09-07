package ec2


// Experimental.
type AwsFleet_FleetInstanceSetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#instance_ids AwsFleet#instance_ids}.
	// Experimental.
	InstanceIds *[]*string `field:"optional" json:"instanceIds" yaml:"instanceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#instance_type AwsFleet#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#lifecycle AwsFleet#lifecycle}.
	// Experimental.
	Lifecycle *string `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#platform AwsFleet#platform}.
	// Experimental.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
}

