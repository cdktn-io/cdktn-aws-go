package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsnetworkmanager/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsnetworkmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document aws_networkmanager_core_network_policy_document}.
// Experimental.
type DataTfCoreNetworkPolicyDocument interface {
	cdktn.TerraformDataSource
	// Experimental.
	AttachmentPolicies() DataTfCoreNetworkPolicyDocument_AttachmentPoliciesPropertyList
	// Experimental.
	AttachmentPoliciesInput() interface{}
	// Experimental.
	AttachmentRoutingPolicyRules() DataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyList
	// Experimental.
	AttachmentRoutingPolicyRulesInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	CoreNetworkConfiguration() DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyList
	// Experimental.
	CoreNetworkConfigurationInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	Json() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	NetworkFunctionGroups() DataTfCoreNetworkPolicyDocument_NetworkFunctionGroupsPropertyList
	// Experimental.
	NetworkFunctionGroupsInput() interface{}
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	RoutingPolicies() DataTfCoreNetworkPolicyDocument_RoutingPoliciesPropertyList
	// Experimental.
	RoutingPoliciesInput() interface{}
	// Experimental.
	SegmentActions() DataTfCoreNetworkPolicyDocument_SegmentActionsPropertyList
	// Experimental.
	SegmentActionsInput() interface{}
	// Experimental.
	Segments() DataTfCoreNetworkPolicyDocument_SegmentsPropertyList
	// Experimental.
	SegmentsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Version() *string
	// Experimental.
	SetVersion(val *string)
	// Experimental.
	VersionInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Experimental.
	PutAttachmentPolicies(value interface{})
	// Experimental.
	PutAttachmentRoutingPolicyRules(value interface{})
	// Experimental.
	PutCoreNetworkConfiguration(value interface{})
	// Experimental.
	PutNetworkFunctionGroups(value interface{})
	// Experimental.
	PutRoutingPolicies(value interface{})
	// Experimental.
	PutSegmentActions(value interface{})
	// Experimental.
	PutSegments(value interface{})
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
	ResetAttachmentPolicies()
	// Experimental.
	ResetAttachmentRoutingPolicyRules()
	// Experimental.
	ResetId()
	// Experimental.
	ResetNetworkFunctionGroups()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRoutingPolicies()
	// Experimental.
	ResetSegmentActions()
	// Experimental.
	ResetVersion()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataTfCoreNetworkPolicyDocument
type jsiiProxy_DataTfCoreNetworkPolicyDocument struct {
	internal.Type__cdktnTerraformDataSource
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) AttachmentPolicies() DataTfCoreNetworkPolicyDocument_AttachmentPoliciesPropertyList {
	var returns DataTfCoreNetworkPolicyDocument_AttachmentPoliciesPropertyList
	_jsii_.Get(
		j,
		"attachmentPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) AttachmentPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachmentPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) AttachmentRoutingPolicyRules() DataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyList {
	var returns DataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesPropertyList
	_jsii_.Get(
		j,
		"attachmentRoutingPolicyRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) AttachmentRoutingPolicyRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attachmentRoutingPolicyRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) CoreNetworkConfiguration() DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyList {
	var returns DataTfCoreNetworkPolicyDocument_CoreNetworkConfigurationPropertyList
	_jsii_.Get(
		j,
		"coreNetworkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) CoreNetworkConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"coreNetworkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) Json() *string {
	var returns *string
	_jsii_.Get(
		j,
		"json",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) NetworkFunctionGroups() DataTfCoreNetworkPolicyDocument_NetworkFunctionGroupsPropertyList {
	var returns DataTfCoreNetworkPolicyDocument_NetworkFunctionGroupsPropertyList
	_jsii_.Get(
		j,
		"networkFunctionGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) NetworkFunctionGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkFunctionGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) RoutingPolicies() DataTfCoreNetworkPolicyDocument_RoutingPoliciesPropertyList {
	var returns DataTfCoreNetworkPolicyDocument_RoutingPoliciesPropertyList
	_jsii_.Get(
		j,
		"routingPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) RoutingPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) SegmentActions() DataTfCoreNetworkPolicyDocument_SegmentActionsPropertyList {
	var returns DataTfCoreNetworkPolicyDocument_SegmentActionsPropertyList
	_jsii_.Get(
		j,
		"segmentActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) SegmentActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"segmentActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) Segments() DataTfCoreNetworkPolicyDocument_SegmentsPropertyList {
	var returns DataTfCoreNetworkPolicyDocument_SegmentsPropertyList
	_jsii_.Get(
		j,
		"segments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) SegmentsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"segmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document aws_networkmanager_core_network_policy_document} Data Source.
// Experimental.
func NewDataTfCoreNetworkPolicyDocument(scope constructs.Construct, id *string, config *DataTfCoreNetworkPolicyDocumentConfig) DataTfCoreNetworkPolicyDocument {
	_init_.Initialize()

	if err := validateNewDataTfCoreNetworkPolicyDocumentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfCoreNetworkPolicyDocument{}

	_jsii_.Create(
		"@cdktn/aws-network-manager.DataTfCoreNetworkPolicyDocument",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/networkmanager_core_network_policy_document aws_networkmanager_core_network_policy_document} Data Source.
// Experimental.
func NewDataTfCoreNetworkPolicyDocument_Override(d DataTfCoreNetworkPolicyDocument, scope constructs.Construct, id *string, config *DataTfCoreNetworkPolicyDocumentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-manager.DataTfCoreNetworkPolicyDocument",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTN code for importing a DataTfCoreNetworkPolicyDocument resource upon running "cdktn plan <stack-name>".
// Experimental.
func DataTfCoreNetworkPolicyDocument_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateDataTfCoreNetworkPolicyDocument_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-network-manager.DataTfCoreNetworkPolicyDocument",
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
func DataTfCoreNetworkPolicyDocument_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfCoreNetworkPolicyDocument_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-network-manager.DataTfCoreNetworkPolicyDocument",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfCoreNetworkPolicyDocument_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfCoreNetworkPolicyDocument_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-network-manager.DataTfCoreNetworkPolicyDocument",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataTfCoreNetworkPolicyDocument_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataTfCoreNetworkPolicyDocument_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-network-manager.DataTfCoreNetworkPolicyDocument",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataTfCoreNetworkPolicyDocument_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-network-manager.DataTfCoreNetworkPolicyDocument",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) PutAttachmentPolicies(value interface{}) {
	if err := d.validatePutAttachmentPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAttachmentPolicies",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) PutAttachmentRoutingPolicyRules(value interface{}) {
	if err := d.validatePutAttachmentRoutingPolicyRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAttachmentRoutingPolicyRules",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) PutCoreNetworkConfiguration(value interface{}) {
	if err := d.validatePutCoreNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCoreNetworkConfiguration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) PutNetworkFunctionGroups(value interface{}) {
	if err := d.validatePutNetworkFunctionGroupsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNetworkFunctionGroups",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) PutRoutingPolicies(value interface{}) {
	if err := d.validatePutRoutingPoliciesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRoutingPolicies",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) PutSegmentActions(value interface{}) {
	if err := d.validatePutSegmentActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSegmentActions",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) PutSegments(value interface{}) {
	if err := d.validatePutSegmentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSegments",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := d.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) ResetAttachmentPolicies() {
	_jsii_.InvokeVoid(
		d,
		"resetAttachmentPolicies",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) ResetAttachmentRoutingPolicyRules() {
	_jsii_.InvokeVoid(
		d,
		"resetAttachmentRoutingPolicyRules",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) ResetNetworkFunctionGroups() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkFunctionGroups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) ResetRoutingPolicies() {
	_jsii_.InvokeVoid(
		d,
		"resetRoutingPolicies",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) ResetSegmentActions() {
	_jsii_.InvokeVoid(
		d,
		"resetSegmentActions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		d,
		"with",
		args,
		&returns,
	)

	return returns
}

