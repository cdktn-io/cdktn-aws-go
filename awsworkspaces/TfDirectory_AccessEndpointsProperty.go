package awsworkspaces


// Experimental.
type TfDirectory_AccessEndpointsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#access_endpoint_type TfDirectory#access_endpoint_type}.
	// Experimental.
	AccessEndpointType *string `field:"required" json:"accessEndpointType" yaml:"accessEndpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#vpc_endpoint_id TfDirectory#vpc_endpoint_id}.
	// Experimental.
	VpcEndpointId *string `field:"required" json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

