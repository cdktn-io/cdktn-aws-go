package awsbedrock


// Experimental.
type TfEvaluationJob_EvaluationConfigProperty struct {
	// automated block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#automated TfEvaluationJob#automated}
	// Experimental.
	Automated interface{} `field:"optional" json:"automated" yaml:"automated"`
	// human block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#human TfEvaluationJob#human}
	// Experimental.
	Human interface{} `field:"optional" json:"human" yaml:"human"`
}

