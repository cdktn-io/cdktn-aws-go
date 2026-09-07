package batch


// Experimental.
type AwsJobDefinition_PodPropertiesProperty struct {
	// containers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#containers AwsJobDefinition#containers}
	// Experimental.
	Containers interface{} `field:"required" json:"containers" yaml:"containers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#dns_policy AwsJobDefinition#dns_policy}.
	// Experimental.
	DnsPolicy *string `field:"optional" json:"dnsPolicy" yaml:"dnsPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#host_network AwsJobDefinition#host_network}.
	// Experimental.
	HostNetwork interface{} `field:"optional" json:"hostNetwork" yaml:"hostNetwork"`
	// image_pull_secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#image_pull_secret AwsJobDefinition#image_pull_secret}
	// Experimental.
	ImagePullSecret interface{} `field:"optional" json:"imagePullSecret" yaml:"imagePullSecret"`
	// init_containers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#init_containers AwsJobDefinition#init_containers}
	// Experimental.
	InitContainers interface{} `field:"optional" json:"initContainers" yaml:"initContainers"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#metadata AwsJobDefinition#metadata}
	// Experimental.
	Metadata *AwsJobDefinition_MetadataProperty `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#service_account_name AwsJobDefinition#service_account_name}.
	// Experimental.
	ServiceAccountName *string `field:"optional" json:"serviceAccountName" yaml:"serviceAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#share_process_namespace AwsJobDefinition#share_process_namespace}.
	// Experimental.
	ShareProcessNamespace interface{} `field:"optional" json:"shareProcessNamespace" yaml:"shareProcessNamespace"`
	// volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#volumes AwsJobDefinition#volumes}
	// Experimental.
	Volumes interface{} `field:"optional" json:"volumes" yaml:"volumes"`
}

