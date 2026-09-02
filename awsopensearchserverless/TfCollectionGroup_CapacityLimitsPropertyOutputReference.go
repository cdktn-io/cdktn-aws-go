package awsopensearchserverless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsopensearchserverless/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsopensearchserverless/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCollectionGroup_CapacityLimitsPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MaxIndexingCapacityInOcu() *float64
	// Experimental.
	SetMaxIndexingCapacityInOcu(val *float64)
	// Experimental.
	MaxIndexingCapacityInOcuInput() *float64
	// Experimental.
	MaxSearchCapacityInOcu() *float64
	// Experimental.
	SetMaxSearchCapacityInOcu(val *float64)
	// Experimental.
	MaxSearchCapacityInOcuInput() *float64
	// Experimental.
	MinIndexingCapacityInOcu() *float64
	// Experimental.
	SetMinIndexingCapacityInOcu(val *float64)
	// Experimental.
	MinIndexingCapacityInOcuInput() *float64
	// Experimental.
	MinSearchCapacityInOcu() *float64
	// Experimental.
	SetMinSearchCapacityInOcu(val *float64)
	// Experimental.
	MinSearchCapacityInOcuInput() *float64
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
	ResetMaxIndexingCapacityInOcu()
	// Experimental.
	ResetMaxSearchCapacityInOcu()
	// Experimental.
	ResetMinIndexingCapacityInOcu()
	// Experimental.
	ResetMinSearchCapacityInOcu()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCollectionGroup_CapacityLimitsPropertyOutputReference
type jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) MaxIndexingCapacityInOcu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIndexingCapacityInOcu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) MaxIndexingCapacityInOcuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIndexingCapacityInOcuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) MaxSearchCapacityInOcu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSearchCapacityInOcu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) MaxSearchCapacityInOcuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxSearchCapacityInOcuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) MinIndexingCapacityInOcu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIndexingCapacityInOcu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) MinIndexingCapacityInOcuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minIndexingCapacityInOcuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) MinSearchCapacityInOcu() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSearchCapacityInOcu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) MinSearchCapacityInOcuInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSearchCapacityInOcuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCollectionGroup_CapacityLimitsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfCollectionGroup_CapacityLimitsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCollectionGroup_CapacityLimitsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-opensearch-serverless.TfCollectionGroup.CapacityLimitsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCollectionGroup_CapacityLimitsPropertyOutputReference_Override(t TfCollectionGroup_CapacityLimitsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-opensearch-serverless.TfCollectionGroup.CapacityLimitsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference)SetMaxIndexingCapacityInOcu(val *float64) {
	if err := j.validateSetMaxIndexingCapacityInOcuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIndexingCapacityInOcu",
		val,
	)
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference)SetMaxSearchCapacityInOcu(val *float64) {
	if err := j.validateSetMaxSearchCapacityInOcuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxSearchCapacityInOcu",
		val,
	)
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference)SetMinIndexingCapacityInOcu(val *float64) {
	if err := j.validateSetMinIndexingCapacityInOcuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minIndexingCapacityInOcu",
		val,
	)
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference)SetMinSearchCapacityInOcu(val *float64) {
	if err := j.validateSetMinSearchCapacityInOcuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSearchCapacityInOcu",
		val,
	)
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) ResetMaxIndexingCapacityInOcu() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxIndexingCapacityInOcu",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) ResetMaxSearchCapacityInOcu() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxSearchCapacityInOcu",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) ResetMinIndexingCapacityInOcu() {
	_jsii_.InvokeVoid(
		t,
		"resetMinIndexingCapacityInOcu",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) ResetMinSearchCapacityInOcu() {
	_jsii_.InvokeVoid(
		t,
		"resetMinSearchCapacityInOcu",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCollectionGroup_CapacityLimitsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

