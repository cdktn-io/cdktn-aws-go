package awslightsail


// Experimental.
type AwsLightsailContainerService_CertificateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_container_service#certificate_name AwsLightsailContainerService#certificate_name}.
	// Experimental.
	CertificateName *string `field:"required" json:"certificateName" yaml:"certificateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_container_service#domain_names AwsLightsailContainerService#domain_names}.
	// Experimental.
	DomainNames *[]*string `field:"required" json:"domainNames" yaml:"domainNames"`
}

