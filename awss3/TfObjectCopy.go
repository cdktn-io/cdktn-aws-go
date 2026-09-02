package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awss3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy aws_s3_object_copy}.
// Experimental.
type TfObjectCopy interface {
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
	CopyIfMatch() *string
	// Experimental.
	SetCopyIfMatch(val *string)
	// Experimental.
	CopyIfMatchInput() *string
	// Experimental.
	CopyIfModifiedSince() *string
	// Experimental.
	SetCopyIfModifiedSince(val *string)
	// Experimental.
	CopyIfModifiedSinceInput() *string
	// Experimental.
	CopyIfNoneMatch() *string
	// Experimental.
	SetCopyIfNoneMatch(val *string)
	// Experimental.
	CopyIfNoneMatchInput() *string
	// Experimental.
	CopyIfUnmodifiedSince() *string
	// Experimental.
	SetCopyIfUnmodifiedSince(val *string)
	// Experimental.
	CopyIfUnmodifiedSinceInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	CustomerAlgorithm() *string
	// Experimental.
	SetCustomerAlgorithm(val *string)
	// Experimental.
	CustomerAlgorithmInput() *string
	// Experimental.
	CustomerKey() *string
	// Experimental.
	SetCustomerKey(val *string)
	// Experimental.
	CustomerKeyInput() *string
	// Experimental.
	CustomerKeyMd5() *string
	// Experimental.
	SetCustomerKeyMd5(val *string)
	// Experimental.
	CustomerKeyMd5Input() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	Etag() *string
	// Experimental.
	ExpectedBucketOwner() *string
	// Experimental.
	SetExpectedBucketOwner(val *string)
	// Experimental.
	ExpectedBucketOwnerInput() *string
	// Experimental.
	ExpectedSourceBucketOwner() *string
	// Experimental.
	SetExpectedSourceBucketOwner(val *string)
	// Experimental.
	ExpectedSourceBucketOwnerInput() *string
	// Experimental.
	Expiration() *string
	// Experimental.
	Expires() *string
	// Experimental.
	SetExpires(val *string)
	// Experimental.
	ExpiresInput() *string
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
	Grant() TfObjectCopy_GrantPropertyList
	// Experimental.
	GrantInput() interface{}
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
	KmsEncryptionContext() *string
	// Experimental.
	SetKmsEncryptionContext(val *string)
	// Experimental.
	KmsEncryptionContextInput() *string
	// Experimental.
	KmsKeyId() *string
	// Experimental.
	SetKmsKeyId(val *string)
	// Experimental.
	KmsKeyIdInput() *string
	// Experimental.
	LastModified() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	Metadata() *map[string]*string
	// Experimental.
	SetMetadata(val *map[string]*string)
	// Experimental.
	MetadataDirective() *string
	// Experimental.
	SetMetadataDirective(val *string)
	// Experimental.
	MetadataDirectiveInput() *string
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
	OverrideProvider() TfObjectCopy_OverrideProviderPropertyOutputReference
	// Experimental.
	OverrideProviderInput() *TfObjectCopy_OverrideProviderProperty
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
	RequestCharged() cdktn.IResolvable
	// Experimental.
	RequestPayer() *string
	// Experimental.
	SetRequestPayer(val *string)
	// Experimental.
	RequestPayerInput() *string
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
	SourceCustomerAlgorithm() *string
	// Experimental.
	SetSourceCustomerAlgorithm(val *string)
	// Experimental.
	SourceCustomerAlgorithmInput() *string
	// Experimental.
	SourceCustomerKey() *string
	// Experimental.
	SetSourceCustomerKey(val *string)
	// Experimental.
	SourceCustomerKeyInput() *string
	// Experimental.
	SourceCustomerKeyMd5() *string
	// Experimental.
	SetSourceCustomerKeyMd5(val *string)
	// Experimental.
	SourceCustomerKeyMd5Input() *string
	// Experimental.
	SourceInput() *string
	// Experimental.
	SourceVersionId() *string
	// Experimental.
	StorageClass() *string
	// Experimental.
	SetStorageClass(val *string)
	// Experimental.
	StorageClassInput() *string
	// Experimental.
	TaggingDirective() *string
	// Experimental.
	SetTaggingDirective(val *string)
	// Experimental.
	TaggingDirectiveInput() *string
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
	PutGrant(value interface{})
	// Experimental.
	PutOverrideProvider(value *TfObjectCopy_OverrideProviderProperty)
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
	ResetContentDisposition()
	// Experimental.
	ResetContentEncoding()
	// Experimental.
	ResetContentLanguage()
	// Experimental.
	ResetContentType()
	// Experimental.
	ResetCopyIfMatch()
	// Experimental.
	ResetCopyIfModifiedSince()
	// Experimental.
	ResetCopyIfNoneMatch()
	// Experimental.
	ResetCopyIfUnmodifiedSince()
	// Experimental.
	ResetCustomerAlgorithm()
	// Experimental.
	ResetCustomerKey()
	// Experimental.
	ResetCustomerKeyMd5()
	// Experimental.
	ResetExpectedBucketOwner()
	// Experimental.
	ResetExpectedSourceBucketOwner()
	// Experimental.
	ResetExpires()
	// Experimental.
	ResetForceDestroy()
	// Experimental.
	ResetGrant()
	// Experimental.
	ResetId()
	// Experimental.
	ResetKmsEncryptionContext()
	// Experimental.
	ResetKmsKeyId()
	// Experimental.
	ResetMetadata()
	// Experimental.
	ResetMetadataDirective()
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
	ResetRequestPayer()
	// Experimental.
	ResetServerSideEncryption()
	// Experimental.
	ResetSourceCustomerAlgorithm()
	// Experimental.
	ResetSourceCustomerKey()
	// Experimental.
	ResetSourceCustomerKeyMd5()
	// Experimental.
	ResetStorageClass()
	// Experimental.
	ResetTaggingDirective()
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

// The jsii proxy struct for TfObjectCopy
type jsiiProxy_TfObjectCopy struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfObjectCopy) Acl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) AclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) BucketKeyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bucketKeyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) BucketKeyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bucketKeyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) CacheControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) CacheControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ChecksumAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ChecksumAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ChecksumCrc32() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumCrc32",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ChecksumCrc32C() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumCrc32C",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ChecksumCrc64Nvme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumCrc64Nvme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ChecksumSha1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumSha1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ChecksumSha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumSha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ContentDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ContentDispositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentDispositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ContentEncoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ContentEncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentEncodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ContentLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ContentLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) CopyIfMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) CopyIfMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) CopyIfModifiedSince() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfModifiedSince",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) CopyIfModifiedSinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfModifiedSinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) CopyIfNoneMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfNoneMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) CopyIfNoneMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfNoneMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) CopyIfUnmodifiedSince() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfUnmodifiedSince",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) CopyIfUnmodifiedSinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfUnmodifiedSinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) CustomerAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) CustomerAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) CustomerKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) CustomerKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) CustomerKeyMd5() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerKeyMd5",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) CustomerKeyMd5Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerKeyMd5Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ExpectedBucketOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBucketOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ExpectedBucketOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBucketOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ExpectedSourceBucketOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedSourceBucketOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ExpectedSourceBucketOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedSourceBucketOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Expiration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Expires() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expires",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ExpiresInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Grant() TfObjectCopy_GrantPropertyList {
	var returns TfObjectCopy_GrantPropertyList
	_jsii_.Get(
		j,
		"grant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) GrantInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) KmsEncryptionContext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsEncryptionContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) KmsEncryptionContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsEncryptionContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) MetadataDirective() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataDirective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) MetadataDirectiveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataDirectiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ObjectLockLegalHoldStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockLegalHoldStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ObjectLockLegalHoldStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockLegalHoldStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ObjectLockMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ObjectLockModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ObjectLockRetainUntilDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockRetainUntilDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ObjectLockRetainUntilDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockRetainUntilDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) OverrideProvider() TfObjectCopy_OverrideProviderPropertyOutputReference {
	var returns TfObjectCopy_OverrideProviderPropertyOutputReference
	_jsii_.Get(
		j,
		"overrideProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) OverrideProviderInput() *TfObjectCopy_OverrideProviderProperty {
	var returns *TfObjectCopy_OverrideProviderProperty
	_jsii_.Get(
		j,
		"overrideProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) RequestCharged() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"requestCharged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) RequestPayer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestPayer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) RequestPayerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestPayerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ServerSideEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) ServerSideEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) SourceCustomerAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) SourceCustomerAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) SourceCustomerKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) SourceCustomerKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) SourceCustomerKeyMd5() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerKeyMd5",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) SourceCustomerKeyMd5Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerKeyMd5Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) SourceVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) TaggingDirective() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taggingDirective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) TaggingDirectiveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taggingDirectiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) WebsiteRedirect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfObjectCopy) WebsiteRedirectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteRedirectInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy aws_s3_object_copy} Resource.
// Experimental.
func NewTfObjectCopy(scope constructs.Construct, id *string, config *TfObjectCopyConfig) TfObjectCopy {
	_init_.Initialize()

	if err := validateNewTfObjectCopyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfObjectCopy{}

	_jsii_.Create(
		"@cdktn/aws-s3.TfObjectCopy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy aws_s3_object_copy} Resource.
// Experimental.
func NewTfObjectCopy_Override(t TfObjectCopy, scope constructs.Construct, id *string, config *TfObjectCopyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.TfObjectCopy",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetAcl(val *string) {
	if err := j.validateSetAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acl",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetBucketKeyEnabled(val interface{}) {
	if err := j.validateSetBucketKeyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketKeyEnabled",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetCacheControl(val *string) {
	if err := j.validateSetCacheControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheControl",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetChecksumAlgorithm(val *string) {
	if err := j.validateSetChecksumAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checksumAlgorithm",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetContentDisposition(val *string) {
	if err := j.validateSetContentDispositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentDisposition",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetContentEncoding(val *string) {
	if err := j.validateSetContentEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentEncoding",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetContentLanguage(val *string) {
	if err := j.validateSetContentLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentLanguage",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetCopyIfMatch(val *string) {
	if err := j.validateSetCopyIfMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyIfMatch",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetCopyIfModifiedSince(val *string) {
	if err := j.validateSetCopyIfModifiedSinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyIfModifiedSince",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetCopyIfNoneMatch(val *string) {
	if err := j.validateSetCopyIfNoneMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyIfNoneMatch",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetCopyIfUnmodifiedSince(val *string) {
	if err := j.validateSetCopyIfUnmodifiedSinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyIfUnmodifiedSince",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetCustomerAlgorithm(val *string) {
	if err := j.validateSetCustomerAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerAlgorithm",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetCustomerKey(val *string) {
	if err := j.validateSetCustomerKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerKey",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetCustomerKeyMd5(val *string) {
	if err := j.validateSetCustomerKeyMd5Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerKeyMd5",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetExpectedBucketOwner(val *string) {
	if err := j.validateSetExpectedBucketOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedBucketOwner",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetExpectedSourceBucketOwner(val *string) {
	if err := j.validateSetExpectedSourceBucketOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedSourceBucketOwner",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetExpires(val *string) {
	if err := j.validateSetExpiresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expires",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetKmsEncryptionContext(val *string) {
	if err := j.validateSetKmsEncryptionContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsEncryptionContext",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetMetadataDirective(val *string) {
	if err := j.validateSetMetadataDirectiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataDirective",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetObjectLockLegalHoldStatus(val *string) {
	if err := j.validateSetObjectLockLegalHoldStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockLegalHoldStatus",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetObjectLockMode(val *string) {
	if err := j.validateSetObjectLockModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockMode",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetObjectLockRetainUntilDate(val *string) {
	if err := j.validateSetObjectLockRetainUntilDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockRetainUntilDate",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetRequestPayer(val *string) {
	if err := j.validateSetRequestPayerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestPayer",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetServerSideEncryption(val *string) {
	if err := j.validateSetServerSideEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverSideEncryption",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetSourceCustomerAlgorithm(val *string) {
	if err := j.validateSetSourceCustomerAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCustomerAlgorithm",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetSourceCustomerKey(val *string) {
	if err := j.validateSetSourceCustomerKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCustomerKey",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetSourceCustomerKeyMd5(val *string) {
	if err := j.validateSetSourceCustomerKeyMd5Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCustomerKeyMd5",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetTaggingDirective(val *string) {
	if err := j.validateSetTaggingDirectiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taggingDirective",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfObjectCopy)SetWebsiteRedirect(val *string) {
	if err := j.validateSetWebsiteRedirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"websiteRedirect",
		val,
	)
}

// Generates CDKTN code for importing a TfObjectCopy resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfObjectCopy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfObjectCopy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-s3.TfObjectCopy",
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
func TfObjectCopy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfObjectCopy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-s3.TfObjectCopy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfObjectCopy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfObjectCopy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-s3.TfObjectCopy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfObjectCopy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfObjectCopy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-s3.TfObjectCopy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfObjectCopy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-s3.TfObjectCopy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfObjectCopy) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfObjectCopy) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfObjectCopy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfObjectCopy) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfObjectCopy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfObjectCopy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfObjectCopy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfObjectCopy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfObjectCopy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfObjectCopy) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfObjectCopy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfObjectCopy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObjectCopy) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfObjectCopy) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfObjectCopy) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfObjectCopy) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfObjectCopy) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfObjectCopy) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfObjectCopy) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfObjectCopy) PutGrant(value interface{}) {
	if err := t.validatePutGrantParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGrant",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfObjectCopy) PutOverrideProvider(value *TfObjectCopy_OverrideProviderProperty) {
	if err := t.validatePutOverrideProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOverrideProvider",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfObjectCopy) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetAcl() {
	_jsii_.InvokeVoid(
		t,
		"resetAcl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetBucketKeyEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetBucketKeyEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetCacheControl() {
	_jsii_.InvokeVoid(
		t,
		"resetCacheControl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetChecksumAlgorithm() {
	_jsii_.InvokeVoid(
		t,
		"resetChecksumAlgorithm",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetContentDisposition() {
	_jsii_.InvokeVoid(
		t,
		"resetContentDisposition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetContentEncoding() {
	_jsii_.InvokeVoid(
		t,
		"resetContentEncoding",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetContentLanguage() {
	_jsii_.InvokeVoid(
		t,
		"resetContentLanguage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetContentType() {
	_jsii_.InvokeVoid(
		t,
		"resetContentType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetCopyIfMatch() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyIfMatch",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetCopyIfModifiedSince() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyIfModifiedSince",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetCopyIfNoneMatch() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyIfNoneMatch",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetCopyIfUnmodifiedSince() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyIfUnmodifiedSince",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetCustomerAlgorithm() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomerAlgorithm",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetCustomerKey() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomerKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetCustomerKeyMd5() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomerKeyMd5",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetExpectedBucketOwner() {
	_jsii_.InvokeVoid(
		t,
		"resetExpectedBucketOwner",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetExpectedSourceBucketOwner() {
	_jsii_.InvokeVoid(
		t,
		"resetExpectedSourceBucketOwner",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetExpires() {
	_jsii_.InvokeVoid(
		t,
		"resetExpires",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		t,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetGrant() {
	_jsii_.InvokeVoid(
		t,
		"resetGrant",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetKmsEncryptionContext() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsEncryptionContext",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetMetadata() {
	_jsii_.InvokeVoid(
		t,
		"resetMetadata",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetMetadataDirective() {
	_jsii_.InvokeVoid(
		t,
		"resetMetadataDirective",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetObjectLockLegalHoldStatus() {
	_jsii_.InvokeVoid(
		t,
		"resetObjectLockLegalHoldStatus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetObjectLockMode() {
	_jsii_.InvokeVoid(
		t,
		"resetObjectLockMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetObjectLockRetainUntilDate() {
	_jsii_.InvokeVoid(
		t,
		"resetObjectLockRetainUntilDate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetOverrideProvider() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideProvider",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetRequestPayer() {
	_jsii_.InvokeVoid(
		t,
		"resetRequestPayer",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetServerSideEncryption() {
	_jsii_.InvokeVoid(
		t,
		"resetServerSideEncryption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetSourceCustomerAlgorithm() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceCustomerAlgorithm",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetSourceCustomerKey() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceCustomerKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetSourceCustomerKeyMd5() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceCustomerKeyMd5",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetStorageClass() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageClass",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetTaggingDirective() {
	_jsii_.InvokeVoid(
		t,
		"resetTaggingDirective",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) ResetWebsiteRedirect() {
	_jsii_.InvokeVoid(
		t,
		"resetWebsiteRedirect",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfObjectCopy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObjectCopy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObjectCopy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObjectCopy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObjectCopy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObjectCopy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfObjectCopy) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

