package awsprovider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsprovider/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsprovider/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs aws}.
// Experimental.
type AwsProvider interface {
	cdktn.TerraformProvider
	// Experimental.
	AccessKey() *string
	// Experimental.
	SetAccessKey(val *string)
	// Experimental.
	AccessKeyInput() *string
	// Experimental.
	Alias() *string
	// Experimental.
	SetAlias(val *string)
	// Experimental.
	AliasInput() *string
	// Experimental.
	AllowedAccountIds() *[]*string
	// Experimental.
	SetAllowedAccountIds(val *[]*string)
	// Experimental.
	AllowedAccountIdsInput() *[]*string
	// Experimental.
	AssumeRole() interface{}
	// Experimental.
	SetAssumeRole(val interface{})
	// Experimental.
	AssumeRoleInput() interface{}
	// Experimental.
	AssumeRoleWithWebIdentity() interface{}
	// Experimental.
	SetAssumeRoleWithWebIdentity(val interface{})
	// Experimental.
	AssumeRoleWithWebIdentityInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	CustomCaBundle() *string
	// Experimental.
	SetCustomCaBundle(val *string)
	// Experimental.
	CustomCaBundleInput() *string
	// Experimental.
	DefaultTags() interface{}
	// Experimental.
	SetDefaultTags(val interface{})
	// Experimental.
	DefaultTagsInput() interface{}
	// Experimental.
	Ec2MetadataServiceEndpoint() *string
	// Experimental.
	SetEc2MetadataServiceEndpoint(val *string)
	// Experimental.
	Ec2MetadataServiceEndpointInput() *string
	// Experimental.
	Ec2MetadataServiceEndpointMode() *string
	// Experimental.
	SetEc2MetadataServiceEndpointMode(val *string)
	// Experimental.
	Ec2MetadataServiceEndpointModeInput() *string
	// Experimental.
	Endpoints() interface{}
	// Experimental.
	SetEndpoints(val interface{})
	// Experimental.
	EndpointsInput() interface{}
	// Experimental.
	ForbiddenAccountIds() *[]*string
	// Experimental.
	SetForbiddenAccountIds(val *[]*string)
	// Experimental.
	ForbiddenAccountIdsInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Provider-defined functions of the aws provider.
	// Experimental.
	Functions() AwsProviderFunctions
	// Experimental.
	HttpProxy() *string
	// Experimental.
	SetHttpProxy(val *string)
	// Experimental.
	HttpProxyInput() *string
	// Experimental.
	HttpsProxy() *string
	// Experimental.
	SetHttpsProxy(val *string)
	// Experimental.
	HttpsProxyInput() *string
	// Experimental.
	IgnoreTags() interface{}
	// Experimental.
	SetIgnoreTags(val interface{})
	// Experimental.
	IgnoreTagsInput() interface{}
	// Experimental.
	Insecure() interface{}
	// Experimental.
	SetInsecure(val interface{})
	// Experimental.
	InsecureInput() interface{}
	// Experimental.
	MaxRetries() *float64
	// Experimental.
	SetMaxRetries(val *float64)
	// Experimental.
	MaxRetriesInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	NoProxy() *string
	// Experimental.
	SetNoProxy(val *string)
	// Experimental.
	NoProxyInput() *string
	// Experimental.
	Profile() *string
	// Experimental.
	SetProfile(val *string)
	// Experimental.
	ProfileInput() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	RetryMode() *string
	// Experimental.
	SetRetryMode(val *string)
	// Experimental.
	RetryModeInput() *string
	// Experimental.
	S3UsEast1RegionalEndpoint() *string
	// Experimental.
	SetS3UsEast1RegionalEndpoint(val *string)
	// Experimental.
	S3UsEast1RegionalEndpointInput() *string
	// Experimental.
	S3UsePathStyle() interface{}
	// Experimental.
	SetS3UsePathStyle(val interface{})
	// Experimental.
	S3UsePathStyleInput() interface{}
	// Experimental.
	SecretKey() *string
	// Experimental.
	SetSecretKey(val *string)
	// Experimental.
	SecretKeyInput() *string
	// Experimental.
	SharedConfigFiles() *[]*string
	// Experimental.
	SetSharedConfigFiles(val *[]*string)
	// Experimental.
	SharedConfigFilesInput() *[]*string
	// Experimental.
	SharedCredentialsFiles() *[]*string
	// Experimental.
	SetSharedCredentialsFiles(val *[]*string)
	// Experimental.
	SharedCredentialsFilesInput() *[]*string
	// Experimental.
	SkipCredentialsValidation() interface{}
	// Experimental.
	SetSkipCredentialsValidation(val interface{})
	// Experimental.
	SkipCredentialsValidationInput() interface{}
	// Experimental.
	SkipMetadataApiCheck() *string
	// Experimental.
	SetSkipMetadataApiCheck(val *string)
	// Experimental.
	SkipMetadataApiCheckInput() *string
	// Experimental.
	SkipRegionValidation() interface{}
	// Experimental.
	SetSkipRegionValidation(val interface{})
	// Experimental.
	SkipRegionValidationInput() interface{}
	// Experimental.
	SkipRequestingAccountId() interface{}
	// Experimental.
	SetSkipRequestingAccountId(val interface{})
	// Experimental.
	SkipRequestingAccountIdInput() interface{}
	// Experimental.
	StsRegion() *string
	// Experimental.
	SetStsRegion(val *string)
	// Experimental.
	StsRegionInput() *string
	// Experimental.
	TagPolicyCompliance() *string
	// Experimental.
	SetTagPolicyCompliance(val *string)
	// Experimental.
	TagPolicyComplianceInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Token() *string
	// Experimental.
	SetToken(val *string)
	// Experimental.
	TokenBucketRateLimiterCapacity() *float64
	// Experimental.
	SetTokenBucketRateLimiterCapacity(val *float64)
	// Experimental.
	TokenBucketRateLimiterCapacityInput() *float64
	// Experimental.
	TokenInput() *string
	// Experimental.
	UseDualstackEndpoint() interface{}
	// Experimental.
	SetUseDualstackEndpoint(val interface{})
	// Experimental.
	UseDualstackEndpointInput() interface{}
	// Experimental.
	UseFipsEndpoint() interface{}
	// Experimental.
	SetUseFipsEndpoint(val interface{})
	// Experimental.
	UseFipsEndpointInput() interface{}
	// Experimental.
	UserAgent() *[]*string
	// Experimental.
	SetUserAgent(val *[]*string)
	// Experimental.
	UserAgentInput() *[]*string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	// Experimental.
	ResetAccessKey()
	// Experimental.
	ResetAlias()
	// Experimental.
	ResetAllowedAccountIds()
	// Experimental.
	ResetAssumeRole()
	// Experimental.
	ResetAssumeRoleWithWebIdentity()
	// Experimental.
	ResetCustomCaBundle()
	// Experimental.
	ResetDefaultTags()
	// Experimental.
	ResetEc2MetadataServiceEndpoint()
	// Experimental.
	ResetEc2MetadataServiceEndpointMode()
	// Experimental.
	ResetEndpoints()
	// Experimental.
	ResetForbiddenAccountIds()
	// Experimental.
	ResetHttpProxy()
	// Experimental.
	ResetHttpsProxy()
	// Experimental.
	ResetIgnoreTags()
	// Experimental.
	ResetInsecure()
	// Experimental.
	ResetMaxRetries()
	// Experimental.
	ResetNoProxy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetProfile()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRetryMode()
	// Experimental.
	ResetS3UsEast1RegionalEndpoint()
	// Experimental.
	ResetS3UsePathStyle()
	// Experimental.
	ResetSecretKey()
	// Experimental.
	ResetSharedConfigFiles()
	// Experimental.
	ResetSharedCredentialsFiles()
	// Experimental.
	ResetSkipCredentialsValidation()
	// Experimental.
	ResetSkipMetadataApiCheck()
	// Experimental.
	ResetSkipRegionValidation()
	// Experimental.
	ResetSkipRequestingAccountId()
	// Experimental.
	ResetStsRegion()
	// Experimental.
	ResetTagPolicyCompliance()
	// Experimental.
	ResetToken()
	// Experimental.
	ResetTokenBucketRateLimiterCapacity()
	// Experimental.
	ResetUseDualstackEndpoint()
	// Experimental.
	ResetUseFipsEndpoint()
	// Experimental.
	ResetUserAgent()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for AwsProvider
type jsiiProxy_AwsProvider struct {
	internal.Type__cdktnTerraformProvider
}

func (j *jsiiProxy_AwsProvider) AccessKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AccessKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AllowedAccountIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAccountIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AllowedAccountIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedAccountIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AssumeRole() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assumeRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AssumeRoleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assumeRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AssumeRoleWithWebIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assumeRoleWithWebIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) AssumeRoleWithWebIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assumeRoleWithWebIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) CustomCaBundle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCaBundle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) CustomCaBundleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customCaBundleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) DefaultTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) DefaultTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Ec2MetadataServiceEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2MetadataServiceEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Ec2MetadataServiceEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2MetadataServiceEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Ec2MetadataServiceEndpointMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2MetadataServiceEndpointMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Ec2MetadataServiceEndpointModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2MetadataServiceEndpointModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Endpoints() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) EndpointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) ForbiddenAccountIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"forbiddenAccountIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) ForbiddenAccountIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"forbiddenAccountIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Functions() AwsProviderFunctions {
	var returns AwsProviderFunctions
	_jsii_.Get(
		j,
		"functions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) HttpProxy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) HttpProxyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) HttpsProxy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) HttpsProxyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) IgnoreTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) IgnoreTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Insecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) InsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) NoProxy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noProxy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) NoProxyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noProxyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) RetryMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retryMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) RetryModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"retryModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) S3UsEast1RegionalEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3UsEast1RegionalEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) S3UsEast1RegionalEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3UsEast1RegionalEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) S3UsePathStyle() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3UsePathStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) S3UsePathStyleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3UsePathStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SharedConfigFiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedConfigFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SharedConfigFilesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedConfigFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SharedCredentialsFiles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedCredentialsFiles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SharedCredentialsFilesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sharedCredentialsFilesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipCredentialsValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipCredentialsValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipCredentialsValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipCredentialsValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipMetadataApiCheck() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipMetadataApiCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipMetadataApiCheckInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipMetadataApiCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipRegionValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipRegionValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipRegionValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipRegionValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipRequestingAccountId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipRequestingAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) SkipRequestingAccountIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipRequestingAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) StsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) StsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) TagPolicyCompliance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagPolicyCompliance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) TagPolicyComplianceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagPolicyComplianceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) TokenBucketRateLimiterCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenBucketRateLimiterCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) TokenBucketRateLimiterCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenBucketRateLimiterCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) UseDualstackEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDualstackEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) UseDualstackEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDualstackEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) UseFipsEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useFipsEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) UseFipsEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useFipsEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) UserAgent() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsProvider) UserAgentInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userAgentInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs aws} Resource.
// Experimental.
func NewAwsProvider(scope constructs.Construct, id *string, config *AwsProviderConfig) AwsProvider {
	_init_.Initialize()

	if err := validateNewAwsProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsProvider{}

	_jsii_.Create(
		"@cdktn/aws-provider.AwsProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs aws} Resource.
// Experimental.
func NewAwsProvider_Override(a AwsProvider, scope constructs.Construct, id *string, config *AwsProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-provider.AwsProvider",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsProvider)SetAccessKey(val *string) {
	_jsii_.Set(
		j,
		"accessKey",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetAllowedAccountIds(val *[]*string) {
	_jsii_.Set(
		j,
		"allowedAccountIds",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetAssumeRole(val interface{}) {
	if err := j.validateSetAssumeRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assumeRole",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetAssumeRoleWithWebIdentity(val interface{}) {
	if err := j.validateSetAssumeRoleWithWebIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assumeRoleWithWebIdentity",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetCustomCaBundle(val *string) {
	_jsii_.Set(
		j,
		"customCaBundle",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetDefaultTags(val interface{}) {
	if err := j.validateSetDefaultTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTags",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetEc2MetadataServiceEndpoint(val *string) {
	_jsii_.Set(
		j,
		"ec2MetadataServiceEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetEc2MetadataServiceEndpointMode(val *string) {
	_jsii_.Set(
		j,
		"ec2MetadataServiceEndpointMode",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetEndpoints(val interface{}) {
	if err := j.validateSetEndpointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoints",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetForbiddenAccountIds(val *[]*string) {
	_jsii_.Set(
		j,
		"forbiddenAccountIds",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetHttpProxy(val *string) {
	_jsii_.Set(
		j,
		"httpProxy",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetHttpsProxy(val *string) {
	_jsii_.Set(
		j,
		"httpsProxy",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetIgnoreTags(val interface{}) {
	if err := j.validateSetIgnoreTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ignoreTags",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetInsecure(val interface{}) {
	if err := j.validateSetInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecure",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetNoProxy(val *string) {
	_jsii_.Set(
		j,
		"noProxy",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetProfile(val *string) {
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetRegion(val *string) {
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetRetryMode(val *string) {
	_jsii_.Set(
		j,
		"retryMode",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetS3UsEast1RegionalEndpoint(val *string) {
	_jsii_.Set(
		j,
		"s3UsEast1RegionalEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetS3UsePathStyle(val interface{}) {
	if err := j.validateSetS3UsePathStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3UsePathStyle",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetSecretKey(val *string) {
	_jsii_.Set(
		j,
		"secretKey",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetSharedConfigFiles(val *[]*string) {
	_jsii_.Set(
		j,
		"sharedConfigFiles",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetSharedCredentialsFiles(val *[]*string) {
	_jsii_.Set(
		j,
		"sharedCredentialsFiles",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetSkipCredentialsValidation(val interface{}) {
	if err := j.validateSetSkipCredentialsValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipCredentialsValidation",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetSkipMetadataApiCheck(val *string) {
	_jsii_.Set(
		j,
		"skipMetadataApiCheck",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetSkipRegionValidation(val interface{}) {
	if err := j.validateSetSkipRegionValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipRegionValidation",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetSkipRequestingAccountId(val interface{}) {
	if err := j.validateSetSkipRequestingAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipRequestingAccountId",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetStsRegion(val *string) {
	_jsii_.Set(
		j,
		"stsRegion",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetTagPolicyCompliance(val *string) {
	_jsii_.Set(
		j,
		"tagPolicyCompliance",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetTokenBucketRateLimiterCapacity(val *float64) {
	_jsii_.Set(
		j,
		"tokenBucketRateLimiterCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetUseDualstackEndpoint(val interface{}) {
	if err := j.validateSetUseDualstackEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDualstackEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetUseFipsEndpoint(val interface{}) {
	if err := j.validateSetUseFipsEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useFipsEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsProvider)SetUserAgent(val *[]*string) {
	_jsii_.Set(
		j,
		"userAgent",
		val,
	)
}

// Generates CDKTN code for importing a AwsProvider resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-provider.AwsProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func AwsProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-provider.AwsProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-provider.AwsProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-provider.AwsProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-provider.AwsProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsProvider) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsProvider) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsProvider) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsProvider) ResetAccessKey() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		a,
		"resetAlias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetAllowedAccountIds() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedAccountIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetAssumeRole() {
	_jsii_.InvokeVoid(
		a,
		"resetAssumeRole",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetAssumeRoleWithWebIdentity() {
	_jsii_.InvokeVoid(
		a,
		"resetAssumeRoleWithWebIdentity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetCustomCaBundle() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomCaBundle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetDefaultTags() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetEc2MetadataServiceEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2MetadataServiceEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetEc2MetadataServiceEndpointMode() {
	_jsii_.InvokeVoid(
		a,
		"resetEc2MetadataServiceEndpointMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetEndpoints() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpoints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetForbiddenAccountIds() {
	_jsii_.InvokeVoid(
		a,
		"resetForbiddenAccountIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetHttpProxy() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpProxy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetHttpsProxy() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpsProxy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetIgnoreTags() {
	_jsii_.InvokeVoid(
		a,
		"resetIgnoreTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetInsecure() {
	_jsii_.InvokeVoid(
		a,
		"resetInsecure",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetNoProxy() {
	_jsii_.InvokeVoid(
		a,
		"resetNoProxy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetProfile() {
	_jsii_.InvokeVoid(
		a,
		"resetProfile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetRetryMode() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetS3UsEast1RegionalEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetS3UsEast1RegionalEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetS3UsePathStyle() {
	_jsii_.InvokeVoid(
		a,
		"resetS3UsePathStyle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSecretKey() {
	_jsii_.InvokeVoid(
		a,
		"resetSecretKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSharedConfigFiles() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedConfigFiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSharedCredentialsFiles() {
	_jsii_.InvokeVoid(
		a,
		"resetSharedCredentialsFiles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSkipCredentialsValidation() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipCredentialsValidation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSkipMetadataApiCheck() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipMetadataApiCheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSkipRegionValidation() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipRegionValidation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetSkipRequestingAccountId() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipRequestingAccountId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetStsRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetStsRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetTagPolicyCompliance() {
	_jsii_.InvokeVoid(
		a,
		"resetTagPolicyCompliance",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetToken() {
	_jsii_.InvokeVoid(
		a,
		"resetToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetTokenBucketRateLimiterCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenBucketRateLimiterCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetUseDualstackEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetUseDualstackEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetUseFipsEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetUseFipsEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) ResetUserAgent() {
	_jsii_.InvokeVoid(
		a,
		"resetUserAgent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsProvider) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

