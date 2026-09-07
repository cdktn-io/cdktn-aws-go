package elementalmedialive


// Experimental.
type AwsChannel_MultiplexSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#multiplex_id AwsChannel#multiplex_id}.
	// Experimental.
	MultiplexId *string `field:"required" json:"multiplexId" yaml:"multiplexId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#program_name AwsChannel#program_name}.
	// Experimental.
	ProgramName *string `field:"required" json:"programName" yaml:"programName"`
}

