package awseventbridgepipes


// Experimental.
type AwsPipesPipe_SourceParametersManagedStreamingKafkaParametersCredentialsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#client_certificate_tls_auth AwsPipesPipe#client_certificate_tls_auth}.
	// Experimental.
	ClientCertificateTlsAuth *string `field:"optional" json:"clientCertificateTlsAuth" yaml:"clientCertificateTlsAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#sasl_scram_512_auth AwsPipesPipe#sasl_scram_512_auth}.
	// Experimental.
	SaslScram512Auth *string `field:"optional" json:"saslScram512Auth" yaml:"saslScram512Auth"`
}

