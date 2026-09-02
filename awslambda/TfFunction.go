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
type TfFunction interface {
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
	CapacityProviderConfig() TfFunction_CapacityProviderConfigPropertyOutputReference
	// Experimental.
	CapacityProviderConfigInput() *TfFunction_CapacityProviderConfigProperty
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
	DeadLetterConfig() TfFunction_DeadLetterConfigPropertyOutputReference
	// Experimental.
	DeadLetterConfigInput() *TfFunction_DeadLetterConfigProperty
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
	DurableConfig() TfFunction_DurableConfigPropertyOutputReference
	// Experimental.
	DurableConfigInput() *TfFunction_DurableConfigProperty
	// Experimental.
	Environment() TfFunction_EnvironmentPropertyOutputReference
	// Experimental.
	EnvironmentInput() *TfFunction_EnvironmentProperty
	// Experimental.
	EphemeralStorage() TfFunction_EphemeralStoragePropertyOutputReference
	// Experimental.
	EphemeralStorageInput() *TfFunction_EphemeralStorageProperty
	// Experimental.
	Filename() *string
	// Experimental.
	SetFilename(val *string)
	// Experimental.
	FilenameInput() *string
	// Experimental.
	FileSystemConfig() TfFunction_FileSystemConfigPropertyOutputReference
	// Experimental.
	FileSystemConfigInput() *TfFunction_FileSystemConfigProperty
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
	ImageConfig() TfFunction_ImageConfigPropertyOutputReference
	// Experimental.
	ImageConfigInput() *TfFunction_ImageConfigProperty
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
	LoggingConfig() TfFunction_LoggingConfigPropertyOutputReference
	// Experimental.
	LoggingConfigInput() *TfFunction_LoggingConfigProperty
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
	SnapStart() TfFunction_SnapStartPropertyOutputReference
	// Experimental.
	SnapStartInput() *TfFunction_SnapStartProperty
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
	TenancyConfig() TfFunction_TenancyConfigPropertyOutputReference
	// Experimental.
	TenancyConfigInput() *TfFunction_TenancyConfigProperty
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
	Timeouts() TfFunction_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TracingConfig() TfFunction_TracingConfigPropertyOutputReference
	// Experimental.
	TracingConfigInput() *TfFunction_TracingConfigProperty
	// Experimental.
	UseResourceTimeoutForPropagation() interface{}
	// Experimental.
	SetUseResourceTimeoutForPropagation(val interface{})
	// Experimental.
	UseResourceTimeoutForPropagationInput() interface{}
	// Experimental.
	Version() *string
	// Experimental.
	VpcConfig() TfFunction_VpcConfigPropertyOutputReference
	// Experimental.
	VpcConfigInput() *TfFunction_VpcConfigProperty
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
	PutCapacityProviderConfig(value *TfFunction_CapacityProviderConfigProperty)
	// Experimental.
	PutDeadLetterConfig(value *TfFunction_DeadLetterConfigProperty)
	// Experimental.
	PutDurableConfig(value *TfFunction_DurableConfigProperty)
	// Experimental.
	PutEnvironment(value *TfFunction_EnvironmentProperty)
	// Experimental.
	PutEphemeralStorage(value *TfFunction_EphemeralStorageProperty)
	// Experimental.
	PutFileSystemConfig(value *TfFunction_FileSystemConfigProperty)
	// Experimental.
	PutImageConfig(value *TfFunction_ImageConfigProperty)
	// Experimental.
	PutLoggingConfig(value *TfFunction_LoggingConfigProperty)
	// Experimental.
	PutSnapStart(value *TfFunction_SnapStartProperty)
	// Experimental.
	PutTenancyConfig(value *TfFunction_TenancyConfigProperty)
	// Experimental.
	PutTimeouts(value *TfFunction_TimeoutsProperty)
	// Experimental.
	PutTracingConfig(value *TfFunction_TracingConfigProperty)
	// Experimental.
	PutVpcConfig(value *TfFunction_VpcConfigProperty)
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

// The jsii proxy struct for TfFunction
type jsiiProxy_TfFunction struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfFunction) Architectures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architectures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) ArchitecturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"architecturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) CapacityProviderConfig() TfFunction_CapacityProviderConfigPropertyOutputReference {
	var returns TfFunction_CapacityProviderConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"capacityProviderConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) CapacityProviderConfigInput() *TfFunction_CapacityProviderConfigProperty {
	var returns *TfFunction_CapacityProviderConfigProperty
	_jsii_.Get(
		j,
		"capacityProviderConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) CodeSha256() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSha256",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) CodeSha256Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSha256Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) CodeSigningConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) CodeSigningConfigArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"codeSigningConfigArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) DeadLetterConfig() TfFunction_DeadLetterConfigPropertyOutputReference {
	var returns TfFunction_DeadLetterConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"deadLetterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) DeadLetterConfigInput() *TfFunction_DeadLetterConfigProperty {
	var returns *TfFunction_DeadLetterConfigProperty
	_jsii_.Get(
		j,
		"deadLetterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) DurableConfig() TfFunction_DurableConfigPropertyOutputReference {
	var returns TfFunction_DurableConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"durableConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) DurableConfigInput() *TfFunction_DurableConfigProperty {
	var returns *TfFunction_DurableConfigProperty
	_jsii_.Get(
		j,
		"durableConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Environment() TfFunction_EnvironmentPropertyOutputReference {
	var returns TfFunction_EnvironmentPropertyOutputReference
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) EnvironmentInput() *TfFunction_EnvironmentProperty {
	var returns *TfFunction_EnvironmentProperty
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) EphemeralStorage() TfFunction_EphemeralStoragePropertyOutputReference {
	var returns TfFunction_EphemeralStoragePropertyOutputReference
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) EphemeralStorageInput() *TfFunction_EphemeralStorageProperty {
	var returns *TfFunction_EphemeralStorageProperty
	_jsii_.Get(
		j,
		"ephemeralStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Filename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) FilenameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) FileSystemConfig() TfFunction_FileSystemConfigPropertyOutputReference {
	var returns TfFunction_FileSystemConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"fileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) FileSystemConfigInput() *TfFunction_FileSystemConfigProperty {
	var returns *TfFunction_FileSystemConfigProperty
	_jsii_.Get(
		j,
		"fileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) HandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) ImageConfig() TfFunction_ImageConfigPropertyOutputReference {
	var returns TfFunction_ImageConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"imageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) ImageConfigInput() *TfFunction_ImageConfigProperty {
	var returns *TfFunction_ImageConfigProperty
	_jsii_.Get(
		j,
		"imageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) ImageUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) ImageUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) InvokeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"invokeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Layers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) LayersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) LoggingConfig() TfFunction_LoggingConfigPropertyOutputReference {
	var returns TfFunction_LoggingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"loggingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) LoggingConfigInput() *TfFunction_LoggingConfigProperty {
	var returns *TfFunction_LoggingConfigProperty
	_jsii_.Get(
		j,
		"loggingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) MemorySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) MemorySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memorySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) PackageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) PackageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Publish() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) PublishInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) PublishTo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publishTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) PublishToInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publishToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) QualifiedArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifiedArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) QualifiedInvokeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifiedInvokeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) ReplacementSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replacementSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) ReplacementSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replacementSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) ReplaceSecurityGroupsOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceSecurityGroupsOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) ReplaceSecurityGroupsOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceSecurityGroupsOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) ReservedConcurrentExecutions() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedConcurrentExecutions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) ReservedConcurrentExecutionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"reservedConcurrentExecutionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) ResponseStreamingInvokeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"responseStreamingInvokeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) S3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) S3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) S3Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) S3KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) S3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) S3ObjectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) SigningJobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingJobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) SigningProfileVersionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingProfileVersionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) SkipDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) SkipDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) SnapStart() TfFunction_SnapStartPropertyOutputReference {
	var returns TfFunction_SnapStartPropertyOutputReference
	_jsii_.Get(
		j,
		"snapStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) SnapStartInput() *TfFunction_SnapStartProperty {
	var returns *TfFunction_SnapStartProperty
	_jsii_.Get(
		j,
		"snapStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) SourceCodeHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodeHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) SourceCodeHashInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceCodeHashInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) SourceCodeSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sourceCodeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) SourceKmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceKmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) SourceKmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceKmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) TenancyConfig() TfFunction_TenancyConfigPropertyOutputReference {
	var returns TfFunction_TenancyConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"tenancyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) TenancyConfigInput() *TfFunction_TenancyConfigProperty {
	var returns *TfFunction_TenancyConfigProperty
	_jsii_.Get(
		j,
		"tenancyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Timeouts() TfFunction_TimeoutsPropertyOutputReference {
	var returns TfFunction_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) TracingConfig() TfFunction_TracingConfigPropertyOutputReference {
	var returns TfFunction_TracingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"tracingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) TracingConfigInput() *TfFunction_TracingConfigProperty {
	var returns *TfFunction_TracingConfigProperty
	_jsii_.Get(
		j,
		"tracingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) UseResourceTimeoutForPropagation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useResourceTimeoutForPropagation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) UseResourceTimeoutForPropagationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useResourceTimeoutForPropagationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) VpcConfig() TfFunction_VpcConfigPropertyOutputReference {
	var returns TfFunction_VpcConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFunction) VpcConfigInput() *TfFunction_VpcConfigProperty {
	var returns *TfFunction_VpcConfigProperty
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function aws_lambda_function} Resource.
// Experimental.
func NewTfFunction(scope constructs.Construct, id *string, config *TfFunctionConfig) TfFunction {
	_init_.Initialize()

	if err := validateNewTfFunctionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFunction{}

	_jsii_.Create(
		"@cdktn/aws-lambda.TfFunction",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function aws_lambda_function} Resource.
// Experimental.
func NewTfFunction_Override(t TfFunction, scope constructs.Construct, id *string, config *TfFunctionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lambda.TfFunction",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfFunction)SetArchitectures(val *[]*string) {
	if err := j.validateSetArchitecturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architectures",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetCodeSha256(val *string) {
	if err := j.validateSetCodeSha256Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeSha256",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetCodeSigningConfigArn(val *string) {
	if err := j.validateSetCodeSigningConfigArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"codeSigningConfigArn",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetFilename(val *string) {
	if err := j.validateSetFilenameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filename",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetFunctionName(val *string) {
	if err := j.validateSetFunctionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetHandler(val *string) {
	if err := j.validateSetHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetImageUri(val *string) {
	if err := j.validateSetImageUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageUri",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetLayers(val *[]*string) {
	if err := j.validateSetLayersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"layers",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetMemorySize(val *float64) {
	if err := j.validateSetMemorySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memorySize",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetPackageType(val *string) {
	if err := j.validateSetPackageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packageType",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetPublish(val interface{}) {
	if err := j.validateSetPublishParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publish",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetPublishTo(val *string) {
	if err := j.validateSetPublishToParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publishTo",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetReplacementSecurityGroupIds(val *[]*string) {
	if err := j.validateSetReplacementSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replacementSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetReplaceSecurityGroupsOnDestroy(val interface{}) {
	if err := j.validateSetReplaceSecurityGroupsOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceSecurityGroupsOnDestroy",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetReservedConcurrentExecutions(val *float64) {
	if err := j.validateSetReservedConcurrentExecutionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reservedConcurrentExecutions",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetS3Bucket(val *string) {
	if err := j.validateSetS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Bucket",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetS3Key(val *string) {
	if err := j.validateSetS3KeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Key",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetS3ObjectVersion(val *string) {
	if err := j.validateSetS3ObjectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetSkipDestroy(val interface{}) {
	if err := j.validateSetSkipDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipDestroy",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetSourceCodeHash(val *string) {
	if err := j.validateSetSourceCodeHashParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceCodeHash",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetSourceKmsKeyArn(val *string) {
	if err := j.validateSetSourceKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceKmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_TfFunction)SetUseResourceTimeoutForPropagation(val interface{}) {
	if err := j.validateSetUseResourceTimeoutForPropagationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useResourceTimeoutForPropagation",
		val,
	)
}

// Generates CDKTN code for importing a TfFunction resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfFunction_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfFunction_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-lambda.TfFunction",
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
func TfFunction_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfFunction_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lambda.TfFunction",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfFunction_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfFunction_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lambda.TfFunction",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfFunction_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfFunction_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lambda.TfFunction",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfFunction_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-lambda.TfFunction",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfFunction) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfFunction) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfFunction) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFunction) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFunction) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFunction) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFunction) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFunction) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFunction) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFunction) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFunction) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFunction) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFunction) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfFunction) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFunction) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfFunction) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfFunction) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfFunction) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfFunction) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfFunction) PutCapacityProviderConfig(value *TfFunction_CapacityProviderConfigProperty) {
	if err := t.validatePutCapacityProviderConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCapacityProviderConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFunction) PutDeadLetterConfig(value *TfFunction_DeadLetterConfigProperty) {
	if err := t.validatePutDeadLetterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeadLetterConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFunction) PutDurableConfig(value *TfFunction_DurableConfigProperty) {
	if err := t.validatePutDurableConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDurableConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFunction) PutEnvironment(value *TfFunction_EnvironmentProperty) {
	if err := t.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFunction) PutEphemeralStorage(value *TfFunction_EphemeralStorageProperty) {
	if err := t.validatePutEphemeralStorageParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEphemeralStorage",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFunction) PutFileSystemConfig(value *TfFunction_FileSystemConfigProperty) {
	if err := t.validatePutFileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFileSystemConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFunction) PutImageConfig(value *TfFunction_ImageConfigProperty) {
	if err := t.validatePutImageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putImageConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFunction) PutLoggingConfig(value *TfFunction_LoggingConfigProperty) {
	if err := t.validatePutLoggingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLoggingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFunction) PutSnapStart(value *TfFunction_SnapStartProperty) {
	if err := t.validatePutSnapStartParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSnapStart",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFunction) PutTenancyConfig(value *TfFunction_TenancyConfigProperty) {
	if err := t.validatePutTenancyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTenancyConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFunction) PutTimeouts(value *TfFunction_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFunction) PutTracingConfig(value *TfFunction_TracingConfigProperty) {
	if err := t.validatePutTracingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTracingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFunction) PutVpcConfig(value *TfFunction_VpcConfigProperty) {
	if err := t.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFunction) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfFunction) ResetArchitectures() {
	_jsii_.InvokeVoid(
		t,
		"resetArchitectures",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetCapacityProviderConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCapacityProviderConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetCodeSha256() {
	_jsii_.InvokeVoid(
		t,
		"resetCodeSha256",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetCodeSigningConfigArn() {
	_jsii_.InvokeVoid(
		t,
		"resetCodeSigningConfigArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetDeadLetterConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetDeadLetterConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetDurableConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetDurableConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetEnvironment() {
	_jsii_.InvokeVoid(
		t,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetEphemeralStorage() {
	_jsii_.InvokeVoid(
		t,
		"resetEphemeralStorage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetFilename() {
	_jsii_.InvokeVoid(
		t,
		"resetFilename",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetFileSystemConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetFileSystemConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetHandler() {
	_jsii_.InvokeVoid(
		t,
		"resetHandler",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetImageConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetImageConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetImageUri() {
	_jsii_.InvokeVoid(
		t,
		"resetImageUri",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetLayers() {
	_jsii_.InvokeVoid(
		t,
		"resetLayers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetLoggingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetLoggingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetMemorySize() {
	_jsii_.InvokeVoid(
		t,
		"resetMemorySize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetPackageType() {
	_jsii_.InvokeVoid(
		t,
		"resetPackageType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetPublish() {
	_jsii_.InvokeVoid(
		t,
		"resetPublish",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetPublishTo() {
	_jsii_.InvokeVoid(
		t,
		"resetPublishTo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetReplacementSecurityGroupIds() {
	_jsii_.InvokeVoid(
		t,
		"resetReplacementSecurityGroupIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetReplaceSecurityGroupsOnDestroy() {
	_jsii_.InvokeVoid(
		t,
		"resetReplaceSecurityGroupsOnDestroy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetReservedConcurrentExecutions() {
	_jsii_.InvokeVoid(
		t,
		"resetReservedConcurrentExecutions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetRuntime() {
	_jsii_.InvokeVoid(
		t,
		"resetRuntime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetS3Bucket() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Bucket",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetS3Key() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Key",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetS3ObjectVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetS3ObjectVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetSkipDestroy() {
	_jsii_.InvokeVoid(
		t,
		"resetSkipDestroy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetSnapStart() {
	_jsii_.InvokeVoid(
		t,
		"resetSnapStart",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetSourceCodeHash() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceCodeHash",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetSourceKmsKeyArn() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceKmsKeyArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetTenancyConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetTenancyConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetTracingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetTracingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetUseResourceTimeoutForPropagation() {
	_jsii_.InvokeVoid(
		t,
		"resetUseResourceTimeoutForPropagation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFunction) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFunction) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFunction) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFunction) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFunction) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFunction) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFunction) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

