package awsroute53


// Experimental.
type TfRecordsExclusive_GeolocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#continent_code TfRecordsExclusive#continent_code}.
	// Experimental.
	ContinentCode *string `field:"optional" json:"continentCode" yaml:"continentCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#country_code TfRecordsExclusive#country_code}.
	// Experimental.
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#subdivision_code TfRecordsExclusive#subdivision_code}.
	// Experimental.
	SubdivisionCode *string `field:"optional" json:"subdivisionCode" yaml:"subdivisionCode"`
}

