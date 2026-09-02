package awsec2


// Experimental.
type TfFleet_FleetInstanceSetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#instance_ids TfFleet#instance_ids}.
	// Experimental.
	InstanceIds *[]*string `field:"optional" json:"instanceIds" yaml:"instanceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#instance_type TfFleet#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#lifecycle TfFleet#lifecycle}.
	// Experimental.
	Lifecycle *string `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_fleet#platform TfFleet#platform}.
	// Experimental.
	Platform *string `field:"optional" json:"platform" yaml:"platform"`
}

