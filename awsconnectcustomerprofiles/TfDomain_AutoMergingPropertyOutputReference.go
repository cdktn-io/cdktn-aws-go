package awsconnectcustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconnectcustomerprofiles/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsconnectcustomerprofiles/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_AutoMergingPropertyOutputReference interface {
	cdktn.ComplexObject
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
	ConflictResolution() TfDomain_MatchingAutoMergingConflictResolutionPropertyOutputReference
	// Experimental.
	ConflictResolutionInput() *TfDomain_MatchingAutoMergingConflictResolutionProperty
	// Experimental.
	Consolidation() TfDomain_ConsolidationPropertyOutputReference
	// Experimental.
	ConsolidationInput() *TfDomain_ConsolidationProperty
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
	Fqn() *string
	// Experimental.
	InternalValue() *TfDomain_AutoMergingProperty
	// Experimental.
	SetInternalValue(val *TfDomain_AutoMergingProperty)
	// Experimental.
	MinAllowedConfidenceScoreForMerging() *float64
	// Experimental.
	SetMinAllowedConfidenceScoreForMerging(val *float64)
	// Experimental.
	MinAllowedConfidenceScoreForMergingInput() *float64
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
	PutConflictResolution(value *TfDomain_MatchingAutoMergingConflictResolutionProperty)
	// Experimental.
	PutConsolidation(value *TfDomain_ConsolidationProperty)
	// Experimental.
	ResetConflictResolution()
	// Experimental.
	ResetConsolidation()
	// Experimental.
	ResetMinAllowedConfidenceScoreForMerging()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDomain_AutoMergingPropertyOutputReference
type jsiiProxy_TfDomain_AutoMergingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) ConflictResolution() TfDomain_MatchingAutoMergingConflictResolutionPropertyOutputReference {
	var returns TfDomain_MatchingAutoMergingConflictResolutionPropertyOutputReference
	_jsii_.Get(
		j,
		"conflictResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) ConflictResolutionInput() *TfDomain_MatchingAutoMergingConflictResolutionProperty {
	var returns *TfDomain_MatchingAutoMergingConflictResolutionProperty
	_jsii_.Get(
		j,
		"conflictResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) Consolidation() TfDomain_ConsolidationPropertyOutputReference {
	var returns TfDomain_ConsolidationPropertyOutputReference
	_jsii_.Get(
		j,
		"consolidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) ConsolidationInput() *TfDomain_ConsolidationProperty {
	var returns *TfDomain_ConsolidationProperty
	_jsii_.Get(
		j,
		"consolidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) InternalValue() *TfDomain_AutoMergingProperty {
	var returns *TfDomain_AutoMergingProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) MinAllowedConfidenceScoreForMerging() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAllowedConfidenceScoreForMerging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) MinAllowedConfidenceScoreForMergingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAllowedConfidenceScoreForMergingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_AutoMergingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDomain_AutoMergingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_AutoMergingPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_AutoMergingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.TfDomain.AutoMergingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_AutoMergingPropertyOutputReference_Override(t TfDomain_AutoMergingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.TfDomain.AutoMergingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference)SetInternalValue(val *TfDomain_AutoMergingProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference)SetMinAllowedConfidenceScoreForMerging(val *float64) {
	if err := j.validateSetMinAllowedConfidenceScoreForMergingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minAllowedConfidenceScoreForMerging",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) PutConflictResolution(value *TfDomain_MatchingAutoMergingConflictResolutionProperty) {
	if err := t.validatePutConflictResolutionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConflictResolution",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) PutConsolidation(value *TfDomain_ConsolidationProperty) {
	if err := t.validatePutConsolidationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConsolidation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) ResetConflictResolution() {
	_jsii_.InvokeVoid(
		t,
		"resetConflictResolution",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) ResetConsolidation() {
	_jsii_.InvokeVoid(
		t,
		"resetConsolidation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) ResetMinAllowedConfidenceScoreForMerging() {
	_jsii_.InvokeVoid(
		t,
		"resetMinAllowedConfidenceScoreForMerging",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDomain_AutoMergingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

