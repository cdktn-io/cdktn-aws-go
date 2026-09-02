package awsoutposts


// Experimental.
type TfCapacityTask_InstancePoolProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/outposts_capacity_task#count TfCapacityTask#count}.
	// Experimental.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/outposts_capacity_task#instance_type TfCapacityTask#instance_type}.
	// Experimental.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
}

