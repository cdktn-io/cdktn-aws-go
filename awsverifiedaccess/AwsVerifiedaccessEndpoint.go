package awsverifiedaccess

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsverifiedaccess/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsverifiedaccess/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint aws_verifiedaccess_endpoint}.
// Experimental.
type AwsVerifiedaccessEndpoint interface {
	cdktn.TerraformResource
	// Experimental.
	ApplicationDomain() *string
	// Experimental.
	SetApplicationDomain(val *string)
	// Experimental.
	ApplicationDomainInput() *string
	// Experimental.
	AttachmentType() *string
	// Experimental.
	SetAttachmentType(val *string)
	// Experimental.
	AttachmentTypeInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CidrOptions() AwsVerifiedaccessEndpoint_CidrOptionsPropertyOutputReference
	// Experimental.
	CidrOptionsInput() *AwsVerifiedaccessEndpoint_CidrOptionsProperty
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
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	DeviceValidationDomain() *string
	// Experimental.
	DomainCertificateArn() *string
	// Experimental.
	SetDomainCertificateArn(val *string)
	// Experimental.
	DomainCertificateArnInput() *string
	// Experimental.
	EndpointDomain() *string
	// Experimental.
	EndpointDomainPrefix() *string
	// Experimental.
	SetEndpointDomainPrefix(val *string)
	// Experimental.
	EndpointDomainPrefixInput() *string
	// Experimental.
	EndpointType() *string
	// Experimental.
	SetEndpointType(val *string)
	// Experimental.
	EndpointTypeInput() *string
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
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LoadBalancerOptions() AwsVerifiedaccessEndpoint_LoadBalancerOptionsPropertyOutputReference
	// Experimental.
	LoadBalancerOptionsInput() *AwsVerifiedaccessEndpoint_LoadBalancerOptionsProperty
	// Experimental.
	NetworkInterfaceOptions() AwsVerifiedaccessEndpoint_NetworkInterfaceOptionsPropertyOutputReference
	// Experimental.
	NetworkInterfaceOptionsInput() *AwsVerifiedaccessEndpoint_NetworkInterfaceOptionsProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PolicyDocument() *string
	// Experimental.
	SetPolicyDocument(val *string)
	// Experimental.
	PolicyDocumentInput() *string
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
	RdsOptions() AwsVerifiedaccessEndpoint_RdsOptionsPropertyOutputReference
	// Experimental.
	RdsOptionsInput() *AwsVerifiedaccessEndpoint_RdsOptionsProperty
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	SecurityGroupIds() *[]*string
	// Experimental.
	SetSecurityGroupIds(val *[]*string)
	// Experimental.
	SecurityGroupIdsInput() *[]*string
	// Experimental.
	SseSpecification() AwsVerifiedaccessEndpoint_SseSpecificationPropertyOutputReference
	// Experimental.
	SseSpecificationInput() *AwsVerifiedaccessEndpoint_SseSpecificationProperty
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
	Timeouts() AwsVerifiedaccessEndpoint_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	VerifiedAccessGroupId() *string
	// Experimental.
	SetVerifiedAccessGroupId(val *string)
	// Experimental.
	VerifiedAccessGroupIdInput() *string
	// Experimental.
	VerifiedAccessInstanceId() *string
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
	PutCidrOptions(value *AwsVerifiedaccessEndpoint_CidrOptionsProperty)
	// Experimental.
	PutLoadBalancerOptions(value *AwsVerifiedaccessEndpoint_LoadBalancerOptionsProperty)
	// Experimental.
	PutNetworkInterfaceOptions(value *AwsVerifiedaccessEndpoint_NetworkInterfaceOptionsProperty)
	// Experimental.
	PutRdsOptions(value *AwsVerifiedaccessEndpoint_RdsOptionsProperty)
	// Experimental.
	PutSseSpecification(value *AwsVerifiedaccessEndpoint_SseSpecificationProperty)
	// Experimental.
	PutTimeouts(value *AwsVerifiedaccessEndpoint_TimeoutsProperty)
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
	ResetApplicationDomain()
	// Experimental.
	ResetCidrOptions()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetDomainCertificateArn()
	// Experimental.
	ResetEndpointDomainPrefix()
	// Experimental.
	ResetId()
	// Experimental.
	ResetLoadBalancerOptions()
	// Experimental.
	ResetNetworkInterfaceOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPolicyDocument()
	// Experimental.
	ResetRdsOptions()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSecurityGroupIds()
	// Experimental.
	ResetSseSpecification()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
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

// The jsii proxy struct for AwsVerifiedaccessEndpoint
type jsiiProxy_AwsVerifiedaccessEndpoint struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) ApplicationDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) ApplicationDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) AttachmentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) AttachmentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attachmentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) CidrOptions() AwsVerifiedaccessEndpoint_CidrOptionsPropertyOutputReference {
	var returns AwsVerifiedaccessEndpoint_CidrOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"cidrOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) CidrOptionsInput() *AwsVerifiedaccessEndpoint_CidrOptionsProperty {
	var returns *AwsVerifiedaccessEndpoint_CidrOptionsProperty
	_jsii_.Get(
		j,
		"cidrOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) DeviceValidationDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceValidationDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) DomainCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) DomainCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) EndpointDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) EndpointDomainPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointDomainPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) EndpointDomainPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointDomainPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) LoadBalancerOptions() AwsVerifiedaccessEndpoint_LoadBalancerOptionsPropertyOutputReference {
	var returns AwsVerifiedaccessEndpoint_LoadBalancerOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"loadBalancerOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) LoadBalancerOptionsInput() *AwsVerifiedaccessEndpoint_LoadBalancerOptionsProperty {
	var returns *AwsVerifiedaccessEndpoint_LoadBalancerOptionsProperty
	_jsii_.Get(
		j,
		"loadBalancerOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) NetworkInterfaceOptions() AwsVerifiedaccessEndpoint_NetworkInterfaceOptionsPropertyOutputReference {
	var returns AwsVerifiedaccessEndpoint_NetworkInterfaceOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"networkInterfaceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) NetworkInterfaceOptionsInput() *AwsVerifiedaccessEndpoint_NetworkInterfaceOptionsProperty {
	var returns *AwsVerifiedaccessEndpoint_NetworkInterfaceOptionsProperty
	_jsii_.Get(
		j,
		"networkInterfaceOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) PolicyDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) PolicyDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) RdsOptions() AwsVerifiedaccessEndpoint_RdsOptionsPropertyOutputReference {
	var returns AwsVerifiedaccessEndpoint_RdsOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"rdsOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) RdsOptionsInput() *AwsVerifiedaccessEndpoint_RdsOptionsProperty {
	var returns *AwsVerifiedaccessEndpoint_RdsOptionsProperty
	_jsii_.Get(
		j,
		"rdsOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) SseSpecification() AwsVerifiedaccessEndpoint_SseSpecificationPropertyOutputReference {
	var returns AwsVerifiedaccessEndpoint_SseSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"sseSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) SseSpecificationInput() *AwsVerifiedaccessEndpoint_SseSpecificationProperty {
	var returns *AwsVerifiedaccessEndpoint_SseSpecificationProperty
	_jsii_.Get(
		j,
		"sseSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) Timeouts() AwsVerifiedaccessEndpoint_TimeoutsPropertyOutputReference {
	var returns AwsVerifiedaccessEndpoint_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) VerifiedAccessGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedAccessGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) VerifiedAccessGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedAccessGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint) VerifiedAccessInstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedAccessInstanceId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint aws_verifiedaccess_endpoint} Resource.
// Experimental.
func NewAwsVerifiedaccessEndpoint(scope constructs.Construct, id *string, config *AwsVerifiedaccessEndpointConfig) AwsVerifiedaccessEndpoint {
	_init_.Initialize()

	if err := validateNewAwsVerifiedaccessEndpointParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsVerifiedaccessEndpoint{}

	_jsii_.Create(
		"@cdktn/aws-verified-access.AwsVerifiedaccessEndpoint",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/verifiedaccess_endpoint aws_verifiedaccess_endpoint} Resource.
// Experimental.
func NewAwsVerifiedaccessEndpoint_Override(a AwsVerifiedaccessEndpoint, scope constructs.Construct, id *string, config *AwsVerifiedaccessEndpointConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-verified-access.AwsVerifiedaccessEndpoint",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetApplicationDomain(val *string) {
	if err := j.validateSetApplicationDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationDomain",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetAttachmentType(val *string) {
	if err := j.validateSetAttachmentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attachmentType",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetDomainCertificateArn(val *string) {
	if err := j.validateSetDomainCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainCertificateArn",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetEndpointDomainPrefix(val *string) {
	if err := j.validateSetEndpointDomainPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointDomainPrefix",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetEndpointType(val *string) {
	if err := j.validateSetEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetPolicyDocument(val *string) {
	if err := j.validateSetPolicyDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyDocument",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsVerifiedaccessEndpoint)SetVerifiedAccessGroupId(val *string) {
	if err := j.validateSetVerifiedAccessGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifiedAccessGroupId",
		val,
	)
}

// Generates CDKTN code for importing a AwsVerifiedaccessEndpoint resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsVerifiedaccessEndpoint_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsVerifiedaccessEndpoint_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-verified-access.AwsVerifiedaccessEndpoint",
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
func AwsVerifiedaccessEndpoint_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsVerifiedaccessEndpoint_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-verified-access.AwsVerifiedaccessEndpoint",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsVerifiedaccessEndpoint_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsVerifiedaccessEndpoint_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-verified-access.AwsVerifiedaccessEndpoint",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsVerifiedaccessEndpoint_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsVerifiedaccessEndpoint_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-verified-access.AwsVerifiedaccessEndpoint",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsVerifiedaccessEndpoint_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-verified-access.AwsVerifiedaccessEndpoint",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) PutCidrOptions(value *AwsVerifiedaccessEndpoint_CidrOptionsProperty) {
	if err := a.validatePutCidrOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCidrOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) PutLoadBalancerOptions(value *AwsVerifiedaccessEndpoint_LoadBalancerOptionsProperty) {
	if err := a.validatePutLoadBalancerOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLoadBalancerOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) PutNetworkInterfaceOptions(value *AwsVerifiedaccessEndpoint_NetworkInterfaceOptionsProperty) {
	if err := a.validatePutNetworkInterfaceOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkInterfaceOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) PutRdsOptions(value *AwsVerifiedaccessEndpoint_RdsOptionsProperty) {
	if err := a.validatePutRdsOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRdsOptions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) PutSseSpecification(value *AwsVerifiedaccessEndpoint_SseSpecificationProperty) {
	if err := a.validatePutSseSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSseSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) PutTimeouts(value *AwsVerifiedaccessEndpoint_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ResetApplicationDomain() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationDomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ResetCidrOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetCidrOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ResetDomainCertificateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetDomainCertificateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ResetEndpointDomainPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetEndpointDomainPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ResetLoadBalancerOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetLoadBalancerOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ResetNetworkInterfaceOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkInterfaceOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ResetPolicyDocument() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyDocument",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ResetRdsOptions() {
	_jsii_.InvokeVoid(
		a,
		"resetRdsOptions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ResetSseSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetSseSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVerifiedaccessEndpoint) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

