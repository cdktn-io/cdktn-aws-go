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
type AwsS3ObjectCopy interface {
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
	Grant() AwsS3ObjectCopy_GrantPropertyList
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
	OverrideProvider() AwsS3ObjectCopy_OverrideProviderPropertyOutputReference
	// Experimental.
	OverrideProviderInput() *AwsS3ObjectCopy_OverrideProviderProperty
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
	PutOverrideProvider(value *AwsS3ObjectCopy_OverrideProviderProperty)
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

// The jsii proxy struct for AwsS3ObjectCopy
type jsiiProxy_AwsS3ObjectCopy struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsS3ObjectCopy) Acl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"acl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) AclInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) BucketKeyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bucketKeyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) BucketKeyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bucketKeyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) CacheControl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheControl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) CacheControlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheControlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ChecksumAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ChecksumAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ChecksumCrc32() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumCrc32",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ChecksumCrc32C() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumCrc32C",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ChecksumCrc64Nvme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumCrc64Nvme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ChecksumSha1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumSha1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ChecksumSha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksumSha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ContentDisposition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentDisposition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ContentDispositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentDispositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ContentEncoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentEncoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ContentEncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentEncodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ContentLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ContentLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ContentType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ContentTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) CopyIfMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) CopyIfMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) CopyIfModifiedSince() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfModifiedSince",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) CopyIfModifiedSinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfModifiedSinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) CopyIfNoneMatch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfNoneMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) CopyIfNoneMatchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfNoneMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) CopyIfUnmodifiedSince() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfUnmodifiedSince",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) CopyIfUnmodifiedSinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyIfUnmodifiedSinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) CustomerAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) CustomerAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) CustomerKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) CustomerKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) CustomerKeyMd5() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerKeyMd5",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) CustomerKeyMd5Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerKeyMd5Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ExpectedBucketOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBucketOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ExpectedBucketOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedBucketOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ExpectedSourceBucketOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedSourceBucketOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ExpectedSourceBucketOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expectedSourceBucketOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Expiration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Expires() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expires",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ExpiresInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ForceDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ForceDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Grant() AwsS3ObjectCopy_GrantPropertyList {
	var returns AwsS3ObjectCopy_GrantPropertyList
	_jsii_.Get(
		j,
		"grant",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) GrantInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"grantInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) KmsEncryptionContext() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsEncryptionContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) KmsEncryptionContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsEncryptionContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) MetadataDirective() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataDirective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) MetadataDirectiveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataDirectiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ObjectLockLegalHoldStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockLegalHoldStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ObjectLockLegalHoldStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockLegalHoldStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ObjectLockMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ObjectLockModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ObjectLockRetainUntilDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockRetainUntilDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ObjectLockRetainUntilDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectLockRetainUntilDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) OverrideProvider() AwsS3ObjectCopy_OverrideProviderPropertyOutputReference {
	var returns AwsS3ObjectCopy_OverrideProviderPropertyOutputReference
	_jsii_.Get(
		j,
		"overrideProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) OverrideProviderInput() *AwsS3ObjectCopy_OverrideProviderProperty {
	var returns *AwsS3ObjectCopy_OverrideProviderProperty
	_jsii_.Get(
		j,
		"overrideProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) RequestCharged() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"requestCharged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) RequestPayer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestPayer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) RequestPayerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"requestPayerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ServerSideEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) ServerSideEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) SourceCustomerAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) SourceCustomerAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) SourceCustomerKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) SourceCustomerKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) SourceCustomerKeyMd5() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerKeyMd5",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) SourceCustomerKeyMd5Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCustomerKeyMd5Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) SourceVersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) StorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) StorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) TaggingDirective() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taggingDirective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) TaggingDirectiveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taggingDirectiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) WebsiteRedirect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteRedirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ObjectCopy) WebsiteRedirectInput() *string {
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
func NewAwsS3ObjectCopy(scope constructs.Construct, id *string, config *AwsS3ObjectCopyConfig) AwsS3ObjectCopy {
	_init_.Initialize()

	if err := validateNewAwsS3ObjectCopyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsS3ObjectCopy{}

	_jsii_.Create(
		"@cdktn/aws-s3.AwsS3ObjectCopy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy aws_s3_object_copy} Resource.
// Experimental.
func NewAwsS3ObjectCopy_Override(a AwsS3ObjectCopy, scope constructs.Construct, id *string, config *AwsS3ObjectCopyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.AwsS3ObjectCopy",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetAcl(val *string) {
	if err := j.validateSetAclParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acl",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetBucket(val *string) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetBucketKeyEnabled(val interface{}) {
	if err := j.validateSetBucketKeyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucketKeyEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetCacheControl(val *string) {
	if err := j.validateSetCacheControlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheControl",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetChecksumAlgorithm(val *string) {
	if err := j.validateSetChecksumAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checksumAlgorithm",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetContentDisposition(val *string) {
	if err := j.validateSetContentDispositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentDisposition",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetContentEncoding(val *string) {
	if err := j.validateSetContentEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentEncoding",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetContentLanguage(val *string) {
	if err := j.validateSetContentLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentLanguage",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetContentType(val *string) {
	if err := j.validateSetContentTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentType",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetCopyIfMatch(val *string) {
	if err := j.validateSetCopyIfMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyIfMatch",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetCopyIfModifiedSince(val *string) {
	if err := j.validateSetCopyIfModifiedSinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyIfModifiedSince",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetCopyIfNoneMatch(val *string) {
	if err := j.validateSetCopyIfNoneMatchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyIfNoneMatch",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetCopyIfUnmodifiedSince(val *string) {
	if err := j.validateSetCopyIfUnmodifiedSinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyIfUnmodifiedSince",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetCustomerAlgorithm(val *string) {
	if err := j.validateSetCustomerAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerAlgorithm",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetCustomerKey(val *string) {
	if err := j.validateSetCustomerKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerKey",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetCustomerKeyMd5(val *string) {
	if err := j.validateSetCustomerKeyMd5Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerKeyMd5",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetExpectedBucketOwner(val *string) {
	if err := j.validateSetExpectedBucketOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedBucketOwner",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetExpectedSourceBucketOwner(val *string) {
	if err := j.validateSetExpectedSourceBucketOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expectedSourceBucketOwner",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetExpires(val *string) {
	if err := j.validateSetExpiresParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expires",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetForceDestroy(val interface{}) {
	if err := j.validateSetForceDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceDestroy",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetKmsEncryptionContext(val *string) {
	if err := j.validateSetKmsEncryptionContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsEncryptionContext",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetMetadataDirective(val *string) {
	if err := j.validateSetMetadataDirectiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataDirective",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetObjectLockLegalHoldStatus(val *string) {
	if err := j.validateSetObjectLockLegalHoldStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockLegalHoldStatus",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetObjectLockMode(val *string) {
	if err := j.validateSetObjectLockModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockMode",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetObjectLockRetainUntilDate(val *string) {
	if err := j.validateSetObjectLockRetainUntilDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectLockRetainUntilDate",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetRequestPayer(val *string) {
	if err := j.validateSetRequestPayerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestPayer",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetServerSideEncryption(val *string) {
	if err := j.validateSetServerSideEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverSideEncryption",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetSource(val *string) {
	if err := j.validateSetSourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"source",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetSourceCustomerAlgorithm(val *string) {
	if err := j.validateSetSourceCustomerAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCustomerAlgorithm",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetSourceCustomerKey(val *string) {
	if err := j.validateSetSourceCustomerKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCustomerKey",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetSourceCustomerKeyMd5(val *string) {
	if err := j.validateSetSourceCustomerKeyMd5Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCustomerKeyMd5",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetStorageClass(val *string) {
	if err := j.validateSetStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageClass",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetTaggingDirective(val *string) {
	if err := j.validateSetTaggingDirectiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"taggingDirective",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsS3ObjectCopy)SetWebsiteRedirect(val *string) {
	if err := j.validateSetWebsiteRedirectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"websiteRedirect",
		val,
	)
}

// Generates CDKTN code for importing a AwsS3ObjectCopy resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsS3ObjectCopy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsS3ObjectCopy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-s3.AwsS3ObjectCopy",
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
func AwsS3ObjectCopy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsS3ObjectCopy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-s3.AwsS3ObjectCopy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsS3ObjectCopy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsS3ObjectCopy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-s3.AwsS3ObjectCopy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsS3ObjectCopy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsS3ObjectCopy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-s3.AwsS3ObjectCopy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsS3ObjectCopy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-s3.AwsS3ObjectCopy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsS3ObjectCopy) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsS3ObjectCopy) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3ObjectCopy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsS3ObjectCopy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsS3ObjectCopy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsS3ObjectCopy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsS3ObjectCopy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsS3ObjectCopy) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsS3ObjectCopy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsS3ObjectCopy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ObjectCopy) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3ObjectCopy) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsS3ObjectCopy) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) PutGrant(value interface{}) {
	if err := a.validatePutGrantParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGrant",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) PutOverrideProvider(value *AwsS3ObjectCopy_OverrideProviderProperty) {
	if err := a.validatePutOverrideProviderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOverrideProvider",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetAcl() {
	_jsii_.InvokeVoid(
		a,
		"resetAcl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetBucketKeyEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetBucketKeyEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetCacheControl() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheControl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetChecksumAlgorithm() {
	_jsii_.InvokeVoid(
		a,
		"resetChecksumAlgorithm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetContentDisposition() {
	_jsii_.InvokeVoid(
		a,
		"resetContentDisposition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetContentEncoding() {
	_jsii_.InvokeVoid(
		a,
		"resetContentEncoding",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetContentLanguage() {
	_jsii_.InvokeVoid(
		a,
		"resetContentLanguage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetContentType() {
	_jsii_.InvokeVoid(
		a,
		"resetContentType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetCopyIfMatch() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyIfMatch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetCopyIfModifiedSince() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyIfModifiedSince",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetCopyIfNoneMatch() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyIfNoneMatch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetCopyIfUnmodifiedSince() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyIfUnmodifiedSince",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetCustomerAlgorithm() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerAlgorithm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetCustomerKey() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetCustomerKeyMd5() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerKeyMd5",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetExpectedBucketOwner() {
	_jsii_.InvokeVoid(
		a,
		"resetExpectedBucketOwner",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetExpectedSourceBucketOwner() {
	_jsii_.InvokeVoid(
		a,
		"resetExpectedSourceBucketOwner",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetExpires() {
	_jsii_.InvokeVoid(
		a,
		"resetExpires",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetForceDestroy() {
	_jsii_.InvokeVoid(
		a,
		"resetForceDestroy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetGrant() {
	_jsii_.InvokeVoid(
		a,
		"resetGrant",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetKmsEncryptionContext() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsEncryptionContext",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetMetadata() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetMetadataDirective() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadataDirective",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetObjectLockLegalHoldStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetObjectLockLegalHoldStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetObjectLockMode() {
	_jsii_.InvokeVoid(
		a,
		"resetObjectLockMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetObjectLockRetainUntilDate() {
	_jsii_.InvokeVoid(
		a,
		"resetObjectLockRetainUntilDate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetOverrideProvider() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideProvider",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetRequestPayer() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestPayer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetServerSideEncryption() {
	_jsii_.InvokeVoid(
		a,
		"resetServerSideEncryption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetSourceCustomerAlgorithm() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceCustomerAlgorithm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetSourceCustomerKey() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceCustomerKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetSourceCustomerKeyMd5() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceCustomerKeyMd5",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetStorageClass() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageClass",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetTaggingDirective() {
	_jsii_.InvokeVoid(
		a,
		"resetTaggingDirective",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) ResetWebsiteRedirect() {
	_jsii_.InvokeVoid(
		a,
		"resetWebsiteRedirect",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ObjectCopy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ObjectCopy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ObjectCopy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ObjectCopy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ObjectCopy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ObjectCopy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ObjectCopy) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

