package oracledatabaseaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/oracledatabaseaws/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/oracledatabaseaws/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network aws_odb_network}.
// Experimental.
type AwsNetwork interface {
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
	ManagedServices() AwsNetwork_ManagedServicesPropertyList
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OciDnsForwardingConfigs() AwsNetwork_OciDnsForwardingConfigsPropertyList
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
	Timeouts() AwsNetwork_TimeoutsPropertyOutputReference
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
	PutTimeouts(value *AwsNetwork_TimeoutsProperty)
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

// The jsii proxy struct for AwsNetwork
type jsiiProxy_AwsNetwork struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsNetwork) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) AvailabilityZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) AvailabilityZoneIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) BackupSubnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupSubnetCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) BackupSubnetCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupSubnetCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) ClientSubnetCidr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSubnetCidr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) ClientSubnetCidrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSubnetCidrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) CrossRegionS3RestoreSourcesAccess() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossRegionS3RestoreSourcesAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) CrossRegionS3RestoreSourcesAccessInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crossRegionS3RestoreSourcesAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) CustomDomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) CustomDomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDomainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) DefaultDnsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDnsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) DefaultDnsPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultDnsPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) DeleteAssociatedResources() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAssociatedResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) DeleteAssociatedResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAssociatedResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) Ec2PlacementGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ec2PlacementGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) KmsAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) KmsAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) KmsPolicyDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsPolicyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) KmsPolicyDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsPolicyDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) ManagedServices() AwsNetwork_ManagedServicesPropertyList {
	var returns AwsNetwork_ManagedServicesPropertyList
	_jsii_.Get(
		j,
		"managedServices",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) OciDnsForwardingConfigs() AwsNetwork_OciDnsForwardingConfigsPropertyList {
	var returns AwsNetwork_OciDnsForwardingConfigsPropertyList
	_jsii_.Get(
		j,
		"ociDnsForwardingConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) OciNetworkAnchorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociNetworkAnchorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) OciNetworkAnchorUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociNetworkAnchorUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) OciResourceAnchorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociResourceAnchorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) OciVcnId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociVcnId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) OciVcnUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ociVcnUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) PeeredCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"peeredCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) PercentProgress() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentProgress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) S3Access() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) S3AccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3AccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) S3PolicyDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3PolicyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) S3PolicyDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3PolicyDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) StatusReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) StsAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) StsAccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) StsPolicyDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsPolicyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) StsPolicyDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stsPolicyDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) Timeouts() AwsNetwork_TimeoutsPropertyOutputReference {
	var returns AwsNetwork_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) ZeroEtlAccess() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zeroEtlAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetwork) ZeroEtlAccessInput() *string {
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
func NewAwsNetwork(scope constructs.Construct, id *string, config *AwsNetworkConfig) AwsNetwork {
	_init_.Initialize()

	if err := validateNewAwsNetworkParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsNetwork{}

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.AwsNetwork",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/odb_network aws_odb_network} Resource.
// Experimental.
func NewAwsNetwork_Override(a AwsNetwork, scope constructs.Construct, id *string, config *AwsNetworkConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.AwsNetwork",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsNetwork)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetAvailabilityZoneId(val *string) {
	if err := j.validateSetAvailabilityZoneIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneId",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetBackupSubnetCidr(val *string) {
	if err := j.validateSetBackupSubnetCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupSubnetCidr",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetClientSubnetCidr(val *string) {
	if err := j.validateSetClientSubnetCidrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSubnetCidr",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetCrossRegionS3RestoreSourcesAccess(val *[]*string) {
	if err := j.validateSetCrossRegionS3RestoreSourcesAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRegionS3RestoreSourcesAccess",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetCustomDomainName(val *string) {
	if err := j.validateSetCustomDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customDomainName",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetDefaultDnsPrefix(val *string) {
	if err := j.validateSetDefaultDnsPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultDnsPrefix",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetDeleteAssociatedResources(val interface{}) {
	if err := j.validateSetDeleteAssociatedResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAssociatedResources",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetKmsAccess(val *string) {
	if err := j.validateSetKmsAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsAccess",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetKmsPolicyDocument(val *string) {
	if err := j.validateSetKmsPolicyDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsPolicyDocument",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetS3Access(val *string) {
	if err := j.validateSetS3AccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Access",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetS3PolicyDocument(val *string) {
	if err := j.validateSetS3PolicyDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3PolicyDocument",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetStsAccess(val *string) {
	if err := j.validateSetStsAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stsAccess",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetStsPolicyDocument(val *string) {
	if err := j.validateSetStsPolicyDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stsPolicyDocument",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsNetwork)SetZeroEtlAccess(val *string) {
	if err := j.validateSetZeroEtlAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zeroEtlAccess",
		val,
	)
}

// Generates CDKTN code for importing a AwsNetwork resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsNetwork_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsNetwork_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.AwsNetwork",
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
func AwsNetwork_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsNetwork_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.AwsNetwork",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsNetwork_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsNetwork_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.AwsNetwork",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsNetwork_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsNetwork_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-oracle-database-aws.AwsNetwork",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsNetwork_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-oracle-database-aws.AwsNetwork",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsNetwork) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsNetwork) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsNetwork) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsNetwork) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsNetwork) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsNetwork) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsNetwork) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsNetwork) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsNetwork) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsNetwork) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsNetwork) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsNetwork) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetwork) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsNetwork) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsNetwork) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsNetwork) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsNetwork) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsNetwork) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsNetwork) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsNetwork) PutTimeouts(value *AwsNetwork_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetwork) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsNetwork) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		a,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetwork) ResetCrossRegionS3RestoreSourcesAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetCrossRegionS3RestoreSourcesAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetwork) ResetCustomDomainName() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomDomainName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetwork) ResetDefaultDnsPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultDnsPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetwork) ResetDeleteAssociatedResources() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteAssociatedResources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetwork) ResetKmsAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetwork) ResetKmsPolicyDocument() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsPolicyDocument",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetwork) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetwork) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetwork) ResetS3PolicyDocument() {
	_jsii_.InvokeVoid(
		a,
		"resetS3PolicyDocument",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetwork) ResetStsAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetStsAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetwork) ResetStsPolicyDocument() {
	_jsii_.InvokeVoid(
		a,
		"resetStsPolicyDocument",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetwork) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetwork) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetwork) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetwork) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetwork) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetwork) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetwork) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetwork) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetwork) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

