package awsoracledatabaseaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network aws_odb_network}.
// Experimental.
type AwsOdbNetwork interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	AvailabilityZone() *string
	// Experimental.
	SetAvailabilityZone(val *string)
	// Experimental.
	AvailabilityZoneId() *string
	// Experimental.
	SetAvailabilityZoneId(val *string)
	// Experimental.
	AvailabilityZoneIdInput() *string
	// Experimental.
	AvailabilityZoneInput() *string
	// Experimental.
	BackupSubnetCidr() *string
	// Experimental.
	SetBackupSubnetCidr(val *string)
	// Experimental.
	BackupSubnetCidrInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClientSubnetCidr() *string
	// Experimental.
	SetClientSubnetCidr(val *string)
	// Experimental.
	ClientSubnetCidrInput() *string
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
	CreatedAt() *string
	// Experimental.
	CrossRegionS3RestoreSourcesAccess() *[]*string
	// Experimental.
	SetCrossRegionS3RestoreSourcesAccess(val *[]*string)
	// Experimental.
	CrossRegionS3RestoreSourcesAccessInput() *[]*string
	// Experimental.
	CustomDomainName() *string
	// Experimental.
	SetCustomDomainName(val *string)
	// Experimental.
	CustomDomainNameInput() *string
	// Experimental.
	DefaultDnsPrefix() *string
	// Experimental.
	SetDefaultDnsPrefix(val *string)
	// Experimental.
	DefaultDnsPrefixInput() *string
	// Experimental.
	DeleteAssociatedResources() interface{}
	// Experimental.
	SetDeleteAssociatedResources(val interface{})
	// Experimental.
	DeleteAssociatedResourcesInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DisplayName() *string
	// Experimental.
	SetDisplayName(val *string)
	// Experimental.
	DisplayNameInput() *string
	// Experimental.
	Ec2PlacementGroupIds() *[]*string
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
	KmsAccess() *string
	// Experimental.
	SetKmsAccess(val *string)
	// Experimental.
	KmsAccessInput() *string
	// Experimental.
	KmsPolicyDocument() *string
	// Experimental.
	SetKmsPolicyDocument(val *string)
	// Experimental.
	KmsPolicyDocumentInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	ManagedServices() AwsOdbNetwork_ManagedServicesPropertyList
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OciDnsForwardingConfigs() AwsOdbNetwork_OciDnsForwardingConfigsPropertyList
	// Experimental.
	OciNetworkAnchorId() *string
	// Experimental.
	OciNetworkAnchorUrl() *string
	// Experimental.
	OciResourceAnchorName() *string
	// Experimental.
	OciVcnId() *string
	// Experimental.
	OciVcnUrl() *string
	// Experimental.
	PeeredCidrs() *[]*string
	// Experimental.
	PercentProgress() *float64
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
	S3Access() *string
	// Experimental.
	SetS3Access(val *string)
	// Experimental.
	S3AccessInput() *string
	// Experimental.
	S3PolicyDocument() *string
	// Experimental.
	SetS3PolicyDocument(val *string)
	// Experimental.
	S3PolicyDocumentInput() *string
	// Experimental.
	Status() *string
	// Experimental.
	StatusReason() *string
	// Experimental.
	StsAccess() *string
	// Experimental.
	SetStsAccess(val *string)
	// Experimental.
	StsAccessInput() *string
	// Experimental.
	StsPolicyDocument() *string
	// Experimental.
	SetStsPolicyDocument(val *string)
	// Experimental.
	StsPolicyDocumentInput() *string
	// Experimental.
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsAll() cdktn.StringMap
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsOdbNetwork_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	ZeroEtlAccess() *string
	// Experimental.
	SetZeroEtlAccess(val *string)
	// Experimental.
	ZeroEtlAccessInput() *string
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
	PutTimeouts(value *AwsOdbNetwork_TimeoutsProperty)
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
	ResetAvailabilityZone()
	// Experimental.
	ResetCrossRegionS3RestoreSourcesAccess()
	// Experimental.
	ResetCustomDomainName()
	// Experimental.
	ResetDefaultDnsPrefix()
	// Experimental.
	ResetDeleteAssociatedResources()
	// Experimental.
	ResetKmsAccess()
	// Experimental.
	ResetKmsPolicyDocument()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetS3PolicyDocument()
	// Experimental.
	ResetStsAccess()
	// Experimental.
	ResetStsPolicyDocument()
	// Experimental.
	ResetTags()
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

// The jsii proxy struct for AwsOdbNetwork
type jsiiProxy_AwsOdbNetwork struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsOdbNetwork) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) AvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) AvailabilityZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) BackupSubnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupSubnetCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) BackupSubnetCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupSubnetCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) ClientSubnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSubnetCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) ClientSubnetCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSubnetCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) CrossRegionS3RestoreSourcesAccess() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossRegionS3RestoreSourcesAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) CrossRegionS3RestoreSourcesAccessInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossRegionS3RestoreSourcesAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) CustomDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) CustomDomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) DefaultDnsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDnsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) DefaultDnsPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDnsPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) DeleteAssociatedResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAssociatedResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) DeleteAssociatedResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAssociatedResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) Ec2PlacementGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ec2PlacementGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) KmsAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) KmsAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) KmsPolicyDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsPolicyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) KmsPolicyDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsPolicyDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) ManagedServices() AwsOdbNetwork_ManagedServicesPropertyList {
	var returns AwsOdbNetwork_ManagedServicesPropertyList
	_jsii_.Get(
		j,
		"managedServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) OciDnsForwardingConfigs() AwsOdbNetwork_OciDnsForwardingConfigsPropertyList {
	var returns AwsOdbNetwork_OciDnsForwardingConfigsPropertyList
	_jsii_.Get(
		j,
		"ociDnsForwardingConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) OciNetworkAnchorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociNetworkAnchorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) OciNetworkAnchorUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociNetworkAnchorUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) OciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) OciVcnId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociVcnId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) OciVcnUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociVcnUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) PeeredCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peeredCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) PercentProgress() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentProgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) S3Access() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) S3AccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3AccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) S3PolicyDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3PolicyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) S3PolicyDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3PolicyDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) StsAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) StsAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) StsPolicyDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsPolicyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) StsPolicyDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsPolicyDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) Timeouts() AwsOdbNetwork_TimeoutsPropertyOutputReference {
	var returns AwsOdbNetwork_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) ZeroEtlAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zeroEtlAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOdbNetwork) ZeroEtlAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zeroEtlAccessInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network aws_odb_network} Resource.
// Experimental.
func NewAwsOdbNetwork(scope constructs.Construct, id *string, config *AwsOdbNetworkConfig) AwsOdbNetwork {
	_init_.Initialize()

	if err := validateNewAwsOdbNetworkParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsOdbNetwork{}

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.AwsOdbNetwork",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network aws_odb_network} Resource.
// Experimental.
func NewAwsOdbNetwork_Override(a AwsOdbNetwork, scope constructs.Construct, id *string, config *AwsOdbNetworkConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.AwsOdbNetwork",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetAvailabilityZoneId(val *string) {
	if err := j.validateSetAvailabilityZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneId",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetBackupSubnetCidr(val *string) {
	if err := j.validateSetBackupSubnetCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupSubnetCidr",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetClientSubnetCidr(val *string) {
	if err := j.validateSetClientSubnetCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSubnetCidr",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetCrossRegionS3RestoreSourcesAccess(val *[]*string) {
	if err := j.validateSetCrossRegionS3RestoreSourcesAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRegionS3RestoreSourcesAccess",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetCustomDomainName(val *string) {
	if err := j.validateSetCustomDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDomainName",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetDefaultDnsPrefix(val *string) {
	if err := j.validateSetDefaultDnsPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultDnsPrefix",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetDeleteAssociatedResources(val interface{}) {
	if err := j.validateSetDeleteAssociatedResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAssociatedResources",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetKmsAccess(val *string) {
	if err := j.validateSetKmsAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsAccess",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetKmsPolicyDocument(val *string) {
	if err := j.validateSetKmsPolicyDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsPolicyDocument",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetS3Access(val *string) {
	if err := j.validateSetS3AccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Access",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetS3PolicyDocument(val *string) {
	if err := j.validateSetS3PolicyDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3PolicyDocument",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetStsAccess(val *string) {
	if err := j.validateSetStsAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stsAccess",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetStsPolicyDocument(val *string) {
	if err := j.validateSetStsPolicyDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stsPolicyDocument",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsOdbNetwork)SetZeroEtlAccess(val *string) {
	if err := j.validateSetZeroEtlAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zeroEtlAccess",
		val,
	)
}

// Generates CDKTN code for importing a AwsOdbNetwork resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsOdbNetwork_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsOdbNetwork_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.AwsOdbNetwork",
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
func AwsOdbNetwork_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsOdbNetwork_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.AwsOdbNetwork",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsOdbNetwork_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsOdbNetwork_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.AwsOdbNetwork",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsOdbNetwork_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsOdbNetwork_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.AwsOdbNetwork",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsOdbNetwork_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-oracle-database-aws.AwsOdbNetwork",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsOdbNetwork) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsOdbNetwork) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsOdbNetwork) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsOdbNetwork) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOdbNetwork) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsOdbNetwork) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsOdbNetwork) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsOdbNetwork) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsOdbNetwork) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsOdbNetwork) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsOdbNetwork) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsOdbNetwork) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOdbNetwork) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsOdbNetwork) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOdbNetwork) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsOdbNetwork) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsOdbNetwork) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsOdbNetwork) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsOdbNetwork) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsOdbNetwork) PutTimeouts(value *AwsOdbNetwork_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOdbNetwork) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsOdbNetwork) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOdbNetwork) ResetCrossRegionS3RestoreSourcesAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetCrossRegionS3RestoreSourcesAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOdbNetwork) ResetCustomDomainName() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomDomainName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOdbNetwork) ResetDefaultDnsPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultDnsPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOdbNetwork) ResetDeleteAssociatedResources() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteAssociatedResources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOdbNetwork) ResetKmsAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOdbNetwork) ResetKmsPolicyDocument() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsPolicyDocument",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOdbNetwork) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOdbNetwork) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOdbNetwork) ResetS3PolicyDocument() {
	_jsii_.InvokeVoid(
		a,
		"resetS3PolicyDocument",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOdbNetwork) ResetStsAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetStsAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOdbNetwork) ResetStsPolicyDocument() {
	_jsii_.InvokeVoid(
		a,
		"resetStsPolicyDocument",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOdbNetwork) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOdbNetwork) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOdbNetwork) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOdbNetwork) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOdbNetwork) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOdbNetwork) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOdbNetwork) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOdbNetwork) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOdbNetwork) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

