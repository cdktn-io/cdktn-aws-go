package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslambda/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awslambda/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function aws_lambda_function}.
// Experimental.
type AwsLambdaFunction interface {
	cdktn.TerraformResource
	// Experimental.
	Architectures() *[]*string
	// Experimental.
	SetArchitectures(val *[]*string)
	// Experimental.
	ArchitecturesInput() *[]*string
	// Experimental.
	Arn() *string
	// Experimental.
	CapacityProviderConfig() AwsLambdaFunction_CapacityProviderConfigPropertyOutputReference
	// Experimental.
	CapacityProviderConfigInput() *AwsLambdaFunction_CapacityProviderConfigProperty
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CodeSha256() *string
	// Experimental.
	SetCodeSha256(val *string)
	// Experimental.
	CodeSha256Input() *string
	// Experimental.
	CodeSigningConfigArn() *string
	// Experimental.
	SetCodeSigningConfigArn(val *string)
	// Experimental.
	CodeSigningConfigArnInput() *string
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
	DeadLetterConfig() AwsLambdaFunction_DeadLetterConfigPropertyOutputReference
	// Experimental.
	DeadLetterConfigInput() *AwsLambdaFunction_DeadLetterConfigProperty
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
	DurableConfig() AwsLambdaFunction_DurableConfigPropertyOutputReference
	// Experimental.
	DurableConfigInput() *AwsLambdaFunction_DurableConfigProperty
	// Experimental.
	Environment() AwsLambdaFunction_EnvironmentPropertyOutputReference
	// Experimental.
	EnvironmentInput() *AwsLambdaFunction_EnvironmentProperty
	// Experimental.
	EphemeralStorage() AwsLambdaFunction_EphemeralStoragePropertyOutputReference
	// Experimental.
	EphemeralStorageInput() *AwsLambdaFunction_EphemeralStorageProperty
	// Experimental.
	Filename() *string
	// Experimental.
	SetFilename(val *string)
	// Experimental.
	FilenameInput() *string
	// Experimental.
	FileSystemConfig() AwsLambdaFunction_FileSystemConfigPropertyOutputReference
	// Experimental.
	FileSystemConfigInput() *AwsLambdaFunction_FileSystemConfigProperty
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	FunctionName() *string
	// Experimental.
	SetFunctionName(val *string)
	// Experimental.
	FunctionNameInput() *string
	// Experimental.
	Handler() *string
	// Experimental.
	SetHandler(val *string)
	// Experimental.
	HandlerInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	ImageConfig() AwsLambdaFunction_ImageConfigPropertyOutputReference
	// Experimental.
	ImageConfigInput() *AwsLambdaFunction_ImageConfigProperty
	// Experimental.
	ImageUri() *string
	// Experimental.
	SetImageUri(val *string)
	// Experimental.
	ImageUriInput() *string
	// Experimental.
	InvokeArn() *string
	// Experimental.
	KmsKeyArn() *string
	// Experimental.
	SetKmsKeyArn(val *string)
	// Experimental.
	KmsKeyArnInput() *string
	// Experimental.
	LastModified() *string
	// Experimental.
	Layers() *[]*string
	// Experimental.
	SetLayers(val *[]*string)
	// Experimental.
	LayersInput() *[]*string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LoggingConfig() AwsLambdaFunction_LoggingConfigPropertyOutputReference
	// Experimental.
	LoggingConfigInput() *AwsLambdaFunction_LoggingConfigProperty
	// Experimental.
	MemorySize() *float64
	// Experimental.
	SetMemorySize(val *float64)
	// Experimental.
	MemorySizeInput() *float64
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PackageType() *string
	// Experimental.
	SetPackageType(val *string)
	// Experimental.
	PackageTypeInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	Publish() interface{}
	// Experimental.
	SetPublish(val interface{})
	// Experimental.
	PublishInput() interface{}
	// Experimental.
	PublishTo() *string
	// Experimental.
	SetPublishTo(val *string)
	// Experimental.
	PublishToInput() *string
	// Experimental.
	QualifiedArn() *string
	// Experimental.
	QualifiedInvokeArn() *string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	ReplacementSecurityGroupIds() *[]*string
	// Experimental.
	SetReplacementSecurityGroupIds(val *[]*string)
	// Experimental.
	ReplacementSecurityGroupIdsInput() *[]*string
	// Experimental.
	ReplaceSecurityGroupsOnDestroy() interface{}
	// Experimental.
	SetReplaceSecurityGroupsOnDestroy(val interface{})
	// Experimental.
	ReplaceSecurityGroupsOnDestroyInput() interface{}
	// Experimental.
	ReservedConcurrentExecutions() *float64
	// Experimental.
	SetReservedConcurrentExecutions(val *float64)
	// Experimental.
	ReservedConcurrentExecutionsInput() *float64
	// Experimental.
	ResponseStreamingInvokeArn() *string
	// Experimental.
	Role() *string
	// Experimental.
	SetRole(val *string)
	// Experimental.
	RoleInput() *string
	// Experimental.
	Runtime() *string
	// Experimental.
	SetRuntime(val *string)
	// Experimental.
	RuntimeInput() *string
	// Experimental.
	S3Bucket() *string
	// Experimental.
	SetS3Bucket(val *string)
	// Experimental.
	S3BucketInput() *string
	// Experimental.
	S3Key() *string
	// Experimental.
	SetS3Key(val *string)
	// Experimental.
	S3KeyInput() *string
	// Experimental.
	S3ObjectVersion() *string
	// Experimental.
	SetS3ObjectVersion(val *string)
	// Experimental.
	S3ObjectVersionInput() *string
	// Experimental.
	SigningJobArn() *string
	// Experimental.
	SigningProfileVersionArn() *string
	// Experimental.
	SkipDestroy() interface{}
	// Experimental.
	SetSkipDestroy(val interface{})
	// Experimental.
	SkipDestroyInput() interface{}
	// Experimental.
	SnapStart() AwsLambdaFunction_SnapStartPropertyOutputReference
	// Experimental.
	SnapStartInput() *AwsLambdaFunction_SnapStartProperty
	// Experimental.
	SourceCodeHash() *string
	// Experimental.
	SetSourceCodeHash(val *string)
	// Experimental.
	SourceCodeHashInput() *string
	// Experimental.
	SourceCodeSize() *float64
	// Experimental.
	SourceKmsKeyArn() *string
	// Experimental.
	SetSourceKmsKeyArn(val *string)
	// Experimental.
	SourceKmsKeyArnInput() *string
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
	TenancyConfig() AwsLambdaFunction_TenancyConfigPropertyOutputReference
	// Experimental.
	TenancyConfigInput() *AwsLambdaFunction_TenancyConfigProperty
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeout() *float64
	// Experimental.
	SetTimeout(val *float64)
	// Experimental.
	TimeoutInput() *float64
	// Experimental.
	Timeouts() AwsLambdaFunction_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TracingConfig() AwsLambdaFunction_TracingConfigPropertyOutputReference
	// Experimental.
	TracingConfigInput() *AwsLambdaFunction_TracingConfigProperty
	// Experimental.
	UseResourceTimeoutForPropagation() interface{}
	// Experimental.
	SetUseResourceTimeoutForPropagation(val interface{})
	// Experimental.
	UseResourceTimeoutForPropagationInput() interface{}
	// Experimental.
	Version() *string
	// Experimental.
	VpcConfig() AwsLambdaFunction_VpcConfigPropertyOutputReference
	// Experimental.
	VpcConfigInput() *AwsLambdaFunction_VpcConfigProperty
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
	PutCapacityProviderConfig(value *AwsLambdaFunction_CapacityProviderConfigProperty)
	// Experimental.
	PutDeadLetterConfig(value *AwsLambdaFunction_DeadLetterConfigProperty)
	// Experimental.
	PutDurableConfig(value *AwsLambdaFunction_DurableConfigProperty)
	// Experimental.
	PutEnvironment(value *AwsLambdaFunction_EnvironmentProperty)
	// Experimental.
	PutEphemeralStorage(value *AwsLambdaFunction_EphemeralStorageProperty)
	// Experimental.
	PutFileSystemConfig(value *AwsLambdaFunction_FileSystemConfigProperty)
	// Experimental.
	PutImageConfig(value *AwsLambdaFunction_ImageConfigProperty)
	// Experimental.
	PutLoggingConfig(value *AwsLambdaFunction_LoggingConfigProperty)
	// Experimental.
	PutSnapStart(value *AwsLambdaFunction_SnapStartProperty)
	// Experimental.
	PutTenancyConfig(value *AwsLambdaFunction_TenancyConfigProperty)
	// Experimental.
	PutTimeouts(value *AwsLambdaFunction_TimeoutsProperty)
	// Experimental.
	PutTracingConfig(value *AwsLambdaFunction_TracingConfigProperty)
	// Experimental.
	PutVpcConfig(value *AwsLambdaFunction_VpcConfigProperty)
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
	ResetArchitectures()
	// Experimental.
	ResetCapacityProviderConfig()
	// Experimental.
	ResetCodeSha256()
	// Experimental.
	ResetCodeSigningConfigArn()
	// Experimental.
	ResetDeadLetterConfig()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetDurableConfig()
	// Experimental.
	ResetEnvironment()
	// Experimental.
	ResetEphemeralStorage()
	// Experimental.
	ResetFilename()
	// Experimental.
	ResetFileSystemConfig()
	// Experimental.
	ResetHandler()
	// Experimental.
	ResetId()
	// Experimental.
	ResetImageConfig()
	// Experimental.
	ResetImageUri()
	// Experimental.
	ResetKmsKeyArn()
	// Experimental.
	ResetLayers()
	// Experimental.
	ResetLoggingConfig()
	// Experimental.
	ResetMemorySize()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPackageType()
	// Experimental.
	ResetPublish()
	// Experimental.
	ResetPublishTo()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetReplacementSecurityGroupIds()
	// Experimental.
	ResetReplaceSecurityGroupsOnDestroy()
	// Experimental.
	ResetReservedConcurrentExecutions()
	// Experimental.
	ResetRuntime()
	// Experimental.
	ResetS3Bucket()
	// Experimental.
	ResetS3Key()
	// Experimental.
	ResetS3ObjectVersion()
	// Experimental.
	ResetSkipDestroy()
	// Experimental.
	ResetSnapStart()
	// Experimental.
	ResetSourceCodeHash()
	// Experimental.
	ResetSourceKmsKeyArn()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTenancyConfig()
	// Experimental.
	ResetTimeout()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTracingConfig()
	// Experimental.
	ResetUseResourceTimeoutForPropagation()
	// Experimental.
	ResetVpcConfig()
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

// The jsii proxy struct for AwsLambdaFunction
type jsiiProxy_AwsLambdaFunction struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsLambdaFunction) Architectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) ArchitecturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architecturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) CapacityProviderConfig() AwsLambdaFunction_CapacityProviderConfigPropertyOutputReference {
	var returns AwsLambdaFunction_CapacityProviderConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"capacityProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) CapacityProviderConfigInput() *AwsLambdaFunction_CapacityProviderConfigProperty {
	var returns *AwsLambdaFunction_CapacityProviderConfigProperty
	_jsii_.Get(
		j,
		"capacityProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) CodeSha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) CodeSha256Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSha256Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) CodeSigningConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) CodeSigningConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) DeadLetterConfig() AwsLambdaFunction_DeadLetterConfigPropertyOutputReference {
	var returns AwsLambdaFunction_DeadLetterConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) DeadLetterConfigInput() *AwsLambdaFunction_DeadLetterConfigProperty {
	var returns *AwsLambdaFunction_DeadLetterConfigProperty
	_jsii_.Get(
		j,
		"deadLetterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) DurableConfig() AwsLambdaFunction_DurableConfigPropertyOutputReference {
	var returns AwsLambdaFunction_DurableConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"durableConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) DurableConfigInput() *AwsLambdaFunction_DurableConfigProperty {
	var returns *AwsLambdaFunction_DurableConfigProperty
	_jsii_.Get(
		j,
		"durableConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Environment() AwsLambdaFunction_EnvironmentPropertyOutputReference {
	var returns AwsLambdaFunction_EnvironmentPropertyOutputReference
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) EnvironmentInput() *AwsLambdaFunction_EnvironmentProperty {
	var returns *AwsLambdaFunction_EnvironmentProperty
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) EphemeralStorage() AwsLambdaFunction_EphemeralStoragePropertyOutputReference {
	var returns AwsLambdaFunction_EphemeralStoragePropertyOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) EphemeralStorageInput() *AwsLambdaFunction_EphemeralStorageProperty {
	var returns *AwsLambdaFunction_EphemeralStorageProperty
	_jsii_.Get(
		j,
		"ephemeralStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Filename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) FilenameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) FileSystemConfig() AwsLambdaFunction_FileSystemConfigPropertyOutputReference {
	var returns AwsLambdaFunction_FileSystemConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"fileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) FileSystemConfigInput() *AwsLambdaFunction_FileSystemConfigProperty {
	var returns *AwsLambdaFunction_FileSystemConfigProperty
	_jsii_.Get(
		j,
		"fileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) HandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) ImageConfig() AwsLambdaFunction_ImageConfigPropertyOutputReference {
	var returns AwsLambdaFunction_ImageConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"imageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) ImageConfigInput() *AwsLambdaFunction_ImageConfigProperty {
	var returns *AwsLambdaFunction_ImageConfigProperty
	_jsii_.Get(
		j,
		"imageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) ImageUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) InvokeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Layers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) LayersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) LoggingConfig() AwsLambdaFunction_LoggingConfigPropertyOutputReference {
	var returns AwsLambdaFunction_LoggingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) LoggingConfigInput() *AwsLambdaFunction_LoggingConfigProperty {
	var returns *AwsLambdaFunction_LoggingConfigProperty
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) MemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) MemorySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) PackageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) PackageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Publish() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) PublishInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) PublishTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publishTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) PublishToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publishToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) QualifiedArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifiedArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) QualifiedInvokeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifiedInvokeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) ReplacementSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replacementSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) ReplacementSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replacementSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) ReplaceSecurityGroupsOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceSecurityGroupsOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) ReplaceSecurityGroupsOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceSecurityGroupsOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) ReservedConcurrentExecutions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedConcurrentExecutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) ReservedConcurrentExecutionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedConcurrentExecutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) ResponseStreamingInvokeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseStreamingInvokeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) S3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) S3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) S3Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) S3KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) S3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) S3ObjectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) SigningJobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingJobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) SigningProfileVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingProfileVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) SkipDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) SkipDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) SnapStart() AwsLambdaFunction_SnapStartPropertyOutputReference {
	var returns AwsLambdaFunction_SnapStartPropertyOutputReference
	_jsii_.Get(
		j,
		"snapStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) SnapStartInput() *AwsLambdaFunction_SnapStartProperty {
	var returns *AwsLambdaFunction_SnapStartProperty
	_jsii_.Get(
		j,
		"snapStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) SourceCodeHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodeHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) SourceCodeHashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodeHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) SourceCodeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceCodeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) SourceKmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceKmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) SourceKmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceKmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) TenancyConfig() AwsLambdaFunction_TenancyConfigPropertyOutputReference {
	var returns AwsLambdaFunction_TenancyConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"tenancyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) TenancyConfigInput() *AwsLambdaFunction_TenancyConfigProperty {
	var returns *AwsLambdaFunction_TenancyConfigProperty
	_jsii_.Get(
		j,
		"tenancyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Timeouts() AwsLambdaFunction_TimeoutsPropertyOutputReference {
	var returns AwsLambdaFunction_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) TracingConfig() AwsLambdaFunction_TracingConfigPropertyOutputReference {
	var returns AwsLambdaFunction_TracingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"tracingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) TracingConfigInput() *AwsLambdaFunction_TracingConfigProperty {
	var returns *AwsLambdaFunction_TracingConfigProperty
	_jsii_.Get(
		j,
		"tracingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) UseResourceTimeoutForPropagation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useResourceTimeoutForPropagation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) UseResourceTimeoutForPropagationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useResourceTimeoutForPropagationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) VpcConfig() AwsLambdaFunction_VpcConfigPropertyOutputReference {
	var returns AwsLambdaFunction_VpcConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaFunction) VpcConfigInput() *AwsLambdaFunction_VpcConfigProperty {
	var returns *AwsLambdaFunction_VpcConfigProperty
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function aws_lambda_function} Resource.
// Experimental.
func NewAwsLambdaFunction(scope constructs.Construct, id *string, config *AwsLambdaFunctionConfig) AwsLambdaFunction {
	_init_.Initialize()

	if err := validateNewAwsLambdaFunctionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLambdaFunction{}

	_jsii_.Create(
		"@cdktn/aws-lambda.AwsLambdaFunction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function aws_lambda_function} Resource.
// Experimental.
func NewAwsLambdaFunction_Override(a AwsLambdaFunction, scope constructs.Construct, id *string, config *AwsLambdaFunctionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lambda.AwsLambdaFunction",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetArchitectures(val *[]*string) {
	if err := j.validateSetArchitecturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architectures",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetCodeSha256(val *string) {
	if err := j.validateSetCodeSha256Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeSha256",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetCodeSigningConfigArn(val *string) {
	if err := j.validateSetCodeSigningConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeSigningConfigArn",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetFilename(val *string) {
	if err := j.validateSetFilenameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filename",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetFunctionName(val *string) {
	if err := j.validateSetFunctionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetHandler(val *string) {
	if err := j.validateSetHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetImageUri(val *string) {
	if err := j.validateSetImageUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetLayers(val *[]*string) {
	if err := j.validateSetLayersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"layers",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetMemorySize(val *float64) {
	if err := j.validateSetMemorySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySize",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetPackageType(val *string) {
	if err := j.validateSetPackageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packageType",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetPublish(val interface{}) {
	if err := j.validateSetPublishParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publish",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetPublishTo(val *string) {
	if err := j.validateSetPublishToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publishTo",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetReplacementSecurityGroupIds(val *[]*string) {
	if err := j.validateSetReplacementSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replacementSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetReplaceSecurityGroupsOnDestroy(val interface{}) {
	if err := j.validateSetReplaceSecurityGroupsOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceSecurityGroupsOnDestroy",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetReservedConcurrentExecutions(val *float64) {
	if err := j.validateSetReservedConcurrentExecutionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedConcurrentExecutions",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetS3Bucket(val *string) {
	if err := j.validateSetS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Bucket",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetS3Key(val *string) {
	if err := j.validateSetS3KeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Key",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetS3ObjectVersion(val *string) {
	if err := j.validateSetS3ObjectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetSkipDestroy(val interface{}) {
	if err := j.validateSetSkipDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipDestroy",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetSourceCodeHash(val *string) {
	if err := j.validateSetSourceCodeHashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCodeHash",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetSourceKmsKeyArn(val *string) {
	if err := j.validateSetSourceKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceKmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaFunction)SetUseResourceTimeoutForPropagation(val interface{}) {
	if err := j.validateSetUseResourceTimeoutForPropagationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useResourceTimeoutForPropagation",
		val,
	)
}

// Generates CDKTN code for importing a AwsLambdaFunction resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsLambdaFunction_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsLambdaFunction_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-lambda.AwsLambdaFunction",
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
func AwsLambdaFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsLambdaFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lambda.AwsLambdaFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsLambdaFunction_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsLambdaFunction_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lambda.AwsLambdaFunction",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsLambdaFunction_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsLambdaFunction_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lambda.AwsLambdaFunction",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsLambdaFunction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-lambda.AwsLambdaFunction",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsLambdaFunction) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLambdaFunction) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLambdaFunction) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLambdaFunction) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLambdaFunction) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLambdaFunction) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLambdaFunction) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLambdaFunction) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLambdaFunction) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLambdaFunction) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaFunction) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLambdaFunction) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsLambdaFunction) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) PutCapacityProviderConfig(value *AwsLambdaFunction_CapacityProviderConfigProperty) {
	if err := a.validatePutCapacityProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCapacityProviderConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) PutDeadLetterConfig(value *AwsLambdaFunction_DeadLetterConfigProperty) {
	if err := a.validatePutDeadLetterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeadLetterConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) PutDurableConfig(value *AwsLambdaFunction_DurableConfigProperty) {
	if err := a.validatePutDurableConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDurableConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) PutEnvironment(value *AwsLambdaFunction_EnvironmentProperty) {
	if err := a.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) PutEphemeralStorage(value *AwsLambdaFunction_EphemeralStorageProperty) {
	if err := a.validatePutEphemeralStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEphemeralStorage",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) PutFileSystemConfig(value *AwsLambdaFunction_FileSystemConfigProperty) {
	if err := a.validatePutFileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFileSystemConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) PutImageConfig(value *AwsLambdaFunction_ImageConfigProperty) {
	if err := a.validatePutImageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putImageConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) PutLoggingConfig(value *AwsLambdaFunction_LoggingConfigProperty) {
	if err := a.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) PutSnapStart(value *AwsLambdaFunction_SnapStartProperty) {
	if err := a.validatePutSnapStartParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnapStart",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) PutTenancyConfig(value *AwsLambdaFunction_TenancyConfigProperty) {
	if err := a.validatePutTenancyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTenancyConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) PutTimeouts(value *AwsLambdaFunction_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) PutTracingConfig(value *AwsLambdaFunction_TracingConfigProperty) {
	if err := a.validatePutTracingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTracingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) PutVpcConfig(value *AwsLambdaFunction_VpcConfigProperty) {
	if err := a.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetArchitectures() {
	_jsii_.InvokeVoid(
		a,
		"resetArchitectures",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetCapacityProviderConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCapacityProviderConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetCodeSha256() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeSha256",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetCodeSigningConfigArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeSigningConfigArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetDeadLetterConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDeadLetterConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetDurableConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDurableConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetEphemeralStorage() {
	_jsii_.InvokeVoid(
		a,
		"resetEphemeralStorage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetFilename() {
	_jsii_.InvokeVoid(
		a,
		"resetFilename",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetFileSystemConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetFileSystemConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetHandler() {
	_jsii_.InvokeVoid(
		a,
		"resetHandler",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetImageConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetImageConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetImageUri() {
	_jsii_.InvokeVoid(
		a,
		"resetImageUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetLayers() {
	_jsii_.InvokeVoid(
		a,
		"resetLayers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetMemorySize() {
	_jsii_.InvokeVoid(
		a,
		"resetMemorySize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetPackageType() {
	_jsii_.InvokeVoid(
		a,
		"resetPackageType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetPublish() {
	_jsii_.InvokeVoid(
		a,
		"resetPublish",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetPublishTo() {
	_jsii_.InvokeVoid(
		a,
		"resetPublishTo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetReplacementSecurityGroupIds() {
	_jsii_.InvokeVoid(
		a,
		"resetReplacementSecurityGroupIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetReplaceSecurityGroupsOnDestroy() {
	_jsii_.InvokeVoid(
		a,
		"resetReplaceSecurityGroupsOnDestroy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetReservedConcurrentExecutions() {
	_jsii_.InvokeVoid(
		a,
		"resetReservedConcurrentExecutions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetRuntime() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntime",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetS3Bucket() {
	_jsii_.InvokeVoid(
		a,
		"resetS3Bucket",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetS3Key() {
	_jsii_.InvokeVoid(
		a,
		"resetS3Key",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetS3ObjectVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetS3ObjectVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetSkipDestroy() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipDestroy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetSnapStart() {
	_jsii_.InvokeVoid(
		a,
		"resetSnapStart",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetSourceCodeHash() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceCodeHash",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetSourceKmsKeyArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceKmsKeyArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetTenancyConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetTenancyConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetTracingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetTracingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetUseResourceTimeoutForPropagation() {
	_jsii_.InvokeVoid(
		a,
		"resetUseResourceTimeoutForPropagation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaFunction) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaFunction) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaFunction) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaFunction) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaFunction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaFunction) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

