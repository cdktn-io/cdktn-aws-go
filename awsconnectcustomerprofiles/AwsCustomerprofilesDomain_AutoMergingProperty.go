package awsconnectcustomerprofiles


// Experimental.
type AwsCustomerprofilesDomain_AutoMergingProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#enabled AwsCustomerprofilesDomain#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// conflict_resolution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#conflict_resolution AwsCustomerprofilesDomain#conflict_resolution}
	// Experimental.
	ConflictResolution *AwsCustomerprofilesDomain_MatchingAutoMergingConflictResolutionProperty `field:"optional" json:"conflictResolution" yaml:"conflictResolution"`
	// consolidation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#consolidation AwsCustomerprofilesDomain#consolidation}
	// Experimental.
	Consolidation *AwsCustomerprofilesDomain_ConsolidationProperty `field:"optional" json:"consolidation" yaml:"consolidation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/customerprofiles_domain#min_allowed_confidence_score_for_merging AwsCustomerprofilesDomain#min_allowed_confidence_score_for_merging}.
	// Experimental.
	MinAllowedConfidenceScoreForMerging *float64 `field:"optional" json:"minAllowedConfidenceScoreForMerging" yaml:"minAllowedConfidenceScoreForMerging"`
}

