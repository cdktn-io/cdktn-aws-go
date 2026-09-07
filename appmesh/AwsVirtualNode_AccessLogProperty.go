package appmesh


// Experimental.
type AwsVirtualNode_AccessLogProperty struct {
	// file block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#file AwsVirtualNode#file}
	// Experimental.
	File *AwsVirtualNode_SpecLoggingAccessLogFileProperty `field:"optional" json:"file" yaml:"file"`
}

