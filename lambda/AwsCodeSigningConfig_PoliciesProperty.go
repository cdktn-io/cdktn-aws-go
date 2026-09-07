package lambda


// Experimental.
type AwsCodeSigningConfig_PoliciesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_code_signing_config#untrusted_artifact_on_deployment AwsCodeSigningConfig#untrusted_artifact_on_deployment}.
	// Experimental.
	UntrustedArtifactOnDeployment *string `field:"required" json:"untrustedArtifactOnDeployment" yaml:"untrustedArtifactOnDeployment"`
}

