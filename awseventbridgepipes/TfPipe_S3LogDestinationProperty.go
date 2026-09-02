package awseventbridgepipes


// Experimental.
type TfPipe_S3LogDestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#bucket_name TfPipe#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#bucket_owner TfPipe#bucket_owner}.
	// Experimental.
	BucketOwner *string `field:"required" json:"bucketOwner" yaml:"bucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#output_format TfPipe#output_format}.
	// Experimental.
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#prefix TfPipe#prefix}.
	// Experimental.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

