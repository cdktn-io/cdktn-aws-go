package awscloudwatchobservabilityadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudwatchobservabilityadmin/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudwatchobservabilityadmin/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllowFieldUpdates() interface{}
	// Experimental.
	SetAllowFieldUpdates(val interface{})
	// Experimental.
	AllowFieldUpdatesInput() interface{}
	// Experimental.
	AllRegions() interface{}
	// Experimental.
	SetAllRegions(val interface{})
	// Experimental.
	AllRegionsInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DestinationConfiguration() AwsObservabilityadminTelemetryRuleForOrganization_DestinationConfigurationPropertyList
	// Experimental.
	DestinationConfigurationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Regions() *[]*string
	// Experimental.
	SetRegions(val *[]*string)
	// Experimental.
	RegionsInput() *[]*string
	// Experimental.
	ResourceType() *string
	// Experimental.
	SetResourceType(val *string)
	// Experimental.
	ResourceTypeInput() *string
	// Experimental.
	Scope() *string
	// Experimental.
	SetScope(val *string)
	// Experimental.
	ScopeInput() *string
	// Experimental.
	SelectionCriteria() *string
	// Experimental.
	SetSelectionCriteria(val *string)
	// Experimental.
	SelectionCriteriaInput() *string
	// Experimental.
	TelemetrySourceTypes() *[]*string
	// Experimental.
	SetTelemetrySourceTypes(val *[]*string)
	// Experimental.
	TelemetrySourceTypesInput() *[]*string
	// Experimental.
	TelemetryType() *string
	// Experimental.
	SetTelemetryType(val *string)
	// Experimental.
	TelemetryTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	PutDestinationConfiguration(value interface{})
	// Experimental.
	ResetAllowFieldUpdates()
	// Experimental.
	ResetAllRegions()
	// Experimental.
	ResetDestinationConfiguration()
	// Experimental.
	ResetRegions()
	// Experimental.
	ResetResourceType()
	// Experimental.
	ResetScope()
	// Experimental.
	ResetSelectionCriteria()
	// Experimental.
	ResetTelemetrySourceTypes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference
type jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) AllowFieldUpdates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowFieldUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) AllowFieldUpdatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowFieldUpdatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) AllRegions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) AllRegionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) DestinationConfiguration() AwsObservabilityadminTelemetryRuleForOrganization_DestinationConfigurationPropertyList {
	var returns AwsObservabilityadminTelemetryRuleForOrganization_DestinationConfigurationPropertyList
	_jsii_.Get(
		j,
		"destinationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) DestinationConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"destinationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) SelectionCriteria() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectionCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) SelectionCriteriaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectionCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) TelemetrySourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"telemetrySourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) TelemetrySourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"telemetrySourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) TelemetryType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"telemetryType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) TelemetryTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"telemetryTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-observability-admin.AwsObservabilityadminTelemetryRuleForOrganization.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference_Override(a AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-observability-admin.AwsObservabilityadminTelemetryRuleForOrganization.RulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference)SetAllowFieldUpdates(val interface{}) {
	if err := j.validateSetAllowFieldUpdatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowFieldUpdates",
		val,
	)
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference)SetAllRegions(val interface{}) {
	if err := j.validateSetAllRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allRegions",
		val,
	)
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference)SetSelectionCriteria(val *string) {
	if err := j.validateSetSelectionCriteriaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selectionCriteria",
		val,
	)
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference)SetTelemetrySourceTypes(val *[]*string) {
	if err := j.validateSetTelemetrySourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"telemetrySourceTypes",
		val,
	)
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference)SetTelemetryType(val *string) {
	if err := j.validateSetTelemetryTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"telemetryType",
		val,
	)
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) PutDestinationConfiguration(value interface{}) {
	if err := a.validatePutDestinationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDestinationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) ResetAllowFieldUpdates() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowFieldUpdates",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) ResetAllRegions() {
	_jsii_.InvokeVoid(
		a,
		"resetAllRegions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) ResetDestinationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) ResetRegions() {
	_jsii_.InvokeVoid(
		a,
		"resetRegions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		a,
		"resetScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) ResetSelectionCriteria() {
	_jsii_.InvokeVoid(
		a,
		"resetSelectionCriteria",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) ResetTelemetrySourceTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetTelemetrySourceTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsObservabilityadminTelemetryRuleForOrganization_RulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

