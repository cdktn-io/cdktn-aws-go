package awsstoragegateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsstoragegateway/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsstoragegateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share aws_storagegateway_smb_file_share}.
// Experimental.
type AwsStoragegatewaySmbFileShare interface {
	cdktn.TerraformResource
	// Experimental.
	AccessBasedEnumeration() interface{}
	// Experimental.
	SetAccessBasedEnumeration(val interface{})
	// Experimental.
	AccessBasedEnumerationInput() interface{}
	// Experimental.
	AdminUserList() *[]*string
	// Experimental.
	SetAdminUserList(val *[]*string)
	// Experimental.
	AdminUserListInput() *[]*string
	// Experimental.
	Arn() *string
	// Experimental.
	AuditDestinationArn() *string
	// Experimental.
	SetAuditDestinationArn(val *string)
	// Experimental.
	AuditDestinationArnInput() *string
	// Experimental.
	Authentication() *string
	// Experimental.
	SetAuthentication(val *string)
	// Experimental.
	AuthenticationInput() *string
	// Experimental.
	BucketRegion() *string
	// Experimental.
	SetBucketRegion(val *string)
	// Experimental.
	BucketRegionInput() *string
	// Experimental.
	CacheAttributes() AwsStoragegatewaySmbFileShare_CacheAttributesPropertyOutputReference
	// Experimental.
	CacheAttributesInput() *AwsStoragegatewaySmbFileShare_CacheAttributesProperty
	// Experimental.
	CaseSensitivity() *string
	// Experimental.
	SetCaseSensitivity(val *string)
	// Experimental.
	CaseSensitivityInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	InvalidUserList() *[]*string
	// Experimental.
	SetInvalidUserList(val *[]*string)
	// Experimental.
	InvalidUserListInput() *[]*string
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
	OplocksEnabled() interface{}
	// Experimental.
	SetOplocksEnabled(val interface{})
	// Experimental.
	OplocksEnabledInput() interface{}
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
	SmbAclEnabled() interface{}
	// Experimental.
	SetSmbAclEnabled(val interface{})
	// Experimental.
	SmbAclEnabledInput() interface{}
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
	Timeouts() AwsStoragegatewaySmbFileShare_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	ValidUserList() *[]*string
	// Experimental.
	SetValidUserList(val *[]*string)
	// Experimental.
	ValidUserListInput() *[]*string
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
	PutCacheAttributes(value *AwsStoragegatewaySmbFileShare_CacheAttributesProperty)
	// Experimental.
	PutTimeouts(value *AwsStoragegatewaySmbFileShare_TimeoutsProperty)
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
	ResetAccessBasedEnumeration()
	// Experimental.
	ResetAdminUserList()
	// Experimental.
	ResetAuditDestinationArn()
	// Experimental.
	ResetAuthentication()
	// Experimental.
	ResetBucketRegion()
	// Experimental.
	ResetCacheAttributes()
	// Experimental.
	ResetCaseSensitivity()
	// Experimental.
	ResetDefaultStorageClass()
	// Experimental.
	ResetFileShareName()
	// Experimental.
	ResetGuessMimeTypeEnabled()
	// Experimental.
	ResetId()
	// Experimental.
	ResetInvalidUserList()
	// Experimental.
	ResetKmsEncrypted()
	// Experimental.
	ResetKmsKeyArn()
	// Experimental.
	ResetNotificationPolicy()
	// Experimental.
	ResetObjectAcl()
	// Experimental.
	ResetOplocksEnabled()
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
	ResetSmbAclEnabled()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetValidUserList()
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

// The jsii proxy struct for AwsStoragegatewaySmbFileShare
type jsiiProxy_AwsStoragegatewaySmbFileShare struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) AccessBasedEnumeration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessBasedEnumeration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) AccessBasedEnumerationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accessBasedEnumerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) AdminUserList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminUserList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) AdminUserListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminUserListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) AuditDestinationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDestinationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) AuditDestinationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditDestinationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) Authentication() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) AuthenticationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) BucketRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) BucketRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) CacheAttributes() AwsStoragegatewaySmbFileShare_CacheAttributesPropertyOutputReference {
	var returns AwsStoragegatewaySmbFileShare_CacheAttributesPropertyOutputReference
	_jsii_.Get(
		j,
		"cacheAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) CacheAttributesInput() *AwsStoragegatewaySmbFileShare_CacheAttributesProperty {
	var returns *AwsStoragegatewaySmbFileShare_CacheAttributesProperty
	_jsii_.Get(
		j,
		"cacheAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) CaseSensitivity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caseSensitivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) CaseSensitivityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caseSensitivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) DefaultStorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStorageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) DefaultStorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultStorageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) FileshareId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileshareId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) FileShareName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) FileShareNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) GatewayArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) GuessMimeTypeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guessMimeTypeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) GuessMimeTypeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guessMimeTypeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) InvalidUserList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"invalidUserList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) InvalidUserListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"invalidUserListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) KmsEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) KmsEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) LocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) LocationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) NotificationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) NotificationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) ObjectAcl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectAcl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) ObjectAclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectAclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) OplocksEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oplocksEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) OplocksEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oplocksEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) RequesterPays() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) RequesterPaysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requesterPaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) SmbAclEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbAclEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) SmbAclEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"smbAclEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) Timeouts() AwsStoragegatewaySmbFileShare_TimeoutsPropertyOutputReference {
	var returns AwsStoragegatewaySmbFileShare_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) ValidUserList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validUserList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) ValidUserListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validUserListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) VpcEndpointDnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointDnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare) VpcEndpointDnsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcEndpointDnsNameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share aws_storagegateway_smb_file_share} Resource.
// Experimental.
func NewAwsStoragegatewaySmbFileShare(scope constructs.Construct, id *string, config *AwsStoragegatewaySmbFileShareConfig) AwsStoragegatewaySmbFileShare {
	_init_.Initialize()

	if err := validateNewAwsStoragegatewaySmbFileShareParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsStoragegatewaySmbFileShare{}

	_jsii_.Create(
		"@cdktn/aws-storage-gateway.AwsStoragegatewaySmbFileShare",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share aws_storagegateway_smb_file_share} Resource.
// Experimental.
func NewAwsStoragegatewaySmbFileShare_Override(a AwsStoragegatewaySmbFileShare, scope constructs.Construct, id *string, config *AwsStoragegatewaySmbFileShareConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-storage-gateway.AwsStoragegatewaySmbFileShare",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetAccessBasedEnumeration(val interface{}) {
	if err := j.validateSetAccessBasedEnumerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessBasedEnumeration",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetAdminUserList(val *[]*string) {
	if err := j.validateSetAdminUserListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminUserList",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetAuditDestinationArn(val *string) {
	if err := j.validateSetAuditDestinationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditDestinationArn",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetAuthentication(val *string) {
	if err := j.validateSetAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authentication",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetBucketRegion(val *string) {
	if err := j.validateSetBucketRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketRegion",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetCaseSensitivity(val *string) {
	if err := j.validateSetCaseSensitivityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caseSensitivity",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetDefaultStorageClass(val *string) {
	if err := j.validateSetDefaultStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultStorageClass",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetFileShareName(val *string) {
	if err := j.validateSetFileShareNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileShareName",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetGatewayArn(val *string) {
	if err := j.validateSetGatewayArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayArn",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetGuessMimeTypeEnabled(val interface{}) {
	if err := j.validateSetGuessMimeTypeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guessMimeTypeEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetInvalidUserList(val *[]*string) {
	if err := j.validateSetInvalidUserListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"invalidUserList",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetKmsEncrypted(val interface{}) {
	if err := j.validateSetKmsEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsEncrypted",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetLocationArn(val *string) {
	if err := j.validateSetLocationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locationArn",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetNotificationPolicy(val *string) {
	if err := j.validateSetNotificationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetObjectAcl(val *string) {
	if err := j.validateSetObjectAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectAcl",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetOplocksEnabled(val interface{}) {
	if err := j.validateSetOplocksEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oplocksEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetRequesterPays(val interface{}) {
	if err := j.validateSetRequesterPaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requesterPays",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetSmbAclEnabled(val interface{}) {
	if err := j.validateSetSmbAclEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smbAclEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetValidUserList(val *[]*string) {
	if err := j.validateSetValidUserListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validUserList",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewaySmbFileShare)SetVpcEndpointDnsName(val *string) {
	if err := j.validateSetVpcEndpointDnsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcEndpointDnsName",
		val,
	)
}

// Generates CDKTN code for importing a AwsStoragegatewaySmbFileShare resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsStoragegatewaySmbFileShare_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsStoragegatewaySmbFileShare_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-storage-gateway.AwsStoragegatewaySmbFileShare",
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
func AwsStoragegatewaySmbFileShare_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsStoragegatewaySmbFileShare_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-storage-gateway.AwsStoragegatewaySmbFileShare",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsStoragegatewaySmbFileShare_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsStoragegatewaySmbFileShare_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-storage-gateway.AwsStoragegatewaySmbFileShare",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsStoragegatewaySmbFileShare_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsStoragegatewaySmbFileShare_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-storage-gateway.AwsStoragegatewaySmbFileShare",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsStoragegatewaySmbFileShare_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-storage-gateway.AwsStoragegatewaySmbFileShare",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) PutCacheAttributes(value *AwsStoragegatewaySmbFileShare_CacheAttributesProperty) {
	if err := a.validatePutCacheAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCacheAttributes",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) PutTimeouts(value *AwsStoragegatewaySmbFileShare_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetAccessBasedEnumeration() {
	_jsii_.InvokeVoid(
		a,
		"resetAccessBasedEnumeration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetAdminUserList() {
	_jsii_.InvokeVoid(
		a,
		"resetAdminUserList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetAuditDestinationArn() {
	_jsii_.InvokeVoid(
		a,
		"resetAuditDestinationArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetAuthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthentication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetBucketRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetBucketRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetCacheAttributes() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheAttributes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetCaseSensitivity() {
	_jsii_.InvokeVoid(
		a,
		"resetCaseSensitivity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetDefaultStorageClass() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultStorageClass",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetFileShareName() {
	_jsii_.InvokeVoid(
		a,
		"resetFileShareName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetGuessMimeTypeEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetGuessMimeTypeEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetInvalidUserList() {
	_jsii_.InvokeVoid(
		a,
		"resetInvalidUserList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetKmsEncrypted() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsEncrypted",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetNotificationPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetNotificationPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetObjectAcl() {
	_jsii_.InvokeVoid(
		a,
		"resetObjectAcl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetOplocksEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetOplocksEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetReadOnly() {
	_jsii_.InvokeVoid(
		a,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetRequesterPays() {
	_jsii_.InvokeVoid(
		a,
		"resetRequesterPays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetSmbAclEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetSmbAclEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetValidUserList() {
	_jsii_.InvokeVoid(
		a,
		"resetValidUserList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ResetVpcEndpointDnsName() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcEndpointDnsName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStoragegatewaySmbFileShare) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

