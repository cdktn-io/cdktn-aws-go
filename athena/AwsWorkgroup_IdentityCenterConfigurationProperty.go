package athena


// Experimental.
type AwsWorkgroup_IdentityCenterConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#enable_identity_center AwsWorkgroup#enable_identity_center}.
	// Experimental.
	EnableIdentityCenter interface{} `field:"optional" json:"enableIdentityCenter" yaml:"enableIdentityCenter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/athena_workgroup#identity_center_instance_arn AwsWorkgroup#identity_center_instance_arn}.
	// Experimental.
	IdentityCenterInstanceArn *string `field:"optional" json:"identityCenterInstanceArn" yaml:"identityCenterInstanceArn"`
}

