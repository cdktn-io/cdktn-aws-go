package appmesh


// Experimental.
type AwsVirtualNode_LoggingProperty struct {
	// access_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#access_log AwsVirtualNode#access_log}
	// Experimental.
	AccessLog *AwsVirtualNode_AccessLogProperty `field:"optional" json:"accessLog" yaml:"accessLog"`
}

