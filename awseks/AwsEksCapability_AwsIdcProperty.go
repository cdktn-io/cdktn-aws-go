package awseks


// Experimental.
type AwsEksCapability_AwsIdcProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_capability#idc_instance_arn AwsEksCapability#idc_instance_arn}.
	// Experimental.
	IdcInstanceArn *string `field:"required" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_capability#idc_region AwsEksCapability#idc_region}.
	// Experimental.
	IdcRegion *string `field:"optional" json:"idcRegion" yaml:"idcRegion"`
}

