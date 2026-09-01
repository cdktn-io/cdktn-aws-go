package awsstoragegateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsstoragegateway/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsstoragegateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share aws_storagegateway_nfs_file_share}.
// Experimental.
type AwsStoragegatewayNfsFileShare interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	AuditDestinationArn() *string
	// Experimental.
	SetAuditDestinationArn(val *string)
	// Experimental.
	AuditDestinationArnInput() *string
	// Experimental.
	BucketRegion() *string
	// Experimental.
	SetBucketRegion(val *string)
	// Experimental.
	BucketRegionInput() *string
	// Experimental.
	CacheAttributes() AwsStoragegatewayNfsFileShare_CacheAttributesPropertyOutputReference
	// Experimental.
	CacheAttributesInput() *AwsStoragegatewayNfsFileShare_CacheAttributesProperty
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClientList() *[]*string
	// Experimental.
	SetClientList(val *[]*string)
	// Experimental.
	ClientListInput() *[]*string
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
	DefaultStorageClass() *string
	// Experimental.
	SetDefaultStorageClass(val *string)
	// Experimental.
	DefaultStorageClassInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	FileshareId() *string
	// Experimental.
	FileShareName() *string
	// Experimental.
	SetFileShareName(val *string)
	// Experimental.
	FileShareNameInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	GatewayArn() *string
	// Experimental.
	SetGatewayArn(val *string)
	// Experimental.
	GatewayArnInput() *string
	// Experimental.
	GuessMimeTypeEnabled() interface{}
	// Experimental.
	SetGuessMimeTypeEnabled(val interface{})
	// Experimental.
	GuessMimeTypeEnabledInput() interface{}
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	KmsEncrypted() interface{}
	// Experimental.
	SetKmsEncrypted(val interface{})
	// Experimental.
	KmsEncryptedInput() interface{}
	// Experimental.
	KmsKeyArn() *string
	// Experimental.
	SetKmsKeyArn(val *string)
	// Experimental.
	KmsKeyArnInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LocationArn() *string
	// Experimental.
	SetLocationArn(val *string)
	// Experimental.
	LocationArnInput() *string
	// Experimental.
	NfsFileShareDefaults() AwsStoragegatewayNfsFileShare_NfsFileShareDefaultsPropertyOutputReference
	// Experimental.
	NfsFileShareDefaultsInput() *AwsStoragegatewayNfsFileShare_NfsFileShareDefaultsProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	NotificationPolicy() *string
	// Experimental.
	SetNotificationPolicy(val *string)
	// Experimental.
	NotificationPolicyInput() *string
	// Experimental.
	ObjectAcl() *string
	// Experimental.
	SetObjectAcl(val *string)
	// Experimental.
	ObjectAclInput() *string
	// Experimental.
	Path() *string
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
	ReadOnly() interface{}
	// Experimental.
	SetReadOnly(val interface{})
	// Experimental.
	ReadOnlyInput() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	RequesterPays() interface{}
	// Experimental.
	SetRequesterPays(val interface{})
	// Experimental.
	RequesterPaysInput() interface{}
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	Squash() *string
	// Experimental.
	SetSquash(val *string)
	// Experimental.
	SquashInput() *string
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
	Timeouts() AwsStoragegatewayNfsFileShare_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	VpcEndpointDnsName() *string
	// Experimental.
	SetVpcEndpointDnsName(val *string)
	// Experimental.
	VpcEndpointDnsNameInput() *string
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
	PutCacheAttributes(value *AwsStoragegatewayNfsFileShare_CacheAttributesProperty)
	// Experimental.
	PutNfsFileShareDefaults(value *AwsStoragegatewayNfsFileShare_NfsFileShareDefaultsProperty)
	// Experimental.
	PutTimeouts(value *AwsStoragegatewayNfsFileShare_TimeoutsProperty)
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
	ResetAuditDestinationArn()
	// Experimental.
	ResetBucketRegion()
	// Experimental.
	ResetCacheAttributes()
	// Experimental.
	ResetDefaultStorageClass()
	// Experimental.
	ResetFileShareName()
	// Experimental.
	ResetGuessMimeTypeEnabled()
	// Experimental.
	ResetId()
	// Experimental.
	ResetKmsEncrypted()
	// Experimental.
	ResetKmsKeyArn()
	// Experimental.
	ResetNfsFileShareDefaults()
	// Experimental.
	ResetNotificationPolicy()
	// Experimental.
	ResetObjectAcl()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetReadOnly()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRequesterPays()
	// Experimental.
	ResetSquash()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetVpcEndpointDnsName()
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

// The jsii proxy struct for AwsStoragegatewayNfsFileShare
type jsiiProxy_AwsStoragegatewayNfsFileShare struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) AuditDestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDestinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) AuditDestinationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDestinationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) BucketRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) BucketRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) CacheAttributes() AwsStoragegatewayNfsFileShare_CacheAttributesPropertyOutputReference {
	var returns AwsStoragegatewayNfsFileShare_CacheAttributesPropertyOutputReference
	_jsii_.Get(
		j,
		"cacheAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) CacheAttributesInput() *AwsStoragegatewayNfsFileShare_CacheAttributesProperty {
	var returns *AwsStoragegatewayNfsFileShare_CacheAttributesProperty
	_jsii_.Get(
		j,
		"cacheAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) ClientList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) ClientListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) DefaultStorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStorageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) DefaultStorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStorageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) FileshareId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileshareId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) FileShareName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) FileShareNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) GatewayArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) GuessMimeTypeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guessMimeTypeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) GuessMimeTypeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guessMimeTypeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) KmsEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) KmsEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) LocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) LocationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) NfsFileShareDefaults() AwsStoragegatewayNfsFileShare_NfsFileShareDefaultsPropertyOutputReference {
	var returns AwsStoragegatewayNfsFileShare_NfsFileShareDefaultsPropertyOutputReference
	_jsii_.Get(
		j,
		"nfsFileShareDefaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) NfsFileShareDefaultsInput() *AwsStoragegatewayNfsFileShare_NfsFileShareDefaultsProperty {
	var returns *AwsStoragegatewayNfsFileShare_NfsFileShareDefaultsProperty
	_jsii_.Get(
		j,
		"nfsFileShareDefaultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) NotificationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) NotificationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) ObjectAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) ObjectAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) RequesterPays() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) RequesterPaysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) Squash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) SquashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"squashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) Timeouts() AwsStoragegatewayNfsFileShare_TimeoutsPropertyOutputReference {
	var returns AwsStoragegatewayNfsFileShare_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) VpcEndpointDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare) VpcEndpointDnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointDnsNameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share aws_storagegateway_nfs_file_share} Resource.
// Experimental.
func NewAwsStoragegatewayNfsFileShare(scope constructs.Construct, id *string, config *AwsStoragegatewayNfsFileShareConfig) AwsStoragegatewayNfsFileShare {
	_init_.Initialize()

	if err := validateNewAwsStoragegatewayNfsFileShareParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsStoragegatewayNfsFileShare{}

	_jsii_.Create(
		"@cdktn/aws-storage-gateway.AwsStoragegatewayNfsFileShare",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share aws_storagegateway_nfs_file_share} Resource.
// Experimental.
func NewAwsStoragegatewayNfsFileShare_Override(a AwsStoragegatewayNfsFileShare, scope constructs.Construct, id *string, config *AwsStoragegatewayNfsFileShareConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-storage-gateway.AwsStoragegatewayNfsFileShare",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetAuditDestinationArn(val *string) {
	if err := j.validateSetAuditDestinationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditDestinationArn",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetBucketRegion(val *string) {
	if err := j.validateSetBucketRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketRegion",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetClientList(val *[]*string) {
	if err := j.validateSetClientListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientList",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetDefaultStorageClass(val *string) {
	if err := j.validateSetDefaultStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultStorageClass",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetFileShareName(val *string) {
	if err := j.validateSetFileShareNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileShareName",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetGatewayArn(val *string) {
	if err := j.validateSetGatewayArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayArn",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetGuessMimeTypeEnabled(val interface{}) {
	if err := j.validateSetGuessMimeTypeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guessMimeTypeEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetKmsEncrypted(val interface{}) {
	if err := j.validateSetKmsEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsEncrypted",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetLocationArn(val *string) {
	if err := j.validateSetLocationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationArn",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetNotificationPolicy(val *string) {
	if err := j.validateSetNotificationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetObjectAcl(val *string) {
	if err := j.validateSetObjectAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectAcl",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetRequesterPays(val interface{}) {
	if err := j.validateSetRequesterPaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requesterPays",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetSquash(val *string) {
	if err := j.validateSetSquashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"squash",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayNfsFileShare)SetVpcEndpointDnsName(val *string) {
	if err := j.validateSetVpcEndpointDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcEndpointDnsName",
		val,
	)
}

// Generates CDKTN code for importing a AwsStoragegatewayNfsFileShare resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsStoragegatewayNfsFileShare_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsStoragegatewayNfsFileShare_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-storage-gateway.AwsStoragegatewayNfsFileShare",
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
func AwsStoragegatewayNfsFileShare_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsStoragegatewayNfsFileShare_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-storage-gateway.AwsStoragegatewayNfsFileShare",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsStoragegatewayNfsFileShare_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsStoragegatewayNfsFileShare_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-storage-gateway.AwsStoragegatewayNfsFileShare",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsStoragegatewayNfsFileShare_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsStoragegatewayNfsFileShare_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-storage-gateway.AwsStoragegatewayNfsFileShare",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsStoragegatewayNfsFileShare_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-storage-gateway.AwsStoragegatewayNfsFileShare",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) PutCacheAttributes(value *AwsStoragegatewayNfsFileShare_CacheAttributesProperty) {
	if err := a.validatePutCacheAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCacheAttributes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) PutNfsFileShareDefaults(value *AwsStoragegatewayNfsFileShare_NfsFileShareDefaultsProperty) {
	if err := a.validatePutNfsFileShareDefaultsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNfsFileShareDefaults",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) PutTimeouts(value *AwsStoragegatewayNfsFileShare_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetAuditDestinationArn() {
	_jsii_.InvokeVoid(
		a,
		"resetAuditDestinationArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetBucketRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetBucketRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetCacheAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetDefaultStorageClass() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultStorageClass",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetFileShareName() {
	_jsii_.InvokeVoid(
		a,
		"resetFileShareName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetGuessMimeTypeEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetGuessMimeTypeEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetKmsEncrypted() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsEncrypted",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetNfsFileShareDefaults() {
	_jsii_.InvokeVoid(
		a,
		"resetNfsFileShareDefaults",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetNotificationPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetNotificationPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetObjectAcl() {
	_jsii_.InvokeVoid(
		a,
		"resetObjectAcl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetReadOnly() {
	_jsii_.InvokeVoid(
		a,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetRequesterPays() {
	_jsii_.InvokeVoid(
		a,
		"resetRequesterPays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetSquash() {
	_jsii_.InvokeVoid(
		a,
		"resetSquash",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ResetVpcEndpointDnsName() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcEndpointDnsName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStoragegatewayNfsFileShare) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

