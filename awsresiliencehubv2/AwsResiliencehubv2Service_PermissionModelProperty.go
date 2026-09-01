package awsresiliencehubv2


// Experimental.
type AwsResiliencehubv2Service_PermissionModelProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_service#invoker_role_name AwsResiliencehubv2Service#invoker_role_name}.
	// Experimental.
	InvokerRoleName *string `field:"required" json:"invokerRoleName" yaml:"invokerRoleName"`
	// cross_account_role block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_service#cross_account_role AwsResiliencehubv2Service#cross_account_role}
	// Experimental.
	CrossAccountRole interface{} `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
}

