package connectcustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/connectcustomerprofiles/jsii"

	"github.com/cdktn-io/cdktn-aws-go/connectcustomerprofiles/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDomain_RuleBasedMatchingPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AttributeTypesSelector() AwsDomain_AttributeTypesSelectorPropertyOutputReference
	// Experimental.
	AttributeTypesSelectorInput() *AwsDomain_AttributeTypesSelectorProperty
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
	ConflictResolution() AwsDomain_RuleBasedMatchingConflictResolutionPropertyOutputReference
	// Experimental.
	ConflictResolutionInput() *AwsDomain_RuleBasedMatchingConflictResolutionProperty
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
	ExportingConfig() AwsDomain_RuleBasedMatchingExportingConfigPropertyOutputReference
	// Experimental.
	ExportingConfigInput() *AwsDomain_RuleBasedMatchingExportingConfigProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDomain_RuleBasedMatchingProperty
	// Experimental.
	SetInternalValue(val *AwsDomain_RuleBasedMatchingProperty)
	// Experimental.
	MatchingRules() AwsDomain_MatchingRulesPropertyList
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
	PutAttributeTypesSelector(value *AwsDomain_AttributeTypesSelectorProperty)
	// Experimental.
	PutConflictResolution(value *AwsDomain_RuleBasedMatchingConflictResolutionProperty)
	// Experimental.
	PutExportingConfig(value *AwsDomain_RuleBasedMatchingExportingConfigProperty)
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

// The jsii proxy struct for AwsDomain_RuleBasedMatchingPropertyOutputReference
type jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) AttributeTypesSelector() AwsDomain_AttributeTypesSelectorPropertyOutputReference {
	var returns AwsDomain_AttributeTypesSelectorPropertyOutputReference
	_jsii_.Get(
		j,
		"attributeTypesSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) AttributeTypesSelectorInput() *AwsDomain_AttributeTypesSelectorProperty {
	var returns *AwsDomain_AttributeTypesSelectorProperty
	_jsii_.Get(
		j,
		"attributeTypesSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) ConflictResolution() AwsDomain_RuleBasedMatchingConflictResolutionPropertyOutputReference {
	var returns AwsDomain_RuleBasedMatchingConflictResolutionPropertyOutputReference
	_jsii_.Get(
		j,
		"conflictResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) ConflictResolutionInput() *AwsDomain_RuleBasedMatchingConflictResolutionProperty {
	var returns *AwsDomain_RuleBasedMatchingConflictResolutionProperty
	_jsii_.Get(
		j,
		"conflictResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) ExportingConfig() AwsDomain_RuleBasedMatchingExportingConfigPropertyOutputReference {
	var returns AwsDomain_RuleBasedMatchingExportingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"exportingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) ExportingConfigInput() *AwsDomain_RuleBasedMatchingExportingConfigProperty {
	var returns *AwsDomain_RuleBasedMatchingExportingConfigProperty
	_jsii_.Get(
		j,
		"exportingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) InternalValue() *AwsDomain_RuleBasedMatchingProperty {
	var returns *AwsDomain_RuleBasedMatchingProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) MatchingRules() AwsDomain_MatchingRulesPropertyList {
	var returns AwsDomain_MatchingRulesPropertyList
	_jsii_.Get(
		j,
		"matchingRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) MatchingRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"matchingRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) MaxAllowedRuleLevelForMatching() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllowedRuleLevelForMatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) MaxAllowedRuleLevelForMatchingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllowedRuleLevelForMatchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) MaxAllowedRuleLevelForMerging() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllowedRuleLevelForMerging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) MaxAllowedRuleLevelForMergingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllowedRuleLevelForMergingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDomain_RuleBasedMatchingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDomain_RuleBasedMatchingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDomain_RuleBasedMatchingPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.AwsDomain.RuleBasedMatchingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDomain_RuleBasedMatchingPropertyOutputReference_Override(a AwsDomain_RuleBasedMatchingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.AwsDomain.RuleBasedMatchingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference)SetInternalValue(val *AwsDomain_RuleBasedMatchingProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference)SetMaxAllowedRuleLevelForMatching(val *float64) {
	if err := j.validateSetMaxAllowedRuleLevelForMatchingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAllowedRuleLevelForMatching",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference)SetMaxAllowedRuleLevelForMerging(val *float64) {
	if err := j.validateSetMaxAllowedRuleLevelForMergingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAllowedRuleLevelForMerging",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) PutAttributeTypesSelector(value *AwsDomain_AttributeTypesSelectorProperty) {
	if err := a.validatePutAttributeTypesSelectorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAttributeTypesSelector",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) PutConflictResolution(value *AwsDomain_RuleBasedMatchingConflictResolutionProperty) {
	if err := a.validatePutConflictResolutionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConflictResolution",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) PutExportingConfig(value *AwsDomain_RuleBasedMatchingExportingConfigProperty) {
	if err := a.validatePutExportingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExportingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) PutMatchingRules(value interface{}) {
	if err := a.validatePutMatchingRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMatchingRules",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) ResetAttributeTypesSelector() {
	_jsii_.InvokeVoid(
		a,
		"resetAttributeTypesSelector",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) ResetConflictResolution() {
	_jsii_.InvokeVoid(
		a,
		"resetConflictResolution",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) ResetExportingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetExportingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) ResetMatchingRules() {
	_jsii_.InvokeVoid(
		a,
		"resetMatchingRules",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) ResetMaxAllowedRuleLevelForMatching() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxAllowedRuleLevelForMatching",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) ResetMaxAllowedRuleLevelForMerging() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxAllowedRuleLevelForMerging",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDomain_RuleBasedMatchingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

