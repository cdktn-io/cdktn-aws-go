package auditmanager


// Experimental.
type AwsFramework_ControlSetsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/auditmanager_framework#name AwsFramework#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// controls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/auditmanager_framework#controls AwsFramework#controls}
	// Experimental.
	Controls interface{} `field:"optional" json:"controls" yaml:"controls"`
}

