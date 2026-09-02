package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awss3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object aws_s3_object}.
// Experimental.
type TfObject interface {
	cdktn.TerraformResource
	// Experimental.
	Acl() *string
	// Experimental.
	SetAcl(val *string)
	// Experimental.
	AclInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	Bucket() *string
	// Experimental.
	SetBucket(val *string)
	// Experimental.
	BucketInput() *string
	// Experimental.
	BucketKeyEnabled() interface{}
	// Experimental.
	SetBucketKeyEnabled(val interface{})
	// Experimental.
	BucketKeyEnabledInput() interface{}
	// Experimental.
	CacheControl() *string
	// Experimental.
	SetCacheControl(val *string)
	// Experimental.
	CacheControlInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ChecksumAlgorithm() *string
	// Experimental.
	SetChecksumAlgorithm(val *string)
	// Experimental.
	ChecksumAlgorithmInput() *string
	// Experimental.
	ChecksumCrc32() *string
	// Experimental.
	ChecksumCrc32C() *string
	// Experimental.
	ChecksumCrc64Nvme() *string
	// Experimental.
	ChecksumSha1() *string
	// Experimental.
	ChecksumSha256() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Content() *string
	// Experimental.
	SetContent(val *string)
	// Experimental.
	ContentBase64() *string
	// Experimental.
	SetContentBase64(val *string)
	// Experimental.
	ContentBase64Input() *string
	// Experimental.
	ContentDisposition() *string
	// Experimental.
	SetContentDisposition(val *string)
	// Experimental.
	ContentDispositionInput() *string
	// Experimental.
	ContentEncoding() *string
	// Experimental.
	SetContentEncoding(val *string)
	// Experimental.
	ContentEncodingInput() *string
	// Experimental.
	ContentInput() *string
	// Experimental.
	ContentLanguage() *string
	// Experimental.
	SetContentLanguage(val *string)
	// Experimental.
	ContentLanguageInput() *string
	// Experimental.
	ContentType() *string
	// Experimental.
	SetContentType(val *string)
	// Experimental.
	ContentTypeInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	Etag() *string
	// Experimental.
	SetEtag(val *string)
	// Experimental.
	EtagInput() *string
	// Experimental.
	ForceDestroy() interface{}
	// Experimental.
	SetForceDestroy(val interface{})
	// Experimental.
	ForceDestroyInput() interface{}
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
	Key() *string
	// Experimental.
	SetKey(val *string)
	// Experimental.
	KeyInput() *string
	// Experimental.
	KmsKeyId() *string
	// Experimental.
	SetKmsKeyId(val *string)
	// Experimental.
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	Metadata() *map[string]*string
	// Experimental.
	SetMetadata(val *map[string]*string)
	// Experimental.
	MetadataInput() *map[string]*string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	ObjectLockLegalHoldStatus() *string
	// Experimental.
	SetObjectLockLegalHoldStatus(val *string)
	// Experimental.
	ObjectLockLegalHoldStatusInput() *string
	// Experimental.
	ObjectLockMode() *string
	// Experimental.
	SetObjectLockMode(val *string)
	// Experimental.
	ObjectLockModeInput() *string
	// Experimental.
	ObjectLockRetainUntilDate() *string
	// Experimental.
	SetObjectLockRetainUntilDate(val *string)
	// Experimental.
	ObjectLockRetainUntilDateInput() *string
	// Experimental.
	OverrideProvider() TfObject_OverrideProviderPropertyOutputReference
	// Experimental.
	OverrideProviderInput() *TfObject_OverrideProviderProperty
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
	ServerSideEncryption() *string
	// Experimental.
	SetServerSideEncryption(val *string)
	// Experimental.
	ServerSideEncryptionInput() *string
	// Experimental.
	Source() *string
	// Experimental.
	SetSource(val *string)
	// Experimental.
	SourceHash() *string
	// Experimental.
	SetSourceHash(val *string)
	// Experimental.
	SourceHashInput() *string
	// Experimental.
	SourceInput() *string
	// Experimental.
	StorageClass() *string
	// Experimental.
	SetStorageClass(val *string)
	// Experimental.
	StorageClassInput() *string
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
	VersionId() *string
	// Experimental.
	WebsiteRedirect() *string
	// Experimental.
	SetWebsiteRedirect(val *string)
	// Experimental.
	WebsiteRedirectInput() *string
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
	PutOverrideProvider(value *TfObject_OverrideProviderProperty)
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
	ResetAcl()
	// Experimental.
	ResetBucketKeyEnabled()
	// Experimental.
	ResetCacheControl()
	// Experimental.
	ResetChecksumAlgorithm()
	// Experimental.
	ResetContent()
	// Experimental.
	ResetContentBase64()
	// Experimental.
	ResetContentDisposition()
	// Experimental.
	ResetContentEncoding()
	// Experimental.
	ResetContentLanguage()
	// Experimental.
	ResetContentType()
	// Experimental.
	ResetEtag()
	// Experimental.
	ResetForceDestroy()
	// Experimental.
	ResetId()
	// Experimental.
	ResetKmsKeyId()
	// Experimental.
	ResetMetadata()
	// Experimental.
	ResetObjectLockLegalHoldStatus()
	// Experimental.
	ResetObjectLockMode()
	// Experimental.
	ResetObjectLockRetainUntilDate()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetOverrideProvider()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetServerSideEncryption()
	// Experimental.
	ResetSource()
	// Experimental.
	ResetSourceHash()
	// Experimental.
	ResetStorageClass()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetWebsiteRedirect()
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

// The jsii proxy struct for TfObject
type jsiiProxy_TfObject struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfObject) Acl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) AclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) BucketKeyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bucketKeyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) BucketKeyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bucketKeyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) CacheControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) CacheControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ChecksumAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ChecksumAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ChecksumCrc32() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumCrc32",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ChecksumCrc32C() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumCrc32C",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ChecksumCrc64Nvme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumCrc64Nvme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ChecksumSha1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumSha1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ChecksumSha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumSha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ContentBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ContentBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ContentDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ContentDispositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentDispositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ContentEncoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ContentEncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentEncodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ContentLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ContentLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) EtagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ObjectLockLegalHoldStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockLegalHoldStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ObjectLockLegalHoldStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockLegalHoldStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ObjectLockMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ObjectLockModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ObjectLockRetainUntilDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockRetainUntilDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ObjectLockRetainUntilDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockRetainUntilDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) OverrideProvider() TfObject_OverrideProviderPropertyOutputReference {
	var returns TfObject_OverrideProviderPropertyOutputReference
	_jsii_.Get(
		j,
		"overrideProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) OverrideProviderInput() *TfObject_OverrideProviderProperty {
	var returns *TfObject_OverrideProviderProperty
	_jsii_.Get(
		j,
		"overrideProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ServerSideEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) ServerSideEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) SourceHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) SourceHashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) WebsiteRedirect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObject) WebsiteRedirectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteRedirectInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object aws_s3_object} Resource.
// Experimental.
func NewTfObject(scope constructs.Construct, id *string, config *TfObjectConfig) TfObject {
	_init_.Initialize()

	if err := validateNewTfObjectParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfObject{}

	_jsii_.Create(
		"@cdktn/aws-s3.TfObject",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object aws_s3_object} Resource.
// Experimental.
func NewTfObject_Override(t TfObject, scope constructs.Construct, id *string, config *TfObjectConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.TfObject",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfObject)SetAcl(val *string) {
	if err := j.validateSetAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acl",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetBucketKeyEnabled(val interface{}) {
	if err := j.validateSetBucketKeyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketKeyEnabled",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetCacheControl(val *string) {
	if err := j.validateSetCacheControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheControl",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetChecksumAlgorithm(val *string) {
	if err := j.validateSetChecksumAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checksumAlgorithm",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetContent(val *string) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetContentBase64(val *string) {
	if err := j.validateSetContentBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentBase64",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetContentDisposition(val *string) {
	if err := j.validateSetContentDispositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentDisposition",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetContentEncoding(val *string) {
	if err := j.validateSetContentEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentEncoding",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetContentLanguage(val *string) {
	if err := j.validateSetContentLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentLanguage",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetEtag(val *string) {
	if err := j.validateSetEtagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etag",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetObjectLockLegalHoldStatus(val *string) {
	if err := j.validateSetObjectLockLegalHoldStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockLegalHoldStatus",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetObjectLockMode(val *string) {
	if err := j.validateSetObjectLockModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockMode",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetObjectLockRetainUntilDate(val *string) {
	if err := j.validateSetObjectLockRetainUntilDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockRetainUntilDate",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetServerSideEncryption(val *string) {
	if err := j.validateSetServerSideEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverSideEncryption",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetSourceHash(val *string) {
	if err := j.validateSetSourceHashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceHash",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfObject)SetWebsiteRedirect(val *string) {
	if err := j.validateSetWebsiteRedirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"websiteRedirect",
		val,
	)
}

// Generates CDKTN code for importing a TfObject resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfObject_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfObject_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-s3.TfObject",
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
func TfObject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfObject_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-s3.TfObject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfObject_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfObject_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-s3.TfObject",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfObject_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfObject_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-s3.TfObject",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfObject_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-s3.TfObject",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfObject) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfObject) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfObject) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfObject) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := t.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfObject) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfObject) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfObject) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfObject) PutOverrideProvider(value *TfObject_OverrideProviderProperty) {
	if err := t.validatePutOverrideProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOverrideProvider",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfObject) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfObject) ResetAcl() {
	_jsii_.InvokeVoid(
		t,
		"resetAcl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetBucketKeyEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetBucketKeyEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetCacheControl() {
	_jsii_.InvokeVoid(
		t,
		"resetCacheControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetChecksumAlgorithm() {
	_jsii_.InvokeVoid(
		t,
		"resetChecksumAlgorithm",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetContent() {
	_jsii_.InvokeVoid(
		t,
		"resetContent",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetContentBase64() {
	_jsii_.InvokeVoid(
		t,
		"resetContentBase64",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetContentDisposition() {
	_jsii_.InvokeVoid(
		t,
		"resetContentDisposition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetContentEncoding() {
	_jsii_.InvokeVoid(
		t,
		"resetContentEncoding",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetContentLanguage() {
	_jsii_.InvokeVoid(
		t,
		"resetContentLanguage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetContentType() {
	_jsii_.InvokeVoid(
		t,
		"resetContentType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetEtag() {
	_jsii_.InvokeVoid(
		t,
		"resetEtag",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		t,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetMetadata() {
	_jsii_.InvokeVoid(
		t,
		"resetMetadata",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetObjectLockLegalHoldStatus() {
	_jsii_.InvokeVoid(
		t,
		"resetObjectLockLegalHoldStatus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetObjectLockMode() {
	_jsii_.InvokeVoid(
		t,
		"resetObjectLockMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetObjectLockRetainUntilDate() {
	_jsii_.InvokeVoid(
		t,
		"resetObjectLockRetainUntilDate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetOverrideProvider() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideProvider",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetServerSideEncryption() {
	_jsii_.InvokeVoid(
		t,
		"resetServerSideEncryption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetSource() {
	_jsii_.InvokeVoid(
		t,
		"resetSource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetSourceHash() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceHash",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetStorageClass() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageClass",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) ResetWebsiteRedirect() {
	_jsii_.InvokeVoid(
		t,
		"resetWebsiteRedirect",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObject) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObject) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		t,
		"with",
		args,
		&returns,
	)

	return returns
}

