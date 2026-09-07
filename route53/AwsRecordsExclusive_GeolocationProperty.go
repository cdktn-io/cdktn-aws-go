package route53


// Experimental.
type AwsRecordsExclusive_GeolocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#continent_code AwsRecordsExclusive#continent_code}.
	// Experimental.
	ContinentCode *string `field:"optional" json:"continentCode" yaml:"continentCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#country_code AwsRecordsExclusive#country_code}.
	// Experimental.
	CountryCode *string `field:"optional" json:"countryCode" yaml:"countryCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#subdivision_code AwsRecordsExclusive#subdivision_code}.
	// Experimental.
	SubdivisionCode *string `field:"optional" json:"subdivisionCode" yaml:"subdivisionCode"`
}

