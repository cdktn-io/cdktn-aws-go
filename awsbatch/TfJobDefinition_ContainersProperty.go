package awsbatch


// Experimental.
type TfJobDefinition_ContainersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#image TfJobDefinition#image}.
	// Experimental.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#args TfJobDefinition#args}.
	// Experimental.
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#command TfJobDefinition#command}.
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#env TfJobDefinition#env}
	// Experimental.
	Env interface{} `field:"optional" json:"env" yaml:"env"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#image_pull_policy TfJobDefinition#image_pull_policy}.
	// Experimental.
	ImagePullPolicy *string `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#name TfJobDefinition#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#resources TfJobDefinition#resources}
	// Experimental.
	Resources *TfJobDefinition_EksPropertiesPodPropertiesContainersResourcesProperty `field:"optional" json:"resources" yaml:"resources"`
	// security_context block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#security_context TfJobDefinition#security_context}
	// Experimental.
	SecurityContext *TfJobDefinition_EksPropertiesPodPropertiesContainersSecurityContextProperty `field:"optional" json:"securityContext" yaml:"securityContext"`
	// volume_mounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#volume_mounts TfJobDefinition#volume_mounts}
	// Experimental.
	VolumeMounts interface{} `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
}

