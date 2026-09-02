package awsec2imagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline aws_imagebuilder_image_pipeline}.
// Experimental.
type TfImagePipeline interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	ContainerRecipeArn() *string
	// Experimental.
	SetContainerRecipeArn(val *string)
	// Experimental.
	ContainerRecipeArnInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DateCreated() *string
	// Experimental.
	DateLastRun() *string
	// Experimental.
	DateNextRun() *string
	// Experimental.
	DateUpdated() *string
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
	DistributionConfigurationArn() *string
	// Experimental.
	SetDistributionConfigurationArn(val *string)
	// Experimental.
	DistributionConfigurationArnInput() *string
	// Experimental.
	EnhancedImageMetadataEnabled() interface{}
	// Experimental.
	SetEnhancedImageMetadataEnabled(val interface{})
	// Experimental.
	EnhancedImageMetadataEnabledInput() interface{}
	// Experimental.
	ExecutionRole() *string
	// Experimental.
	SetExecutionRole(val *string)
	// Experimental.
	ExecutionRoleInput() *string
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
	ImageRecipeArn() *string
	// Experimental.
	SetImageRecipeArn(val *string)
	// Experimental.
	ImageRecipeArnInput() *string
	// Experimental.
	ImageScanningConfiguration() TfImagePipeline_ImageScanningConfigurationPropertyOutputReference
	// Experimental.
	ImageScanningConfigurationInput() *TfImagePipeline_ImageScanningConfigurationProperty
	// Experimental.
	ImageTestsConfiguration() TfImagePipeline_ImageTestsConfigurationPropertyOutputReference
	// Experimental.
	ImageTestsConfigurationInput() *TfImagePipeline_ImageTestsConfigurationProperty
	// Experimental.
	InfrastructureConfigurationArn() *string
	// Experimental.
	SetInfrastructureConfigurationArn(val *string)
	// Experimental.
	InfrastructureConfigurationArnInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LoggingConfiguration() TfImagePipeline_LoggingConfigurationPropertyOutputReference
	// Experimental.
	LoggingConfigurationInput() *TfImagePipeline_LoggingConfigurationProperty
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
	Platform() *string
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
	Schedule() TfImagePipeline_SchedulePropertyOutputReference
	// Experimental.
	ScheduleInput() *TfImagePipeline_ScheduleProperty
	// Experimental.
	Status() *string
	// Experimental.
	SetStatus(val *string)
	// Experimental.
	StatusInput() *string
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
	Workflow() TfImagePipeline_WorkflowPropertyList
	// Experimental.
	WorkflowInput() interface{}
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
	PutImageScanningConfiguration(value *TfImagePipeline_ImageScanningConfigurationProperty)
	// Experimental.
	PutImageTestsConfiguration(value *TfImagePipeline_ImageTestsConfigurationProperty)
	// Experimental.
	PutLoggingConfiguration(value *TfImagePipeline_LoggingConfigurationProperty)
	// Experimental.
	PutSchedule(value *TfImagePipeline_ScheduleProperty)
	// Experimental.
	PutWorkflow(value interface{})
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
	ResetContainerRecipeArn()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetDistributionConfigurationArn()
	// Experimental.
	ResetEnhancedImageMetadataEnabled()
	// Experimental.
	ResetExecutionRole()
	// Experimental.
	ResetId()
	// Experimental.
	ResetImageRecipeArn()
	// Experimental.
	ResetImageScanningConfiguration()
	// Experimental.
	ResetImageTestsConfiguration()
	// Experimental.
	ResetLoggingConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSchedule()
	// Experimental.
	ResetStatus()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetWorkflow()
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

// The jsii proxy struct for TfImagePipeline
type jsiiProxy_TfImagePipeline struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfImagePipeline) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) ContainerRecipeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRecipeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) ContainerRecipeArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRecipeArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) DateCreated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateCreated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) DateLastRun() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateLastRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) DateNextRun() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateNextRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) DateUpdated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dateUpdated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) DistributionConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) DistributionConfigurationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionConfigurationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) EnhancedImageMetadataEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedImageMetadataEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) EnhancedImageMetadataEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedImageMetadataEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) ImageRecipeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRecipeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) ImageRecipeArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRecipeArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) ImageScanningConfiguration() TfImagePipeline_ImageScanningConfigurationPropertyOutputReference {
	var returns TfImagePipeline_ImageScanningConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"imageScanningConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) ImageScanningConfigurationInput() *TfImagePipeline_ImageScanningConfigurationProperty {
	var returns *TfImagePipeline_ImageScanningConfigurationProperty
	_jsii_.Get(
		j,
		"imageScanningConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) ImageTestsConfiguration() TfImagePipeline_ImageTestsConfigurationPropertyOutputReference {
	var returns TfImagePipeline_ImageTestsConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"imageTestsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) ImageTestsConfigurationInput() *TfImagePipeline_ImageTestsConfigurationProperty {
	var returns *TfImagePipeline_ImageTestsConfigurationProperty
	_jsii_.Get(
		j,
		"imageTestsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) InfrastructureConfigurationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureConfigurationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) InfrastructureConfigurationArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureConfigurationArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) LoggingConfiguration() TfImagePipeline_LoggingConfigurationPropertyOutputReference {
	var returns TfImagePipeline_LoggingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"loggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) LoggingConfigurationInput() *TfImagePipeline_LoggingConfigurationProperty {
	var returns *TfImagePipeline_LoggingConfigurationProperty
	_jsii_.Get(
		j,
		"loggingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) Schedule() TfImagePipeline_SchedulePropertyOutputReference {
	var returns TfImagePipeline_SchedulePropertyOutputReference
	_jsii_.Get(
		j,
		"schedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) ScheduleInput() *TfImagePipeline_ScheduleProperty {
	var returns *TfImagePipeline_ScheduleProperty
	_jsii_.Get(
		j,
		"scheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) Workflow() TfImagePipeline_WorkflowPropertyList {
	var returns TfImagePipeline_WorkflowPropertyList
	_jsii_.Get(
		j,
		"workflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfImagePipeline) WorkflowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline aws_imagebuilder_image_pipeline} Resource.
// Experimental.
func NewTfImagePipeline(scope constructs.Construct, id *string, config *TfImagePipelineConfig) TfImagePipeline {
	_init_.Initialize()

	if err := validateNewTfImagePipelineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfImagePipeline{}

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.TfImagePipeline",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_image_pipeline aws_imagebuilder_image_pipeline} Resource.
// Experimental.
func NewTfImagePipeline_Override(t TfImagePipeline, scope constructs.Construct, id *string, config *TfImagePipelineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.TfImagePipeline",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetContainerRecipeArn(val *string) {
	if err := j.validateSetContainerRecipeArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerRecipeArn",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetDistributionConfigurationArn(val *string) {
	if err := j.validateSetDistributionConfigurationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"distributionConfigurationArn",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetEnhancedImageMetadataEnabled(val interface{}) {
	if err := j.validateSetEnhancedImageMetadataEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enhancedImageMetadataEnabled",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetImageRecipeArn(val *string) {
	if err := j.validateSetImageRecipeArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageRecipeArn",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetInfrastructureConfigurationArn(val *string) {
	if err := j.validateSetInfrastructureConfigurationArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"infrastructureConfigurationArn",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfImagePipeline)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a TfImagePipeline resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfImagePipeline_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfImagePipeline_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2-image-builder.TfImagePipeline",
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
func TfImagePipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfImagePipeline_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2-image-builder.TfImagePipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfImagePipeline_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfImagePipeline_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2-image-builder.TfImagePipeline",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfImagePipeline_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfImagePipeline_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-ec2-image-builder.TfImagePipeline",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfImagePipeline_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-ec2-image-builder.TfImagePipeline",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfImagePipeline) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfImagePipeline) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfImagePipeline) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfImagePipeline) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfImagePipeline) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfImagePipeline) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfImagePipeline) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfImagePipeline) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfImagePipeline) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfImagePipeline) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfImagePipeline) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfImagePipeline) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfImagePipeline) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfImagePipeline) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfImagePipeline) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfImagePipeline) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfImagePipeline) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfImagePipeline) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfImagePipeline) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfImagePipeline) PutImageScanningConfiguration(value *TfImagePipeline_ImageScanningConfigurationProperty) {
	if err := t.validatePutImageScanningConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putImageScanningConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfImagePipeline) PutImageTestsConfiguration(value *TfImagePipeline_ImageTestsConfigurationProperty) {
	if err := t.validatePutImageTestsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putImageTestsConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfImagePipeline) PutLoggingConfiguration(value *TfImagePipeline_LoggingConfigurationProperty) {
	if err := t.validatePutLoggingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLoggingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfImagePipeline) PutSchedule(value *TfImagePipeline_ScheduleProperty) {
	if err := t.validatePutScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSchedule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfImagePipeline) PutWorkflow(value interface{}) {
	if err := t.validatePutWorkflowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWorkflow",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfImagePipeline) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfImagePipeline) ResetContainerRecipeArn() {
	_jsii_.InvokeVoid(
		t,
		"resetContainerRecipeArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfImagePipeline) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfImagePipeline) ResetDistributionConfigurationArn() {
	_jsii_.InvokeVoid(
		t,
		"resetDistributionConfigurationArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfImagePipeline) ResetEnhancedImageMetadataEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetEnhancedImageMetadataEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfImagePipeline) ResetExecutionRole() {
	_jsii_.InvokeVoid(
		t,
		"resetExecutionRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfImagePipeline) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfImagePipeline) ResetImageRecipeArn() {
	_jsii_.InvokeVoid(
		t,
		"resetImageRecipeArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfImagePipeline) ResetImageScanningConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetImageScanningConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfImagePipeline) ResetImageTestsConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetImageTestsConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfImagePipeline) ResetLoggingConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetLoggingConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfImagePipeline) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfImagePipeline) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfImagePipeline) ResetSchedule() {
	_jsii_.InvokeVoid(
		t,
		"resetSchedule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfImagePipeline) ResetStatus() {
	_jsii_.InvokeVoid(
		t,
		"resetStatus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfImagePipeline) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfImagePipeline) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfImagePipeline) ResetWorkflow() {
	_jsii_.InvokeVoid(
		t,
		"resetWorkflow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfImagePipeline) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfImagePipeline) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfImagePipeline) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfImagePipeline) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfImagePipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfImagePipeline) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfImagePipeline) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

