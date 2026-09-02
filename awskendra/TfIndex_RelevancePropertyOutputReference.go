package awskendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskendra/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskendra/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIndex_RelevancePropertyOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Duration() *string
	// Experimental.
	SetDuration(val *string)
	// Experimental.
	DurationInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	Freshness() interface{}
	// Experimental.
	SetFreshness(val interface{})
	// Experimental.
	FreshnessInput() interface{}
	// Experimental.
	Importance() *float64
	// Experimental.
	SetImportance(val *float64)
	// Experimental.
	ImportanceInput() *float64
	// Experimental.
	InternalValue() *TfIndex_RelevanceProperty
	// Experimental.
	SetInternalValue(val *TfIndex_RelevanceProperty)
	// Experimental.
	RankOrder() *string
	// Experimental.
	SetRankOrder(val *string)
	// Experimental.
	RankOrderInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ValuesImportanceMap() *map[string]*float64
	// Experimental.
	SetValuesImportanceMap(val *map[string]*float64)
	// Experimental.
	ValuesImportanceMapInput() *map[string]*float64
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
	ResetDuration()
	// Experimental.
	ResetFreshness()
	// Experimental.
	ResetImportance()
	// Experimental.
	ResetRankOrder()
	// Experimental.
	ResetValuesImportanceMap()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfIndex_RelevancePropertyOutputReference
type jsiiProxy_TfIndex_RelevancePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference) Freshness() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"freshness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference) FreshnessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"freshnessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference) Importance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference) ImportanceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference) InternalValue() *TfIndex_RelevanceProperty {
	var returns *TfIndex_RelevanceProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference) RankOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rankOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference) RankOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rankOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference) ValuesImportanceMap() *map[string]*float64 {
	var returns *map[string]*float64
	_jsii_.Get(
		j,
		"valuesImportanceMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference) ValuesImportanceMapInput() *map[string]*float64 {
	var returns *map[string]*float64
	_jsii_.Get(
		j,
		"valuesImportanceMapInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfIndex_RelevancePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfIndex_RelevancePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfIndex_RelevancePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfIndex_RelevancePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kendra.TfIndex.RelevancePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfIndex_RelevancePropertyOutputReference_Override(t TfIndex_RelevancePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kendra.TfIndex.RelevancePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference)SetDuration(val *string) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference)SetFreshness(val interface{}) {
	if err := j.validateSetFreshnessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"freshness",
		val,
	)
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference)SetImportance(val *float64) {
	if err := j.validateSetImportanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importance",
		val,
	)
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference)SetInternalValue(val *TfIndex_RelevanceProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference)SetRankOrder(val *string) {
	if err := j.validateSetRankOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rankOrder",
		val,
	)
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfIndex_RelevancePropertyOutputReference)SetValuesImportanceMap(val *map[string]*float64) {
	if err := j.validateSetValuesImportanceMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valuesImportanceMap",
		val,
	)
}

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) ResetDuration() {
	_jsii_.InvokeVoid(
		t,
		"resetDuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) ResetFreshness() {
	_jsii_.InvokeVoid(
		t,
		"resetFreshness",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) ResetImportance() {
	_jsii_.InvokeVoid(
		t,
		"resetImportance",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) ResetRankOrder() {
	_jsii_.InvokeVoid(
		t,
		"resetRankOrder",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) ResetValuesImportanceMap() {
	_jsii_.InvokeVoid(
		t,
		"resetValuesImportanceMap",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfIndex_RelevancePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

