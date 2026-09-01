package awseventbridgepipes


// Experimental.
type AwsPipesPipe_SourceParametersSelfManagedKafkaParametersCredentialsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#basic_auth AwsPipesPipe#basic_auth}.
	// Experimental.
	BasicAuth *string `field:"optional" json:"basicAuth" yaml:"basicAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#client_certificate_tls_auth AwsPipesPipe#client_certificate_tls_auth}.
	// Experimental.
	ClientCertificateTlsAuth *string `field:"optional" json:"clientCertificateTlsAuth" yaml:"clientCertificateTlsAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#sasl_scram_256_auth AwsPipesPipe#sasl_scram_256_auth}.
	// Experimental.
	SaslScram256Auth *string `field:"optional" json:"saslScram256Auth" yaml:"saslScram256Auth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#sasl_scram_512_auth AwsPipesPipe#sasl_scram_512_auth}.
	// Experimental.
	SaslScram512Auth *string `field:"optional" json:"saslScram512Auth" yaml:"saslScram512Auth"`
}

