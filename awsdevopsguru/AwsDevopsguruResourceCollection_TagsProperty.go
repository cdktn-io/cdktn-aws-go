package awsdevopsguru


// Experimental.
type AwsDevopsguruResourceCollection_TagsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devopsguru_resource_collection#app_boundary_key AwsDevopsguruResourceCollection#app_boundary_key}.
	// Experimental.
	AppBoundaryKey *string `field:"required" json:"appBoundaryKey" yaml:"appBoundaryKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devopsguru_resource_collection#tag_values AwsDevopsguruResourceCollection#tag_values}.
	// Experimental.
	TagValues *[]*string `field:"required" json:"tagValues" yaml:"tagValues"`
}

