package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/glue/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/glue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job aws_glue_job}.
// Experimental.
type AwsJob interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Command() AwsJob_CommandPropertyOutputReference
	// Experimental.
	CommandInput() *AwsJob_CommandProperty
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	Connections() *[]*string
	// Experimental.
	SetConnections(val *[]*string)
	// Experimental.
	ConnectionsInput() *[]*string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DefaultArguments() *map[string]*string
	// Experimental.
	SetDefaultArguments(val *map[string]*string)
	// Experimental.
	DefaultArgumentsInput() *map[string]*string
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
	ExecutionClass() *string
	// Experimental.
	SetExecutionClass(val *string)
	// Experimental.
	ExecutionClassInput() *string
	// Experimental.
	ExecutionProperty() AwsJob_ExecutionPropertyPropertyOutputReference
	// Experimental.
	ExecutionPropertyInput() *AwsJob_ExecutionPropertyProperty
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	GlueVersion() *string
	// Experimental.
	SetGlueVersion(val *string)
	// Experimental.
	GlueVersionInput() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	JobMode() *string
	// Experimental.
	SetJobMode(val *string)
	// Experimental.
	JobModeInput() *string
	// Experimental.
	JobRunQueuingEnabled() interface{}
	// Experimental.
	SetJobRunQueuingEnabled(val interface{})
	// Experimental.
	JobRunQueuingEnabledInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaintenanceWindow() *string
	// Experimental.
	SetMaintenanceWindow(val *string)
	// Experimental.
	MaintenanceWindowInput() *string
	// Experimental.
	MaxCapacity() *float64
	// Experimental.
	SetMaxCapacity(val *float64)
	// Experimental.
	MaxCapacityInput() *float64
	// Experimental.
	MaxRetries() *float64
	// Experimental.
	SetMaxRetries(val *float64)
	// Experimental.
	MaxRetriesInput() *float64
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
	NonOverridableArguments() *map[string]*string
	// Experimental.
	SetNonOverridableArguments(val *map[string]*string)
	// Experimental.
	NonOverridableArgumentsInput() *map[string]*string
	// Experimental.
	NotificationProperty() AwsJob_NotificationPropertyPropertyOutputReference
	// Experimental.
	NotificationPropertyInput() *AwsJob_NotificationPropertyProperty
	// Experimental.
	NumberOfWorkers() *float64
	// Experimental.
	SetNumberOfWorkers(val *float64)
	// Experimental.
	NumberOfWorkersInput() *float64
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
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	SecurityConfiguration() *string
	// Experimental.
	SetSecurityConfiguration(val *string)
	// Experimental.
	SecurityConfigurationInput() *string
	// Experimental.
	SourceControlDetails() AwsJob_SourceControlDetailsPropertyOutputReference
	// Experimental.
	SourceControlDetailsInput() *AwsJob_SourceControlDetailsProperty
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
	Timeout() *float64
	// Experimental.
	SetTimeout(val *float64)
	// Experimental.
	TimeoutInput() *float64
	// Experimental.
	WorkerType() *string
	// Experimental.
	SetWorkerType(val *string)
	// Experimental.
	WorkerTypeInput() *string
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
	PutCommand(value *AwsJob_CommandProperty)
	// Experimental.
	PutExecutionProperty(value *AwsJob_ExecutionPropertyProperty)
	// Experimental.
	PutNotificationProperty(value *AwsJob_NotificationPropertyProperty)
	// Experimental.
	PutSourceControlDetails(value *AwsJob_SourceControlDetailsProperty)
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
	ResetConnections()
	// Experimental.
	ResetDefaultArguments()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetExecutionClass()
	// Experimental.
	ResetExecutionProperty()
	// Experimental.
	ResetGlueVersion()
	// Experimental.
	ResetId()
	// Experimental.
	ResetJobMode()
	// Experimental.
	ResetJobRunQueuingEnabled()
	// Experimental.
	ResetMaintenanceWindow()
	// Experimental.
	ResetMaxCapacity()
	// Experimental.
	ResetMaxRetries()
	// Experimental.
	ResetNonOverridableArguments()
	// Experimental.
	ResetNotificationProperty()
	// Experimental.
	ResetNumberOfWorkers()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSecurityConfiguration()
	// Experimental.
	ResetSourceControlDetails()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeout()
	// Experimental.
	ResetWorkerType()
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

// The jsii proxy struct for AwsJob
type jsiiProxy_AwsJob struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsJob) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) Command() AwsJob_CommandPropertyOutputReference {
	var returns AwsJob_CommandPropertyOutputReference
	_jsii_.Get(
		j,
		"command",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) CommandInput() *AwsJob_CommandProperty {
	var returns *AwsJob_CommandProperty
	_jsii_.Get(
		j,
		"commandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) Connections() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) ConnectionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) DefaultArguments() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) DefaultArgumentsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"defaultArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) ExecutionClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) ExecutionClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) ExecutionProperty() AwsJob_ExecutionPropertyPropertyOutputReference {
	var returns AwsJob_ExecutionPropertyPropertyOutputReference
	_jsii_.Get(
		j,
		"executionProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) ExecutionPropertyInput() *AwsJob_ExecutionPropertyProperty {
	var returns *AwsJob_ExecutionPropertyProperty
	_jsii_.Get(
		j,
		"executionPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) GlueVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) GlueVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"glueVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) JobMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) JobModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) JobRunQueuingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobRunQueuingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) JobRunQueuingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jobRunQueuingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) MaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) MaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) MaxCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) NonOverridableArguments() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nonOverridableArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) NonOverridableArgumentsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"nonOverridableArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) NotificationProperty() AwsJob_NotificationPropertyPropertyOutputReference {
	var returns AwsJob_NotificationPropertyPropertyOutputReference
	_jsii_.Get(
		j,
		"notificationProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) NotificationPropertyInput() *AwsJob_NotificationPropertyProperty {
	var returns *AwsJob_NotificationPropertyProperty
	_jsii_.Get(
		j,
		"notificationPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) NumberOfWorkers() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) NumberOfWorkersInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfWorkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) SecurityConfiguration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) SecurityConfigurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) SourceControlDetails() AwsJob_SourceControlDetailsPropertyOutputReference {
	var returns AwsJob_SourceControlDetailsPropertyOutputReference
	_jsii_.Get(
		j,
		"sourceControlDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) SourceControlDetailsInput() *AwsJob_SourceControlDetailsProperty {
	var returns *AwsJob_SourceControlDetailsProperty
	_jsii_.Get(
		j,
		"sourceControlDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) Timeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) TimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) WorkerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJob) WorkerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workerTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job aws_glue_job} Resource.
// Experimental.
func NewAwsJob(scope constructs.Construct, id *string, config *AwsJobConfig) AwsJob {
	_init_.Initialize()

	if err := validateNewAwsJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsJob{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_job aws_glue_job} Resource.
// Experimental.
func NewAwsJob_Override(a AwsJob, scope constructs.Construct, id *string, config *AwsJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsJob",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetConnections(val *[]*string) {
	if err := j.validateSetConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connections",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetDefaultArguments(val *map[string]*string) {
	if err := j.validateSetDefaultArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultArguments",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetExecutionClass(val *string) {
	if err := j.validateSetExecutionClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionClass",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetGlueVersion(val *string) {
	if err := j.validateSetGlueVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"glueVersion",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetJobMode(val *string) {
	if err := j.validateSetJobModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobMode",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetJobRunQueuingEnabled(val interface{}) {
	if err := j.validateSetJobRunQueuingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobRunQueuingEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetMaintenanceWindow(val *string) {
	if err := j.validateSetMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetMaxCapacity(val *float64) {
	if err := j.validateSetMaxCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetNonOverridableArguments(val *map[string]*string) {
	if err := j.validateSetNonOverridableArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nonOverridableArguments",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetNumberOfWorkers(val *float64) {
	if err := j.validateSetNumberOfWorkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfWorkers",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetSecurityConfiguration(val *string) {
	if err := j.validateSetSecurityConfigurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityConfiguration",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetTimeout(val *float64) {
	if err := j.validateSetTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeout",
		val,
	)
}

func (j *jsiiProxy_AwsJob)SetWorkerType(val *string) {
	if err := j.validateSetWorkerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workerType",
		val,
	)
}

// Generates CDKTN code for importing a AwsJob resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-glue.AwsJob",
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
func AwsJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-glue.AwsJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-glue.AwsJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-glue.AwsJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-glue.AwsJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsJob) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsJob) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsJob) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsJob) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsJob) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJob) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsJob) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsJob) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsJob) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsJob) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsJob) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsJob) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsJob) PutCommand(value *AwsJob_CommandProperty) {
	if err := a.validatePutCommandParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCommand",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsJob) PutExecutionProperty(value *AwsJob_ExecutionPropertyProperty) {
	if err := a.validatePutExecutionPropertyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExecutionProperty",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsJob) PutNotificationProperty(value *AwsJob_NotificationPropertyProperty) {
	if err := a.validatePutNotificationPropertyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNotificationProperty",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsJob) PutSourceControlDetails(value *AwsJob_SourceControlDetailsProperty) {
	if err := a.validatePutSourceControlDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourceControlDetails",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsJob) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsJob) ResetConnections() {
	_jsii_.InvokeVoid(
		a,
		"resetConnections",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetDefaultArguments() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultArguments",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetExecutionClass() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionClass",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetExecutionProperty() {
	_jsii_.InvokeVoid(
		a,
		"resetExecutionProperty",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetGlueVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetGlueVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetJobMode() {
	_jsii_.InvokeVoid(
		a,
		"resetJobMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetJobRunQueuingEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetJobRunQueuingEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetMaxCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetNonOverridableArguments() {
	_jsii_.InvokeVoid(
		a,
		"resetNonOverridableArguments",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetNotificationProperty() {
	_jsii_.InvokeVoid(
		a,
		"resetNotificationProperty",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetNumberOfWorkers() {
	_jsii_.InvokeVoid(
		a,
		"resetNumberOfWorkers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetSecurityConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetSourceControlDetails() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceControlDetails",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) ResetWorkerType() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkerType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJob) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

