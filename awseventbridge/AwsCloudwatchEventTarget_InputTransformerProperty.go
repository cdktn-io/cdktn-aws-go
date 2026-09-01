package awseventbridge


// Experimental.
type AwsCloudwatchEventTarget_InputTransformerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#input_template AwsCloudwatchEventTarget#input_template}.
	// Experimental.
	InputTemplate *string `field:"required" json:"inputTemplate" yaml:"inputTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_event_target#input_paths AwsCloudwatchEventTarget#input_paths}.
	// Experimental.
	InputPaths *map[string]*string `field:"optional" json:"inputPaths" yaml:"inputPaths"`
}

