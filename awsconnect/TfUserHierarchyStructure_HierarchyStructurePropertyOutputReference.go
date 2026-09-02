package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconnect/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsconnect/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference interface {
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
	InternalValue() *TfUserHierarchyStructure_HierarchyStructureProperty
	// Experimental.
	SetInternalValue(val *TfUserHierarchyStructure_HierarchyStructureProperty)
	// Experimental.
	LevelFive() TfUserHierarchyStructure_LevelFivePropertyOutputReference
	// Experimental.
	LevelFiveInput() *TfUserHierarchyStructure_LevelFiveProperty
	// Experimental.
	LevelFour() TfUserHierarchyStructure_LevelFourPropertyOutputReference
	// Experimental.
	LevelFourInput() *TfUserHierarchyStructure_LevelFourProperty
	// Experimental.
	LevelOne() TfUserHierarchyStructure_LevelOnePropertyOutputReference
	// Experimental.
	LevelOneInput() *TfUserHierarchyStructure_LevelOneProperty
	// Experimental.
	LevelThree() TfUserHierarchyStructure_LevelThreePropertyOutputReference
	// Experimental.
	LevelThreeInput() *TfUserHierarchyStructure_LevelThreeProperty
	// Experimental.
	LevelTwo() TfUserHierarchyStructure_LevelTwoPropertyOutputReference
	// Experimental.
	LevelTwoInput() *TfUserHierarchyStructure_LevelTwoProperty
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
	PutLevelFive(value *TfUserHierarchyStructure_LevelFiveProperty)
	// Experimental.
	PutLevelFour(value *TfUserHierarchyStructure_LevelFourProperty)
	// Experimental.
	PutLevelOne(value *TfUserHierarchyStructure_LevelOneProperty)
	// Experimental.
	PutLevelThree(value *TfUserHierarchyStructure_LevelThreeProperty)
	// Experimental.
	PutLevelTwo(value *TfUserHierarchyStructure_LevelTwoProperty)
	// Experimental.
	ResetLevelFive()
	// Experimental.
	ResetLevelFour()
	// Experimental.
	ResetLevelOne()
	// Experimental.
	ResetLevelThree()
	// Experimental.
	ResetLevelTwo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference
type jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) InternalValue() *TfUserHierarchyStructure_HierarchyStructureProperty {
	var returns *TfUserHierarchyStructure_HierarchyStructureProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelFive() TfUserHierarchyStructure_LevelFivePropertyOutputReference {
	var returns TfUserHierarchyStructure_LevelFivePropertyOutputReference
	_jsii_.Get(
		j,
		"levelFive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelFiveInput() *TfUserHierarchyStructure_LevelFiveProperty {
	var returns *TfUserHierarchyStructure_LevelFiveProperty
	_jsii_.Get(
		j,
		"levelFiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelFour() TfUserHierarchyStructure_LevelFourPropertyOutputReference {
	var returns TfUserHierarchyStructure_LevelFourPropertyOutputReference
	_jsii_.Get(
		j,
		"levelFour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelFourInput() *TfUserHierarchyStructure_LevelFourProperty {
	var returns *TfUserHierarchyStructure_LevelFourProperty
	_jsii_.Get(
		j,
		"levelFourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelOne() TfUserHierarchyStructure_LevelOnePropertyOutputReference {
	var returns TfUserHierarchyStructure_LevelOnePropertyOutputReference
	_jsii_.Get(
		j,
		"levelOne",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelOneInput() *TfUserHierarchyStructure_LevelOneProperty {
	var returns *TfUserHierarchyStructure_LevelOneProperty
	_jsii_.Get(
		j,
		"levelOneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelThree() TfUserHierarchyStructure_LevelThreePropertyOutputReference {
	var returns TfUserHierarchyStructure_LevelThreePropertyOutputReference
	_jsii_.Get(
		j,
		"levelThree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelThreeInput() *TfUserHierarchyStructure_LevelThreeProperty {
	var returns *TfUserHierarchyStructure_LevelThreeProperty
	_jsii_.Get(
		j,
		"levelThreeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelTwo() TfUserHierarchyStructure_LevelTwoPropertyOutputReference {
	var returns TfUserHierarchyStructure_LevelTwoPropertyOutputReference
	_jsii_.Get(
		j,
		"levelTwo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelTwoInput() *TfUserHierarchyStructure_LevelTwoProperty {
	var returns *TfUserHierarchyStructure_LevelTwoProperty
	_jsii_.Get(
		j,
		"levelTwoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfUserHierarchyStructure_HierarchyStructurePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfUserHierarchyStructure_HierarchyStructurePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-connect.TfUserHierarchyStructure.HierarchyStructurePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfUserHierarchyStructure_HierarchyStructurePropertyOutputReference_Override(t TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect.TfUserHierarchyStructure.HierarchyStructurePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference)SetInternalValue(val *TfUserHierarchyStructure_HierarchyStructureProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) PutLevelFive(value *TfUserHierarchyStructure_LevelFiveProperty) {
	if err := t.validatePutLevelFiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLevelFive",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) PutLevelFour(value *TfUserHierarchyStructure_LevelFourProperty) {
	if err := t.validatePutLevelFourParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLevelFour",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) PutLevelOne(value *TfUserHierarchyStructure_LevelOneProperty) {
	if err := t.validatePutLevelOneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLevelOne",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) PutLevelThree(value *TfUserHierarchyStructure_LevelThreeProperty) {
	if err := t.validatePutLevelThreeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLevelThree",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) PutLevelTwo(value *TfUserHierarchyStructure_LevelTwoProperty) {
	if err := t.validatePutLevelTwoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLevelTwo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ResetLevelFive() {
	_jsii_.InvokeVoid(
		t,
		"resetLevelFive",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ResetLevelFour() {
	_jsii_.InvokeVoid(
		t,
		"resetLevelFour",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ResetLevelOne() {
	_jsii_.InvokeVoid(
		t,
		"resetLevelOne",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ResetLevelThree() {
	_jsii_.InvokeVoid(
		t,
		"resetLevelThree",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ResetLevelTwo() {
	_jsii_.InvokeVoid(
		t,
		"resetLevelTwo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

