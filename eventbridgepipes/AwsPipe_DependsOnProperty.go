package eventbridgepipes


// Experimental.
type AwsPipe_DependsOnProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#job_id AwsPipe#job_id}.
	// Experimental.
	JobId *string `field:"optional" json:"jobId" yaml:"jobId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#type AwsPipe#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

