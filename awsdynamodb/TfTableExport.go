package awsdynamodb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdynamodb/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsdynamodb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export aws_dynamodb_table_export}.
// Experimental.
type TfTableExport interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	BilledSizeInBytes() *float64
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
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EndTime() *string
	// Experimental.
	ExportFormat() *string
	// Experimental.
	SetExportFormat(val *string)
	// Experimental.
	ExportFormatInput() *string
	// Experimental.
	ExportStatus() *string
	// Experimental.
	ExportTime() *string
	// Experimental.
	SetExportTime(val *string)
	// Experimental.
	ExportTimeInput() *string
	// Experimental.
	ExportType() *string
	// Experimental.
	SetExportType(val *string)
	// Experimental.
	ExportTypeInput() *string
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
	IncrementalExportSpecification() TfTableExport_IncrementalExportSpecificationPropertyOutputReference
	// Experimental.
	IncrementalExportSpecificationInput() *TfTableExport_IncrementalExportSpecificationProperty
	// Experimental.
	ItemCount() *float64
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	ManifestFilesS3Key() *string
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
	S3Bucket() *string
	// Experimental.
	SetS3Bucket(val *string)
	// Experimental.
	S3BucketInput() *string
	// Experimental.
	S3BucketOwner() *string
	// Experimental.
	SetS3BucketOwner(val *string)
	// Experimental.
	S3BucketOwnerInput() *string
	// Experimental.
	S3Prefix() *string
	// Experimental.
	SetS3Prefix(val *string)
	// Experimental.
	S3PrefixInput() *string
	// Experimental.
	S3SseAlgorithm() *string
	// Experimental.
	SetS3SseAlgorithm(val *string)
	// Experimental.
	S3SseAlgorithmInput() *string
	// Experimental.
	S3SseKmsKeyId() *string
	// Experimental.
	SetS3SseKmsKeyId(val *string)
	// Experimental.
	S3SseKmsKeyIdInput() *string
	// Experimental.
	StartTime() *string
	// Experimental.
	TableArn() *string
	// Experimental.
	SetTableArn(val *string)
	// Experimental.
	TableArnInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() TfTableExport_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
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
	PutIncrementalExportSpecification(value *TfTableExport_IncrementalExportSpecificationProperty)
	// Experimental.
	PutTimeouts(value *TfTableExport_TimeoutsProperty)
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
	ResetExportFormat()
	// Experimental.
	ResetExportTime()
	// Experimental.
	ResetExportType()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIncrementalExportSpecification()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetS3BucketOwner()
	// Experimental.
	ResetS3Prefix()
	// Experimental.
	ResetS3SseAlgorithm()
	// Experimental.
	ResetS3SseKmsKeyId()
	// Experimental.
	ResetTimeouts()
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

// The jsii proxy struct for TfTableExport
type jsiiProxy_TfTableExport struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfTableExport) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) BilledSizeInBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"billedSizeInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) EndTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) ExportFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) ExportFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) ExportStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) ExportTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) ExportTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) ExportType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) ExportTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) IncrementalExportSpecification() TfTableExport_IncrementalExportSpecificationPropertyOutputReference {
	var returns TfTableExport_IncrementalExportSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"incrementalExportSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) IncrementalExportSpecificationInput() *TfTableExport_IncrementalExportSpecificationProperty {
	var returns *TfTableExport_IncrementalExportSpecificationProperty
	_jsii_.Get(
		j,
		"incrementalExportSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) ItemCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"itemCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) ManifestFilesS3Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"manifestFilesS3Key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) S3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) S3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) S3BucketOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) S3BucketOwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketOwnerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) S3Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) S3PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3PrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) S3SseAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3SseAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) S3SseAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3SseAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) S3SseKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3SseKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) S3SseKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3SseKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) StartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) TableArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) TableArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) Timeouts() TfTableExport_TimeoutsPropertyOutputReference {
	var returns TfTableExport_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTableExport) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export aws_dynamodb_table_export} Resource.
// Experimental.
func NewTfTableExport(scope constructs.Construct, id *string, config *TfTableExportConfig) TfTableExport {
	_init_.Initialize()

	if err := validateNewTfTableExportParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTableExport{}

	_jsii_.Create(
		"@cdktn/aws-dynamodb.TfTableExport",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table_export aws_dynamodb_table_export} Resource.
// Experimental.
func NewTfTableExport_Override(t TfTableExport, scope constructs.Construct, id *string, config *TfTableExportConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dynamodb.TfTableExport",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfTableExport)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfTableExport)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfTableExport)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfTableExport)SetExportFormat(val *string) {
	if err := j.validateSetExportFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportFormat",
		val,
	)
}

func (j *jsiiProxy_TfTableExport)SetExportTime(val *string) {
	if err := j.validateSetExportTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportTime",
		val,
	)
}

func (j *jsiiProxy_TfTableExport)SetExportType(val *string) {
	if err := j.validateSetExportTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportType",
		val,
	)
}

func (j *jsiiProxy_TfTableExport)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfTableExport)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfTableExport)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfTableExport)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfTableExport)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfTableExport)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfTableExport)SetS3Bucket(val *string) {
	if err := j.validateSetS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Bucket",
		val,
	)
}

func (j *jsiiProxy_TfTableExport)SetS3BucketOwner(val *string) {
	if err := j.validateSetS3BucketOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketOwner",
		val,
	)
}

func (j *jsiiProxy_TfTableExport)SetS3Prefix(val *string) {
	if err := j.validateSetS3PrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Prefix",
		val,
	)
}

func (j *jsiiProxy_TfTableExport)SetS3SseAlgorithm(val *string) {
	if err := j.validateSetS3SseAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3SseAlgorithm",
		val,
	)
}

func (j *jsiiProxy_TfTableExport)SetS3SseKmsKeyId(val *string) {
	if err := j.validateSetS3SseKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3SseKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfTableExport)SetTableArn(val *string) {
	if err := j.validateSetTableArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableArn",
		val,
	)
}

// Generates CDKTN code for importing a TfTableExport resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfTableExport_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfTableExport_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-dynamodb.TfTableExport",
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
func TfTableExport_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTableExport_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dynamodb.TfTableExport",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfTableExport_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTableExport_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dynamodb.TfTableExport",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfTableExport_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTableExport_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-dynamodb.TfTableExport",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfTableExport_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-dynamodb.TfTableExport",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfTableExport) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfTableExport) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfTableExport) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTableExport) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTableExport) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTableExport) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTableExport) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTableExport) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTableExport) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTableExport) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTableExport) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTableExport) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTableExport) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfTableExport) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTableExport) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfTableExport) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfTableExport) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfTableExport) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfTableExport) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfTableExport) PutIncrementalExportSpecification(value *TfTableExport_IncrementalExportSpecificationProperty) {
	if err := t.validatePutIncrementalExportSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIncrementalExportSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTableExport) PutTimeouts(value *TfTableExport_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTableExport) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfTableExport) ResetExportFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetExportFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTableExport) ResetExportTime() {
	_jsii_.InvokeVoid(
		t,
		"resetExportTime",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTableExport) ResetExportType() {
	_jsii_.InvokeVoid(
		t,
		"resetExportType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTableExport) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTableExport) ResetIncrementalExportSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetIncrementalExportSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTableExport) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTableExport) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTableExport) ResetS3BucketOwner() {
	_jsii_.InvokeVoid(
		t,
		"resetS3BucketOwner",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTableExport) ResetS3Prefix() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Prefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTableExport) ResetS3SseAlgorithm() {
	_jsii_.InvokeVoid(
		t,
		"resetS3SseAlgorithm",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTableExport) ResetS3SseKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetS3SseKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTableExport) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTableExport) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTableExport) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTableExport) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTableExport) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTableExport) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTableExport) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTableExport) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

