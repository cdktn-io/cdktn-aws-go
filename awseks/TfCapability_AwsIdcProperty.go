package awseks


// Experimental.
type TfCapability_AwsIdcProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_capability#idc_instance_arn TfCapability#idc_instance_arn}.
	// Experimental.
	IdcInstanceArn *string `field:"required" json:"idcInstanceArn" yaml:"idcInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_capability#idc_region TfCapability#idc_region}.
	// Experimental.
	IdcRegion *string `field:"optional" json:"idcRegion" yaml:"idcRegion"`
}

