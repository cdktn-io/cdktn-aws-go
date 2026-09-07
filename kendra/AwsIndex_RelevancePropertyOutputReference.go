package kendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/kendra/jsii"

	"github.com/cdktn-io/cdktn-aws-go/kendra/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsIndex_RelevancePropertyOutputReference interface {
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
	InternalValue() *AwsIndex_RelevanceProperty
	// Experimental.
	SetInternalValue(val *AwsIndex_RelevanceProperty)
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

// The jsii proxy struct for AwsIndex_RelevancePropertyOutputReference
type jsiiProxy_AwsIndex_RelevancePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) Duration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"duration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) DurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"durationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) Freshness() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"freshness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) FreshnessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"freshnessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) Importance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) ImportanceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"importanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) InternalValue() *AwsIndex_RelevanceProperty {
	var returns *AwsIndex_RelevanceProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) RankOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rankOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) RankOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rankOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) ValuesImportanceMap() *map[string]*float64 {
	var returns *map[string]*float64
	_jsii_.Get(
		j,
		"valuesImportanceMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) ValuesImportanceMapInput() *map[string]*float64 {
	var returns *map[string]*float64
	_jsii_.Get(
		j,
		"valuesImportanceMapInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsIndex_RelevancePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsIndex_RelevancePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsIndex_RelevancePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsIndex_RelevancePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kendra.AwsIndex.RelevancePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsIndex_RelevancePropertyOutputReference_Override(a AwsIndex_RelevancePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kendra.AwsIndex.RelevancePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference)SetDuration(val *string) {
	if err := j.validateSetDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"duration",
		val,
	)
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference)SetFreshness(val interface{}) {
	if err := j.validateSetFreshnessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"freshness",
		val,
	)
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference)SetImportance(val *float64) {
	if err := j.validateSetImportanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importance",
		val,
	)
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference)SetInternalValue(val *AwsIndex_RelevanceProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference)SetRankOrder(val *string) {
	if err := j.validateSetRankOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rankOrder",
		val,
	)
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsIndex_RelevancePropertyOutputReference)SetValuesImportanceMap(val *map[string]*float64) {
	if err := j.validateSetValuesImportanceMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valuesImportanceMap",
		val,
	)
}

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) ResetDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) ResetFreshness() {
	_jsii_.InvokeVoid(
		a,
		"resetFreshness",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) ResetImportance() {
	_jsii_.InvokeVoid(
		a,
		"resetImportance",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) ResetRankOrder() {
	_jsii_.InvokeVoid(
		a,
		"resetRankOrder",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) ResetValuesImportanceMap() {
	_jsii_.InvokeVoid(
		a,
		"resetValuesImportanceMap",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsIndex_RelevancePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

