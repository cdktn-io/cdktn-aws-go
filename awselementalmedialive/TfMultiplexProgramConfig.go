package awselementalmedialive

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMultiplexProgramConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_multiplex_program#multiplex_id TfMultiplexProgram#multiplex_id}.
	// Experimental.
	MultiplexId *string `field:"required" json:"multiplexId" yaml:"multiplexId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_multiplex_program#program_name TfMultiplexProgram#program_name}.
	// Experimental.
	ProgramName *string `field:"required" json:"programName" yaml:"programName"`
	// multiplex_program_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_multiplex_program#multiplex_program_settings TfMultiplexProgram#multiplex_program_settings}
	// Experimental.
	MultiplexProgramSettings interface{} `field:"optional" json:"multiplexProgramSettings" yaml:"multiplexProgramSettings"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_multiplex_program#region TfMultiplexProgram#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_multiplex_program#timeouts TfMultiplexProgram#timeouts}
	// Experimental.
	Timeouts *TfMultiplexProgram_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

