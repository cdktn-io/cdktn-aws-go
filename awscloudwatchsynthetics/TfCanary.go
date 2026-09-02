package awscloudwatchsynthetics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudwatchsynthetics/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awscloudwatchsynthetics/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary aws_synthetics_canary}.
// Experimental.
type TfCanary interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	ArtifactConfig() TfCanary_ArtifactConfigPropertyOutputReference
	// Experimental.
	ArtifactConfigInput() *TfCanary_ArtifactConfigProperty
	// Experimental.
	ArtifactS3Location() *string
	// Experimental.
	SetArtifactS3Location(val *string)
	// Experimental.
	ArtifactS3LocationInput() *string
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
	DeleteLambda() interface{}
	// Experimental.
	SetDeleteLambda(val interface{})
	// Experimental.
	DeleteLambdaInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EngineArn() *string
	// Experimental.
	ExecutionRoleArn() *string
	// Experimental.
	SetExecutionRoleArn(val *string)
	// Experimental.
	ExecutionRoleArnInput() *string
	// Experimental.
	FailureRetentionPeriod() *float64
	// Experimental.
	SetFailureRetentionPeriod(val *float64)
	// Experimental.
	FailureRetentionPeriodInput() *float64
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
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
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	RunConfig() TfCanary_RunConfigPropertyOutputReference
	// Experimental.
	RunConfigInput() *TfCanary_RunConfigProperty
	// Experimental.
	RuntimeVersion() *string
	// Experimental.
	SetRuntimeVersion(val *string)
	// Experimental.
	RuntimeVersionInput() *string
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
	S3Version() *string
	// Experimental.
	SetS3Version(val *string)
	// Experimental.
	S3VersionInput() *string
	// Experimental.
	Schedule() TfCanary_SchedulePropertyOutputReference
	// Experimental.
	ScheduleInput() *TfCanary_ScheduleProperty
	// Experimental.
	SourceLocationArn() *string
	// Experimental.
	StartCanary() interface{}
	// Experimental.
	SetStartCanary(val interface{})
	// Experimental.
	StartCanaryInput() interface{}
	// Experimental.
	Status() *string
	// Experimental.
	SuccessRetentionPeriod() *float64
	// Experimental.
	SetSuccessRetentionPeriod(val *float64)
	// Experimental.
	SuccessRetentionPeriodInput() *float64
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
	Timeline() TfCanary_TimelinePropertyList
	// Experimental.
	VpcConfig() TfCanary_VpcConfigPropertyOutputReference
	// Experimental.
	VpcConfigInput() *TfCanary_VpcConfigProperty
	// Experimental.
	ZipFile() *string
	// Experimental.
	SetZipFile(val *string)
	// Experimental.
	ZipFileInput() *string
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
	PutArtifactConfig(value *TfCanary_ArtifactConfigProperty)
	// Experimental.
	PutRunConfig(value *TfCanary_RunConfigProperty)
	// Experimental.
	PutSchedule(value *TfCanary_ScheduleProperty)
	// Experimental.
	PutVpcConfig(value *TfCanary_VpcConfigProperty)
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
	ResetArtifactConfig()
	// Experimental.
	ResetDeleteLambda()
	// Experimental.
	ResetFailureRetentionPeriod()
	// Experimental.
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRunConfig()
	// Experimental.
	ResetS3Bucket()
	// Experimental.
	ResetS3Key()
	// Experimental.
	ResetS3Version()
	// Experimental.
	ResetStartCanary()
	// Experimental.
	ResetSuccessRetentionPeriod()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetVpcConfig()
	// Experimental.
	ResetZipFile()
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

// The jsii proxy struct for TfCanary
type jsiiProxy_TfCanary struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfCanary) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) ArtifactConfig() TfCanary_ArtifactConfigPropertyOutputReference {
	var returns TfCanary_ArtifactConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"artifactConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) ArtifactConfigInput() *TfCanary_ArtifactConfigProperty {
	var returns *TfCanary_ArtifactConfigProperty
	_jsii_.Get(
		j,
		"artifactConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) ArtifactS3Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactS3Location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) ArtifactS3LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"artifactS3LocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) DeleteLambda() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteLambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) DeleteLambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteLambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) EngineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) FailureRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) FailureRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) Handler() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) HandlerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"handlerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) RunConfig() TfCanary_RunConfigPropertyOutputReference {
	var returns TfCanary_RunConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"runConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) RunConfigInput() *TfCanary_RunConfigProperty {
	var returns *TfCanary_RunConfigProperty
	_jsii_.Get(
		j,
		"runConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) S3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) S3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) S3Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) S3KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) S3Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) S3VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3VersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) Schedule() TfCanary_SchedulePropertyOutputReference {
	var returns TfCanary_SchedulePropertyOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) ScheduleInput() *TfCanary_ScheduleProperty {
	var returns *TfCanary_ScheduleProperty
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) SourceLocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceLocationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) StartCanary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startCanary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) StartCanaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startCanaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) SuccessRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) SuccessRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"successRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) Timeline() TfCanary_TimelinePropertyList {
	var returns TfCanary_TimelinePropertyList
	_jsii_.Get(
		j,
		"timeline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) VpcConfig() TfCanary_VpcConfigPropertyOutputReference {
	var returns TfCanary_VpcConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) VpcConfigInput() *TfCanary_VpcConfigProperty {
	var returns *TfCanary_VpcConfigProperty
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) ZipFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary) ZipFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipFileInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary aws_synthetics_canary} Resource.
// Experimental.
func NewTfCanary(scope constructs.Construct, id *string, config *TfCanaryConfig) TfCanary {
	_init_.Initialize()

	if err := validateNewTfCanaryParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCanary{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-synthetics.TfCanary",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/synthetics_canary aws_synthetics_canary} Resource.
// Experimental.
func NewTfCanary_Override(t TfCanary, scope constructs.Construct, id *string, config *TfCanaryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-synthetics.TfCanary",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfCanary)SetArtifactS3Location(val *string) {
	if err := j.validateSetArtifactS3LocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"artifactS3Location",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetDeleteLambda(val interface{}) {
	if err := j.validateSetDeleteLambdaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteLambda",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetFailureRetentionPeriod(val *float64) {
	if err := j.validateSetFailureRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetHandler(val *string) {
	if err := j.validateSetHandlerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"handler",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetS3Bucket(val *string) {
	if err := j.validateSetS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Bucket",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetS3Key(val *string) {
	if err := j.validateSetS3KeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Key",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetS3Version(val *string) {
	if err := j.validateSetS3VersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Version",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetStartCanary(val interface{}) {
	if err := j.validateSetStartCanaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startCanary",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetSuccessRetentionPeriod(val *float64) {
	if err := j.validateSetSuccessRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfCanary)SetZipFile(val *string) {
	if err := j.validateSetZipFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zipFile",
		val,
	)
}

// Generates CDKTN code for importing a TfCanary resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfCanary_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfCanary_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudwatch-synthetics.TfCanary",
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
func TfCanary_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCanary_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudwatch-synthetics.TfCanary",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCanary_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCanary_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudwatch-synthetics.TfCanary",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCanary_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCanary_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudwatch-synthetics.TfCanary",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfCanary_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-cloudwatch-synthetics.TfCanary",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfCanary) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfCanary) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfCanary) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCanary) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCanary) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCanary) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCanary) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCanary) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCanary) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCanary) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCanary) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCanary) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCanary) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfCanary) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCanary) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfCanary) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCanary) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfCanary) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCanary) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfCanary) PutArtifactConfig(value *TfCanary_ArtifactConfigProperty) {
	if err := t.validatePutArtifactConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putArtifactConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCanary) PutRunConfig(value *TfCanary_RunConfigProperty) {
	if err := t.validatePutRunConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRunConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCanary) PutSchedule(value *TfCanary_ScheduleProperty) {
	if err := t.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSchedule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCanary) PutVpcConfig(value *TfCanary_VpcConfigProperty) {
	if err := t.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCanary) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfCanary) ResetArtifactConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetArtifactConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary) ResetDeleteLambda() {
	_jsii_.InvokeVoid(
		t,
		"resetDeleteLambda",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary) ResetFailureRetentionPeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetFailureRetentionPeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary) ResetRunConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetRunConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary) ResetS3Bucket() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Bucket",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary) ResetS3Key() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Key",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary) ResetS3Version() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Version",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary) ResetStartCanary() {
	_jsii_.InvokeVoid(
		t,
		"resetStartCanary",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary) ResetSuccessRetentionPeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetSuccessRetentionPeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary) ResetZipFile() {
	_jsii_.InvokeVoid(
		t,
		"resetZipFile",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCanary) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCanary) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCanary) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCanary) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCanary) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCanary) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

