package awseventbridgepipes


// Experimental.
type TfPipe_ContainerOverridesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#command TfPipe#command}.
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#environment TfPipe#environment}
	// Experimental.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#instance_type TfPipe#instance_type}.
	// Experimental.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// resource_requirement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#resource_requirement TfPipe#resource_requirement}
	// Experimental.
	ResourceRequirement interface{} `field:"optional" json:"resourceRequirement" yaml:"resourceRequirement"`
}

