package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_AccessLogProperty struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#file AwsAppmeshVirtualNode#file}
	// Experimental.
	File *AwsAppmeshVirtualNode_SpecLoggingAccessLogFileProperty `field:"optional" json:"file" yaml:"file"`
}

