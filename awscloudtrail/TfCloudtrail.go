package awscloudtrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudtrail/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awscloudtrail/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail aws_cloudtrail}.
// Experimental.
type TfCloudtrail interface {
	cdktn.TerraformResource
	// Experimental.
	AdvancedEventSelector() TfCloudtrail_AdvancedEventSelectorPropertyList
	// Experimental.
	AdvancedEventSelectorInput() interface{}
	// Experimental.
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CloudWatchLogsGroupArn() *string
	// Experimental.
	SetCloudWatchLogsGroupArn(val *string)
	// Experimental.
	CloudWatchLogsGroupArnInput() *string
	// Experimental.
	CloudWatchLogsRoleArn() *string
	// Experimental.
	SetCloudWatchLogsRoleArn(val *string)
	// Experimental.
	CloudWatchLogsRoleArnInput() *string
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
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EnableLogFileValidation() interface{}
	// Experimental.
	SetEnableLogFileValidation(val interface{})
	// Experimental.
	EnableLogFileValidationInput() interface{}
	// Experimental.
	EnableLogging() interface{}
	// Experimental.
	SetEnableLogging(val interface{})
	// Experimental.
	EnableLoggingInput() interface{}
	// Experimental.
	EventSelector() TfCloudtrail_EventSelectorPropertyList
	// Experimental.
	EventSelectorInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HomeRegion() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	IncludeGlobalServiceEvents() interface{}
	// Experimental.
	SetIncludeGlobalServiceEvents(val interface{})
	// Experimental.
	IncludeGlobalServiceEventsInput() interface{}
	// Experimental.
	InsightSelector() TfCloudtrail_InsightSelectorPropertyList
	// Experimental.
	InsightSelectorInput() interface{}
	// Experimental.
	IsMultiRegionTrail() interface{}
	// Experimental.
	SetIsMultiRegionTrail(val interface{})
	// Experimental.
	IsMultiRegionTrailInput() interface{}
	// Experimental.
	IsOrganizationTrail() interface{}
	// Experimental.
	SetIsOrganizationTrail(val interface{})
	// Experimental.
	IsOrganizationTrailInput() interface{}
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
	S3BucketName() *string
	// Experimental.
	SetS3BucketName(val *string)
	// Experimental.
	S3BucketNameInput() *string
	// Experimental.
	S3KeyPrefix() *string
	// Experimental.
	SetS3KeyPrefix(val *string)
	// Experimental.
	S3KeyPrefixInput() *string
	// Experimental.
	SnsTopicArn() *string
	// Experimental.
	SnsTopicName() *string
	// Experimental.
	SetSnsTopicName(val *string)
	// Experimental.
	SnsTopicNameInput() *string
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
	PutAdvancedEventSelector(value interface{})
	// Experimental.
	PutEventSelector(value interface{})
	// Experimental.
	PutInsightSelector(value interface{})
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
	ResetAdvancedEventSelector()
	// Experimental.
	ResetCloudWatchLogsGroupArn()
	// Experimental.
	ResetCloudWatchLogsRoleArn()
	// Experimental.
	ResetEnableLogFileValidation()
	// Experimental.
	ResetEnableLogging()
	// Experimental.
	ResetEventSelector()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIncludeGlobalServiceEvents()
	// Experimental.
	ResetInsightSelector()
	// Experimental.
	ResetIsMultiRegionTrail()
	// Experimental.
	ResetIsOrganizationTrail()
	// Experimental.
	ResetKmsKeyId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetS3KeyPrefix()
	// Experimental.
	ResetSnsTopicName()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
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

// The jsii proxy struct for TfCloudtrail
type jsiiProxy_TfCloudtrail struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfCloudtrail) AdvancedEventSelector() TfCloudtrail_AdvancedEventSelectorPropertyList {
	var returns TfCloudtrail_AdvancedEventSelectorPropertyList
	_jsii_.Get(
		j,
		"advancedEventSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) AdvancedEventSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedEventSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) CloudWatchLogsGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogsGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) CloudWatchLogsGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogsGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) CloudWatchLogsRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogsRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) CloudWatchLogsRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudWatchLogsRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) EnableLogFileValidation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogFileValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) EnableLogFileValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogFileValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) EnableLogging() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLogging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) EnableLoggingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLoggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) EventSelector() TfCloudtrail_EventSelectorPropertyList {
	var returns TfCloudtrail_EventSelectorPropertyList
	_jsii_.Get(
		j,
		"eventSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) EventSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eventSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) HomeRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homeRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) IncludeGlobalServiceEvents() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalServiceEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) IncludeGlobalServiceEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeGlobalServiceEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) InsightSelector() TfCloudtrail_InsightSelectorPropertyList {
	var returns TfCloudtrail_InsightSelectorPropertyList
	_jsii_.Get(
		j,
		"insightSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) InsightSelectorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insightSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) IsMultiRegionTrail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMultiRegionTrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) IsMultiRegionTrailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isMultiRegionTrailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) IsOrganizationTrail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isOrganizationTrail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) IsOrganizationTrailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isOrganizationTrailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) S3BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) S3KeyPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) S3KeyPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) SnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) SnsTopicName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) SnsTopicNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudtrail) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail aws_cloudtrail} Resource.
// Experimental.
func NewTfCloudtrail(scope constructs.Construct, id *string, config *TfCloudtrailConfig) TfCloudtrail {
	_init_.Initialize()

	if err := validateNewTfCloudtrailParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCloudtrail{}

	_jsii_.Create(
		"@cdktn/aws-cloudtrail.TfCloudtrail",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail aws_cloudtrail} Resource.
// Experimental.
func NewTfCloudtrail_Override(t TfCloudtrail, scope constructs.Construct, id *string, config *TfCloudtrailConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudtrail.TfCloudtrail",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetCloudWatchLogsGroupArn(val *string) {
	if err := j.validateSetCloudWatchLogsGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudWatchLogsGroupArn",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetCloudWatchLogsRoleArn(val *string) {
	if err := j.validateSetCloudWatchLogsRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudWatchLogsRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetEnableLogFileValidation(val interface{}) {
	if err := j.validateSetEnableLogFileValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLogFileValidation",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetEnableLogging(val interface{}) {
	if err := j.validateSetEnableLoggingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLogging",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetIncludeGlobalServiceEvents(val interface{}) {
	if err := j.validateSetIncludeGlobalServiceEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeGlobalServiceEvents",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetIsMultiRegionTrail(val interface{}) {
	if err := j.validateSetIsMultiRegionTrailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isMultiRegionTrail",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetIsOrganizationTrail(val interface{}) {
	if err := j.validateSetIsOrganizationTrailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isOrganizationTrail",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetS3BucketName(val *string) {
	if err := j.validateSetS3BucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetS3KeyPrefix(val *string) {
	if err := j.validateSetS3KeyPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3KeyPrefix",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetSnsTopicName(val *string) {
	if err := j.validateSetSnsTopicNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snsTopicName",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfCloudtrail)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a TfCloudtrail resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfCloudtrail_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfCloudtrail_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudtrail.TfCloudtrail",
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
func TfCloudtrail_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCloudtrail_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudtrail.TfCloudtrail",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCloudtrail_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCloudtrail_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudtrail.TfCloudtrail",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfCloudtrail_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfCloudtrail_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-cloudtrail.TfCloudtrail",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfCloudtrail_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-cloudtrail.TfCloudtrail",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfCloudtrail) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfCloudtrail) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfCloudtrail) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCloudtrail) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCloudtrail) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCloudtrail) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCloudtrail) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCloudtrail) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCloudtrail) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCloudtrail) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCloudtrail) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCloudtrail) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudtrail) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfCloudtrail) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCloudtrail) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfCloudtrail) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCloudtrail) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfCloudtrail) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfCloudtrail) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfCloudtrail) PutAdvancedEventSelector(value interface{}) {
	if err := t.validatePutAdvancedEventSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAdvancedEventSelector",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCloudtrail) PutEventSelector(value interface{}) {
	if err := t.validatePutEventSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEventSelector",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCloudtrail) PutInsightSelector(value interface{}) {
	if err := t.validatePutInsightSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInsightSelector",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCloudtrail) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetAdvancedEventSelector() {
	_jsii_.InvokeVoid(
		t,
		"resetAdvancedEventSelector",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetCloudWatchLogsGroupArn() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudWatchLogsGroupArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetCloudWatchLogsRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudWatchLogsRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetEnableLogFileValidation() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableLogFileValidation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetEnableLogging() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableLogging",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetEventSelector() {
	_jsii_.InvokeVoid(
		t,
		"resetEventSelector",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetIncludeGlobalServiceEvents() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeGlobalServiceEvents",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetInsightSelector() {
	_jsii_.InvokeVoid(
		t,
		"resetInsightSelector",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetIsMultiRegionTrail() {
	_jsii_.InvokeVoid(
		t,
		"resetIsMultiRegionTrail",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetIsOrganizationTrail() {
	_jsii_.InvokeVoid(
		t,
		"resetIsOrganizationTrail",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetS3KeyPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetS3KeyPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetSnsTopicName() {
	_jsii_.InvokeVoid(
		t,
		"resetSnsTopicName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudtrail) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudtrail) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudtrail) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudtrail) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudtrail) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudtrail) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudtrail) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

