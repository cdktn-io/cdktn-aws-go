package awsbatch


// Experimental.
type TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#allow_privilege_escalation TfJobDefinition#allow_privilege_escalation}.
	// Experimental.
	AllowPrivilegeEscalation interface{} `field:"optional" json:"allowPrivilegeEscalation" yaml:"allowPrivilegeEscalation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#privileged TfJobDefinition#privileged}.
	// Experimental.
	Privileged interface{} `field:"optional" json:"privileged" yaml:"privileged"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#read_only_root_file_system TfJobDefinition#read_only_root_file_system}.
	// Experimental.
	ReadOnlyRootFileSystem interface{} `field:"optional" json:"readOnlyRootFileSystem" yaml:"readOnlyRootFileSystem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#run_as_group TfJobDefinition#run_as_group}.
	// Experimental.
	RunAsGroup *float64 `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#run_as_non_root TfJobDefinition#run_as_non_root}.
	// Experimental.
	RunAsNonRoot interface{} `field:"optional" json:"runAsNonRoot" yaml:"runAsNonRoot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#run_as_user TfJobDefinition#run_as_user}.
	// Experimental.
	RunAsUser *float64 `field:"optional" json:"runAsUser" yaml:"runAsUser"`
}

