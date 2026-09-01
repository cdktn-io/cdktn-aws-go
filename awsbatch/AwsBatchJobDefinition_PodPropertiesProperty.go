package awsbatch


// Experimental.
type AwsBatchJobDefinition_PodPropertiesProperty struct {
	// containers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#containers AwsBatchJobDefinition#containers}
	// Experimental.
	Containers interface{} `field:"required" json:"containers" yaml:"containers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#dns_policy AwsBatchJobDefinition#dns_policy}.
	// Experimental.
	DnsPolicy *string `field:"optional" json:"dnsPolicy" yaml:"dnsPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#host_network AwsBatchJobDefinition#host_network}.
	// Experimental.
	HostNetwork interface{} `field:"optional" json:"hostNetwork" yaml:"hostNetwork"`
	// image_pull_secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#image_pull_secret AwsBatchJobDefinition#image_pull_secret}
	// Experimental.
	ImagePullSecret interface{} `field:"optional" json:"imagePullSecret" yaml:"imagePullSecret"`
	// init_containers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#init_containers AwsBatchJobDefinition#init_containers}
	// Experimental.
	InitContainers interface{} `field:"optional" json:"initContainers" yaml:"initContainers"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#metadata AwsBatchJobDefinition#metadata}
	// Experimental.
	Metadata *AwsBatchJobDefinition_MetadataProperty `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#service_account_name AwsBatchJobDefinition#service_account_name}.
	// Experimental.
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#share_process_namespace AwsBatchJobDefinition#share_process_namespace}.
	// Experimental.
	ShareProcessNamespace interface{} `field:"optional" json:"shareProcessNamespace" yaml:"shareProcessNamespace"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#volumes AwsBatchJobDefinition#volumes}
	// Experimental.
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

