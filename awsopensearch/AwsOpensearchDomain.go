// cdk-terrain bindings for the OpenSearch group of terraform-provider-aws 6.62.0
package awsopensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsopensearch/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsopensearch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain aws_opensearch_domain}.
// Experimental.
type AwsOpensearchDomain interface {
	cdktn.TerraformResource
	// Experimental.
	AccessPolicies() *string
	// Experimental.
	SetAccessPolicies(val *string)
	// Experimental.
	AccessPoliciesInput() *string
	// Experimental.
	AdvancedOptions() *map[string]*string
	// Experimental.
	SetAdvancedOptions(val *map[string]*string)
	// Experimental.
	AdvancedOptionsInput() *map[string]*string
	// Experimental.
	AdvancedSecurityOptions() AwsOpensearchDomain_AdvancedSecurityOptionsPropertyOutputReference
	// Experimental.
	AdvancedSecurityOptionsInput() *AwsOpensearchDomain_AdvancedSecurityOptionsProperty
	// Experimental.
	AimlOptions() AwsOpensearchDomain_AimlOptionsPropertyOutputReference
	// Experimental.
	AimlOptionsInput() *AwsOpensearchDomain_AimlOptionsProperty
	// Experimental.
	Arn() *string
	// Experimental.
	AutoTuneOptions() AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference
	// Experimental.
	AutoTuneOptionsInput() *AwsOpensearchDomain_AutoTuneOptionsProperty
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClusterConfig() AwsOpensearchDomain_ClusterConfigPropertyOutputReference
	// Experimental.
	ClusterConfigInput() *AwsOpensearchDomain_ClusterConfigProperty
	// Experimental.
	CognitoOptions() AwsOpensearchDomain_CognitoOptionsPropertyOutputReference
	// Experimental.
	CognitoOptionsInput() *AwsOpensearchDomain_CognitoOptionsProperty
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DashboardEndpoint() *string
	// Experimental.
	DashboardEndpointV2() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DeploymentStrategyOptions() AwsOpensearchDomain_DeploymentStrategyOptionsPropertyOutputReference
	// Experimental.
	DeploymentStrategyOptionsInput() *AwsOpensearchDomain_DeploymentStrategyOptionsProperty
	// Experimental.
	DomainEndpointOptions() AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference
	// Experimental.
	DomainEndpointOptionsInput() *AwsOpensearchDomain_DomainEndpointOptionsProperty
	// Experimental.
	DomainEndpointV2HostedZoneId() *string
	// Experimental.
	DomainId() *string
	// Experimental.
	DomainName() *string
	// Experimental.
	SetDomainName(val *string)
	// Experimental.
	DomainNameInput() *string
	// Experimental.
	EbsOptions() AwsOpensearchDomain_EbsOptionsPropertyOutputReference
	// Experimental.
	EbsOptionsInput() *AwsOpensearchDomain_EbsOptionsProperty
	// Experimental.
	EncryptAtRest() AwsOpensearchDomain_EncryptAtRestPropertyOutputReference
	// Experimental.
	EncryptAtRestInput() *AwsOpensearchDomain_EncryptAtRestProperty
	// Experimental.
	Endpoint() *string
	// Experimental.
	EndpointV2() *string
	// Experimental.
	EngineVersion() *string
	// Experimental.
	SetEngineVersion(val *string)
	// Experimental.
	EngineVersionInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdentityCenterOptions() AwsOpensearchDomain_IdentityCenterOptionsPropertyOutputReference
	// Experimental.
	IdentityCenterOptionsInput() *AwsOpensearchDomain_IdentityCenterOptionsProperty
	// Experimental.
	IdInput() *string
	// Experimental.
	IpAddressType() *string
	// Experimental.
	SetIpAddressType(val *string)
	// Experimental.
	IpAddressTypeInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LogPublishingOptions() AwsOpensearchDomain_LogPublishingOptionsPropertyList
	// Experimental.
	LogPublishingOptionsInput() interface{}
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	NodeToNodeEncryption() AwsOpensearchDomain_NodeToNodeEncryptionPropertyOutputReference
	// Experimental.
	NodeToNodeEncryptionInput() *AwsOpensearchDomain_NodeToNodeEncryptionProperty
	// Experimental.
	OffPeakWindowOptions() AwsOpensearchDomain_OffPeakWindowOptionsPropertyOutputReference
	// Experimental.
	OffPeakWindowOptionsInput() *AwsOpensearchDomain_OffPeakWindowOptionsProperty
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	SnapshotOptions() AwsOpensearchDomain_SnapshotOptionsPropertyOutputReference
	// Experimental.
	SnapshotOptionsInput() *AwsOpensearchDomain_SnapshotOptionsProperty
	// Experimental.
	SoftwareUpdateOptions() AwsOpensearchDomain_SoftwareUpdateOptionsPropertyOutputReference
	// Experimental.
	SoftwareUpdateOptionsInput() *AwsOpensearchDomain_SoftwareUpdateOptionsProperty
	// Experimental.
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsAll() *map[string]*string
	// Experimental.
	SetTagsAll(val *map[string]*string)
	// Experimental.
	TagsAllInput() *map[string]*string
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsOpensearchDomain_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	VpcOptions() AwsOpensearchDomain_VpcOptionsPropertyOutputReference
	// Experimental.
	VpcOptionsInput() *AwsOpensearchDomain_VpcOptionsProperty
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Experimental.
	PutAdvancedSecurityOptions(value *AwsOpensearchDomain_AdvancedSecurityOptionsProperty)
	// Experimental.
	PutAimlOptions(value *AwsOpensearchDomain_AimlOptionsProperty)
	// Experimental.
	PutAutoTuneOptions(value *AwsOpensearchDomain_AutoTuneOptionsProperty)
	// Experimental.
	PutClusterConfig(value *AwsOpensearchDomain_ClusterConfigProperty)
	// Experimental.
	PutCognitoOptions(value *AwsOpensearchDomain_CognitoOptionsProperty)
	// Experimental.
	PutDeploymentStrategyOptions(value *AwsOpensearchDomain_DeploymentStrategyOptionsProperty)
	// Experimental.
	PutDomainEndpointOptions(value *AwsOpensearchDomain_DomainEndpointOptionsProperty)
	// Experimental.
	PutEbsOptions(value *AwsOpensearchDomain_EbsOptionsProperty)
	// Experimental.
	PutEncryptAtRest(value *AwsOpensearchDomain_EncryptAtRestProperty)
	// Experimental.
	PutIdentityCenterOptions(value *AwsOpensearchDomain_IdentityCenterOptionsProperty)
	// Experimental.
	PutLogPublishingOptions(value interface{})
	// Experimental.
	PutNodeToNodeEncryption(value *AwsOpensearchDomain_NodeToNodeEncryptionProperty)
	// Experimental.
	PutOffPeakWindowOptions(value *AwsOpensearchDomain_OffPeakWindowOptionsProperty)
	// Experimental.
	PutSnapshotOptions(value *AwsOpensearchDomain_SnapshotOptionsProperty)
	// Experimental.
	PutSoftwareUpdateOptions(value *AwsOpensearchDomain_SoftwareUpdateOptionsProperty)
	// Experimental.
	PutTimeouts(value *AwsOpensearchDomain_TimeoutsProperty)
	// Experimental.
	PutVpcOptions(value *AwsOpensearchDomain_VpcOptionsProperty)
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
	ResetAccessPolicies()
	// Experimental.
	ResetAdvancedOptions()
	// Experimental.
	ResetAdvancedSecurityOptions()
	// Experimental.
	ResetAimlOptions()
	// Experimental.
	ResetAutoTuneOptions()
	// Experimental.
	ResetClusterConfig()
	// Experimental.
	ResetCognitoOptions()
	// Experimental.
	ResetDeploymentStrategyOptions()
	// Experimental.
	ResetDomainEndpointOptions()
	// Experimental.
	ResetEbsOptions()
	// Experimental.
	ResetEncryptAtRest()
	// Experimental.
	ResetEngineVersion()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIdentityCenterOptions()
	// Experimental.
	ResetIpAddressType()
	// Experimental.
	ResetLogPublishingOptions()
	// Experimental.
	ResetNodeToNodeEncryption()
	// Experimental.
	ResetOffPeakWindowOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSnapshotOptions()
	// Experimental.
	ResetSoftwareUpdateOptions()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetVpcOptions()
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

// The jsii proxy struct for AwsOpensearchDomain
type jsiiProxy_AwsOpensearchDomain struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsOpensearchDomain) AccessPolicies() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) AccessPoliciesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) AdvancedOptions() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"advancedOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) AdvancedOptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"advancedOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) AdvancedSecurityOptions() AwsOpensearchDomain_AdvancedSecurityOptionsPropertyOutputReference {
	var returns AwsOpensearchDomain_AdvancedSecurityOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedSecurityOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) AdvancedSecurityOptionsInput() *AwsOpensearchDomain_AdvancedSecurityOptionsProperty {
	var returns *AwsOpensearchDomain_AdvancedSecurityOptionsProperty
	_jsii_.Get(
		j,
		"advancedSecurityOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) AimlOptions() AwsOpensearchDomain_AimlOptionsPropertyOutputReference {
	var returns AwsOpensearchDomain_AimlOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"aimlOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) AimlOptionsInput() *AwsOpensearchDomain_AimlOptionsProperty {
	var returns *AwsOpensearchDomain_AimlOptionsProperty
	_jsii_.Get(
		j,
		"aimlOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) AutoTuneOptions() AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference {
	var returns AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"autoTuneOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) AutoTuneOptionsInput() *AwsOpensearchDomain_AutoTuneOptionsProperty {
	var returns *AwsOpensearchDomain_AutoTuneOptionsProperty
	_jsii_.Get(
		j,
		"autoTuneOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) ClusterConfig() AwsOpensearchDomain_ClusterConfigPropertyOutputReference {
	var returns AwsOpensearchDomain_ClusterConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"clusterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) ClusterConfigInput() *AwsOpensearchDomain_ClusterConfigProperty {
	var returns *AwsOpensearchDomain_ClusterConfigProperty
	_jsii_.Get(
		j,
		"clusterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) CognitoOptions() AwsOpensearchDomain_CognitoOptionsPropertyOutputReference {
	var returns AwsOpensearchDomain_CognitoOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"cognitoOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) CognitoOptionsInput() *AwsOpensearchDomain_CognitoOptionsProperty {
	var returns *AwsOpensearchDomain_CognitoOptionsProperty
	_jsii_.Get(
		j,
		"cognitoOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) DashboardEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) DashboardEndpointV2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dashboardEndpointV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) DeploymentStrategyOptions() AwsOpensearchDomain_DeploymentStrategyOptionsPropertyOutputReference {
	var returns AwsOpensearchDomain_DeploymentStrategyOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"deploymentStrategyOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) DeploymentStrategyOptionsInput() *AwsOpensearchDomain_DeploymentStrategyOptionsProperty {
	var returns *AwsOpensearchDomain_DeploymentStrategyOptionsProperty
	_jsii_.Get(
		j,
		"deploymentStrategyOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) DomainEndpointOptions() AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference {
	var returns AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"domainEndpointOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) DomainEndpointOptionsInput() *AwsOpensearchDomain_DomainEndpointOptionsProperty {
	var returns *AwsOpensearchDomain_DomainEndpointOptionsProperty
	_jsii_.Get(
		j,
		"domainEndpointOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) DomainEndpointV2HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainEndpointV2HostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) DomainId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) EbsOptions() AwsOpensearchDomain_EbsOptionsPropertyOutputReference {
	var returns AwsOpensearchDomain_EbsOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"ebsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) EbsOptionsInput() *AwsOpensearchDomain_EbsOptionsProperty {
	var returns *AwsOpensearchDomain_EbsOptionsProperty
	_jsii_.Get(
		j,
		"ebsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) EncryptAtRest() AwsOpensearchDomain_EncryptAtRestPropertyOutputReference {
	var returns AwsOpensearchDomain_EncryptAtRestPropertyOutputReference
	_jsii_.Get(
		j,
		"encryptAtRest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) EncryptAtRestInput() *AwsOpensearchDomain_EncryptAtRestProperty {
	var returns *AwsOpensearchDomain_EncryptAtRestProperty
	_jsii_.Get(
		j,
		"encryptAtRestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) EndpointV2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) IdentityCenterOptions() AwsOpensearchDomain_IdentityCenterOptionsPropertyOutputReference {
	var returns AwsOpensearchDomain_IdentityCenterOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"identityCenterOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) IdentityCenterOptionsInput() *AwsOpensearchDomain_IdentityCenterOptionsProperty {
	var returns *AwsOpensearchDomain_IdentityCenterOptionsProperty
	_jsii_.Get(
		j,
		"identityCenterOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) LogPublishingOptions() AwsOpensearchDomain_LogPublishingOptionsPropertyList {
	var returns AwsOpensearchDomain_LogPublishingOptionsPropertyList
	_jsii_.Get(
		j,
		"logPublishingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) LogPublishingOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logPublishingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) NodeToNodeEncryption() AwsOpensearchDomain_NodeToNodeEncryptionPropertyOutputReference {
	var returns AwsOpensearchDomain_NodeToNodeEncryptionPropertyOutputReference
	_jsii_.Get(
		j,
		"nodeToNodeEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) NodeToNodeEncryptionInput() *AwsOpensearchDomain_NodeToNodeEncryptionProperty {
	var returns *AwsOpensearchDomain_NodeToNodeEncryptionProperty
	_jsii_.Get(
		j,
		"nodeToNodeEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) OffPeakWindowOptions() AwsOpensearchDomain_OffPeakWindowOptionsPropertyOutputReference {
	var returns AwsOpensearchDomain_OffPeakWindowOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"offPeakWindowOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) OffPeakWindowOptionsInput() *AwsOpensearchDomain_OffPeakWindowOptionsProperty {
	var returns *AwsOpensearchDomain_OffPeakWindowOptionsProperty
	_jsii_.Get(
		j,
		"offPeakWindowOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) SnapshotOptions() AwsOpensearchDomain_SnapshotOptionsPropertyOutputReference {
	var returns AwsOpensearchDomain_SnapshotOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"snapshotOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) SnapshotOptionsInput() *AwsOpensearchDomain_SnapshotOptionsProperty {
	var returns *AwsOpensearchDomain_SnapshotOptionsProperty
	_jsii_.Get(
		j,
		"snapshotOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) SoftwareUpdateOptions() AwsOpensearchDomain_SoftwareUpdateOptionsPropertyOutputReference {
	var returns AwsOpensearchDomain_SoftwareUpdateOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"softwareUpdateOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) SoftwareUpdateOptionsInput() *AwsOpensearchDomain_SoftwareUpdateOptionsProperty {
	var returns *AwsOpensearchDomain_SoftwareUpdateOptionsProperty
	_jsii_.Get(
		j,
		"softwareUpdateOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) Timeouts() AwsOpensearchDomain_TimeoutsPropertyOutputReference {
	var returns AwsOpensearchDomain_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) VpcOptions() AwsOpensearchDomain_VpcOptionsPropertyOutputReference {
	var returns AwsOpensearchDomain_VpcOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain) VpcOptionsInput() *AwsOpensearchDomain_VpcOptionsProperty {
	var returns *AwsOpensearchDomain_VpcOptionsProperty
	_jsii_.Get(
		j,
		"vpcOptionsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain aws_opensearch_domain} Resource.
// Experimental.
func NewAwsOpensearchDomain(scope constructs.Construct, id *string, config *AwsOpensearchDomainConfig) AwsOpensearchDomain {
	_init_.Initialize()

	if err := validateNewAwsOpensearchDomainParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsOpensearchDomain{}

	_jsii_.Create(
		"@cdktn/aws-opensearch.AwsOpensearchDomain",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain aws_opensearch_domain} Resource.
// Experimental.
func NewAwsOpensearchDomain_Override(a AwsOpensearchDomain, scope constructs.Construct, id *string, config *AwsOpensearchDomainConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-opensearch.AwsOpensearchDomain",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain)SetAccessPolicies(val *string) {
	if err := j.validateSetAccessPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessPolicies",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain)SetAdvancedOptions(val *map[string]*string) {
	if err := j.validateSetAdvancedOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advancedOptions",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a AwsOpensearchDomain resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsOpensearchDomain_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsOpensearchDomain_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-opensearch.AwsOpensearchDomain",
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
func AwsOpensearchDomain_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsOpensearchDomain_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-opensearch.AwsOpensearchDomain",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsOpensearchDomain_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsOpensearchDomain_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-opensearch.AwsOpensearchDomain",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsOpensearchDomain_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsOpensearchDomain_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-opensearch.AwsOpensearchDomain",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsOpensearchDomain_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-opensearch.AwsOpensearchDomain",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := a.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) PutAdvancedSecurityOptions(value *AwsOpensearchDomain_AdvancedSecurityOptionsProperty) {
	if err := a.validatePutAdvancedSecurityOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdvancedSecurityOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) PutAimlOptions(value *AwsOpensearchDomain_AimlOptionsProperty) {
	if err := a.validatePutAimlOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAimlOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) PutAutoTuneOptions(value *AwsOpensearchDomain_AutoTuneOptionsProperty) {
	if err := a.validatePutAutoTuneOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutoTuneOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) PutClusterConfig(value *AwsOpensearchDomain_ClusterConfigProperty) {
	if err := a.validatePutClusterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putClusterConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) PutCognitoOptions(value *AwsOpensearchDomain_CognitoOptionsProperty) {
	if err := a.validatePutCognitoOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCognitoOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) PutDeploymentStrategyOptions(value *AwsOpensearchDomain_DeploymentStrategyOptionsProperty) {
	if err := a.validatePutDeploymentStrategyOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeploymentStrategyOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) PutDomainEndpointOptions(value *AwsOpensearchDomain_DomainEndpointOptionsProperty) {
	if err := a.validatePutDomainEndpointOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDomainEndpointOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) PutEbsOptions(value *AwsOpensearchDomain_EbsOptionsProperty) {
	if err := a.validatePutEbsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEbsOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) PutEncryptAtRest(value *AwsOpensearchDomain_EncryptAtRestProperty) {
	if err := a.validatePutEncryptAtRestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEncryptAtRest",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) PutIdentityCenterOptions(value *AwsOpensearchDomain_IdentityCenterOptionsProperty) {
	if err := a.validatePutIdentityCenterOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIdentityCenterOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) PutLogPublishingOptions(value interface{}) {
	if err := a.validatePutLogPublishingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogPublishingOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) PutNodeToNodeEncryption(value *AwsOpensearchDomain_NodeToNodeEncryptionProperty) {
	if err := a.validatePutNodeToNodeEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNodeToNodeEncryption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) PutOffPeakWindowOptions(value *AwsOpensearchDomain_OffPeakWindowOptionsProperty) {
	if err := a.validatePutOffPeakWindowOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOffPeakWindowOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) PutSnapshotOptions(value *AwsOpensearchDomain_SnapshotOptionsProperty) {
	if err := a.validatePutSnapshotOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnapshotOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) PutSoftwareUpdateOptions(value *AwsOpensearchDomain_SoftwareUpdateOptionsProperty) {
	if err := a.validatePutSoftwareUpdateOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSoftwareUpdateOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) PutTimeouts(value *AwsOpensearchDomain_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) PutVpcOptions(value *AwsOpensearchDomain_VpcOptionsProperty) {
	if err := a.validatePutVpcOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpcOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetAccessPolicies() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessPolicies",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetAdvancedOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetAdvancedSecurityOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedSecurityOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetAimlOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetAimlOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetAutoTuneOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoTuneOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetClusterConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetClusterConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetCognitoOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetCognitoOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetDeploymentStrategyOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetDeploymentStrategyOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetDomainEndpointOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainEndpointOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetEbsOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetEbsOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetEncryptAtRest() {
	_jsii_.InvokeVoid(
		a,
		"resetEncryptAtRest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetIdentityCenterOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetIdentityCenterOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		a,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetLogPublishingOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetLogPublishingOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetNodeToNodeEncryption() {
	_jsii_.InvokeVoid(
		a,
		"resetNodeToNodeEncryption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetOffPeakWindowOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetOffPeakWindowOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetSnapshotOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapshotOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetSoftwareUpdateOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetSoftwareUpdateOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) ResetVpcOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

