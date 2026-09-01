package awsecs


// Experimental.
type AwsEcsCluster_SettingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#name AwsEcsCluster#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#value AwsEcsCluster#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

