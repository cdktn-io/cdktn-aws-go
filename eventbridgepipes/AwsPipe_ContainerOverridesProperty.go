package eventbridgepipes


// Experimental.
type AwsPipe_ContainerOverridesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#command AwsPipe#command}.
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#environment AwsPipe#environment}
	// Experimental.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#instance_type AwsPipe#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// resource_requirement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#resource_requirement AwsPipe#resource_requirement}
	// Experimental.
	ResourceRequirement interface{} `field:"optional" json:"resourceRequirement" yaml:"resourceRequirement"`
}

