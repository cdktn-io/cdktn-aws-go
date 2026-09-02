package awsconnectcustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconnectcustomerprofiles/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsconnectcustomerprofiles/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_RuleBasedMatchingPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AttributeTypesSelector() TfDomain_AttributeTypesSelectorPropertyOutputReference
	// Experimental.
	AttributeTypesSelectorInput() *TfDomain_AttributeTypesSelectorProperty
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
	// Experimental.
	ConflictResolution() TfDomain_RuleBasedMatchingConflictResolutionPropertyOutputReference
	// Experimental.
	ConflictResolutionInput() *TfDomain_RuleBasedMatchingConflictResolutionProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	ExportingConfig() TfDomain_RuleBasedMatchingExportingConfigPropertyOutputReference
	// Experimental.
	ExportingConfigInput() *TfDomain_RuleBasedMatchingExportingConfigProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDomain_RuleBasedMatchingProperty
	// Experimental.
	SetInternalValue(val *TfDomain_RuleBasedMatchingProperty)
	// Experimental.
	MatchingRules() TfDomain_MatchingRulesPropertyList
	// Experimental.
	MatchingRulesInput() interface{}
	// Experimental.
	MaxAllowedRuleLevelForMatching() *float64
	// Experimental.
	SetMaxAllowedRuleLevelForMatching(val *float64)
	// Experimental.
	MaxAllowedRuleLevelForMatchingInput() *float64
	// Experimental.
	MaxAllowedRuleLevelForMerging() *float64
	// Experimental.
	SetMaxAllowedRuleLevelForMerging(val *float64)
	// Experimental.
	MaxAllowedRuleLevelForMergingInput() *float64
	// Experimental.
	Status() *string
	// Experimental.
	SetStatus(val *string)
	// Experimental.
	StatusInput() *string
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
	PutAttributeTypesSelector(value *TfDomain_AttributeTypesSelectorProperty)
	// Experimental.
	PutConflictResolution(value *TfDomain_RuleBasedMatchingConflictResolutionProperty)
	// Experimental.
	PutExportingConfig(value *TfDomain_RuleBasedMatchingExportingConfigProperty)
	// Experimental.
	PutMatchingRules(value interface{})
	// Experimental.
	ResetAttributeTypesSelector()
	// Experimental.
	ResetConflictResolution()
	// Experimental.
	ResetExportingConfig()
	// Experimental.
	ResetMatchingRules()
	// Experimental.
	ResetMaxAllowedRuleLevelForMatching()
	// Experimental.
	ResetMaxAllowedRuleLevelForMerging()
	// Experimental.
	ResetStatus()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDomain_RuleBasedMatchingPropertyOutputReference
type jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) AttributeTypesSelector() TfDomain_AttributeTypesSelectorPropertyOutputReference {
	var returns TfDomain_AttributeTypesSelectorPropertyOutputReference
	_jsii_.Get(
		j,
		"attributeTypesSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) AttributeTypesSelectorInput() *TfDomain_AttributeTypesSelectorProperty {
	var returns *TfDomain_AttributeTypesSelectorProperty
	_jsii_.Get(
		j,
		"attributeTypesSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) ConflictResolution() TfDomain_RuleBasedMatchingConflictResolutionPropertyOutputReference {
	var returns TfDomain_RuleBasedMatchingConflictResolutionPropertyOutputReference
	_jsii_.Get(
		j,
		"conflictResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) ConflictResolutionInput() *TfDomain_RuleBasedMatchingConflictResolutionProperty {
	var returns *TfDomain_RuleBasedMatchingConflictResolutionProperty
	_jsii_.Get(
		j,
		"conflictResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) ExportingConfig() TfDomain_RuleBasedMatchingExportingConfigPropertyOutputReference {
	var returns TfDomain_RuleBasedMatchingExportingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"exportingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) ExportingConfigInput() *TfDomain_RuleBasedMatchingExportingConfigProperty {
	var returns *TfDomain_RuleBasedMatchingExportingConfigProperty
	_jsii_.Get(
		j,
		"exportingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) InternalValue() *TfDomain_RuleBasedMatchingProperty {
	var returns *TfDomain_RuleBasedMatchingProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) MatchingRules() TfDomain_MatchingRulesPropertyList {
	var returns TfDomain_MatchingRulesPropertyList
	_jsii_.Get(
		j,
		"matchingRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) MatchingRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchingRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) MaxAllowedRuleLevelForMatching() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllowedRuleLevelForMatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) MaxAllowedRuleLevelForMatchingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllowedRuleLevelForMatchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) MaxAllowedRuleLevelForMerging() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllowedRuleLevelForMerging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) MaxAllowedRuleLevelForMergingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllowedRuleLevelForMergingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_RuleBasedMatchingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDomain_RuleBasedMatchingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_RuleBasedMatchingPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.TfDomain.RuleBasedMatchingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_RuleBasedMatchingPropertyOutputReference_Override(t TfDomain_RuleBasedMatchingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.TfDomain.RuleBasedMatchingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference)SetInternalValue(val *TfDomain_RuleBasedMatchingProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference)SetMaxAllowedRuleLevelForMatching(val *float64) {
	if err := j.validateSetMaxAllowedRuleLevelForMatchingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAllowedRuleLevelForMatching",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference)SetMaxAllowedRuleLevelForMerging(val *float64) {
	if err := j.validateSetMaxAllowedRuleLevelForMergingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAllowedRuleLevelForMerging",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) PutAttributeTypesSelector(value *TfDomain_AttributeTypesSelectorProperty) {
	if err := t.validatePutAttributeTypesSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAttributeTypesSelector",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) PutConflictResolution(value *TfDomain_RuleBasedMatchingConflictResolutionProperty) {
	if err := t.validatePutConflictResolutionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConflictResolution",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) PutExportingConfig(value *TfDomain_RuleBasedMatchingExportingConfigProperty) {
	if err := t.validatePutExportingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExportingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) PutMatchingRules(value interface{}) {
	if err := t.validatePutMatchingRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMatchingRules",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) ResetAttributeTypesSelector() {
	_jsii_.InvokeVoid(
		t,
		"resetAttributeTypesSelector",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) ResetConflictResolution() {
	_jsii_.InvokeVoid(
		t,
		"resetConflictResolution",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) ResetExportingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetExportingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) ResetMatchingRules() {
	_jsii_.InvokeVoid(
		t,
		"resetMatchingRules",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) ResetMaxAllowedRuleLevelForMatching() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxAllowedRuleLevelForMatching",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) ResetMaxAllowedRuleLevelForMerging() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxAllowedRuleLevelForMerging",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		t,
		"resetStatus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_RuleBasedMatchingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

