package awsssoadmin


// Experimental.
type AwsSsoadminInstanceAccessControlAttributes_AttributeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_instance_access_control_attributes#key AwsSsoadminInstanceAccessControlAttributes#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssoadmin_instance_access_control_attributes#value AwsSsoadminInstanceAccessControlAttributes#value}
	// Experimental.
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

