package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconnect/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsconnect/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference interface {
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
	InternalValue() *AwsConnectUserHierarchyStructure_HierarchyStructureProperty
	// Experimental.
	SetInternalValue(val *AwsConnectUserHierarchyStructure_HierarchyStructureProperty)
	// Experimental.
	LevelFive() AwsConnectUserHierarchyStructure_LevelFivePropertyOutputReference
	// Experimental.
	LevelFiveInput() *AwsConnectUserHierarchyStructure_LevelFiveProperty
	// Experimental.
	LevelFour() AwsConnectUserHierarchyStructure_LevelFourPropertyOutputReference
	// Experimental.
	LevelFourInput() *AwsConnectUserHierarchyStructure_LevelFourProperty
	// Experimental.
	LevelOne() AwsConnectUserHierarchyStructure_LevelOnePropertyOutputReference
	// Experimental.
	LevelOneInput() *AwsConnectUserHierarchyStructure_LevelOneProperty
	// Experimental.
	LevelThree() AwsConnectUserHierarchyStructure_LevelThreePropertyOutputReference
	// Experimental.
	LevelThreeInput() *AwsConnectUserHierarchyStructure_LevelThreeProperty
	// Experimental.
	LevelTwo() AwsConnectUserHierarchyStructure_LevelTwoPropertyOutputReference
	// Experimental.
	LevelTwoInput() *AwsConnectUserHierarchyStructure_LevelTwoProperty
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
	PutLevelFive(value *AwsConnectUserHierarchyStructure_LevelFiveProperty)
	// Experimental.
	PutLevelFour(value *AwsConnectUserHierarchyStructure_LevelFourProperty)
	// Experimental.
	PutLevelOne(value *AwsConnectUserHierarchyStructure_LevelOneProperty)
	// Experimental.
	PutLevelThree(value *AwsConnectUserHierarchyStructure_LevelThreeProperty)
	// Experimental.
	PutLevelTwo(value *AwsConnectUserHierarchyStructure_LevelTwoProperty)
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

// The jsii proxy struct for AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference
type jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) InternalValue() *AwsConnectUserHierarchyStructure_HierarchyStructureProperty {
	var returns *AwsConnectUserHierarchyStructure_HierarchyStructureProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelFive() AwsConnectUserHierarchyStructure_LevelFivePropertyOutputReference {
	var returns AwsConnectUserHierarchyStructure_LevelFivePropertyOutputReference
	_jsii_.Get(
		j,
		"levelFive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelFiveInput() *AwsConnectUserHierarchyStructure_LevelFiveProperty {
	var returns *AwsConnectUserHierarchyStructure_LevelFiveProperty
	_jsii_.Get(
		j,
		"levelFiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelFour() AwsConnectUserHierarchyStructure_LevelFourPropertyOutputReference {
	var returns AwsConnectUserHierarchyStructure_LevelFourPropertyOutputReference
	_jsii_.Get(
		j,
		"levelFour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelFourInput() *AwsConnectUserHierarchyStructure_LevelFourProperty {
	var returns *AwsConnectUserHierarchyStructure_LevelFourProperty
	_jsii_.Get(
		j,
		"levelFourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelOne() AwsConnectUserHierarchyStructure_LevelOnePropertyOutputReference {
	var returns AwsConnectUserHierarchyStructure_LevelOnePropertyOutputReference
	_jsii_.Get(
		j,
		"levelOne",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelOneInput() *AwsConnectUserHierarchyStructure_LevelOneProperty {
	var returns *AwsConnectUserHierarchyStructure_LevelOneProperty
	_jsii_.Get(
		j,
		"levelOneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelThree() AwsConnectUserHierarchyStructure_LevelThreePropertyOutputReference {
	var returns AwsConnectUserHierarchyStructure_LevelThreePropertyOutputReference
	_jsii_.Get(
		j,
		"levelThree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelThreeInput() *AwsConnectUserHierarchyStructure_LevelThreeProperty {
	var returns *AwsConnectUserHierarchyStructure_LevelThreeProperty
	_jsii_.Get(
		j,
		"levelThreeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelTwo() AwsConnectUserHierarchyStructure_LevelTwoPropertyOutputReference {
	var returns AwsConnectUserHierarchyStructure_LevelTwoPropertyOutputReference
	_jsii_.Get(
		j,
		"levelTwo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) LevelTwoInput() *AwsConnectUserHierarchyStructure_LevelTwoProperty {
	var returns *AwsConnectUserHierarchyStructure_LevelTwoProperty
	_jsii_.Get(
		j,
		"levelTwoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-connect.AwsConnectUserHierarchyStructure.HierarchyStructurePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference_Override(a AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect.AwsConnectUserHierarchyStructure.HierarchyStructurePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference)SetInternalValue(val *AwsConnectUserHierarchyStructure_HierarchyStructureProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) PutLevelFive(value *AwsConnectUserHierarchyStructure_LevelFiveProperty) {
	if err := a.validatePutLevelFiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLevelFive",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) PutLevelFour(value *AwsConnectUserHierarchyStructure_LevelFourProperty) {
	if err := a.validatePutLevelFourParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLevelFour",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) PutLevelOne(value *AwsConnectUserHierarchyStructure_LevelOneProperty) {
	if err := a.validatePutLevelOneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLevelOne",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) PutLevelThree(value *AwsConnectUserHierarchyStructure_LevelThreeProperty) {
	if err := a.validatePutLevelThreeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLevelThree",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) PutLevelTwo(value *AwsConnectUserHierarchyStructure_LevelTwoProperty) {
	if err := a.validatePutLevelTwoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLevelTwo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ResetLevelFive() {
	_jsii_.InvokeVoid(
		a,
		"resetLevelFive",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ResetLevelFour() {
	_jsii_.InvokeVoid(
		a,
		"resetLevelFour",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ResetLevelOne() {
	_jsii_.InvokeVoid(
		a,
		"resetLevelOne",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ResetLevelThree() {
	_jsii_.InvokeVoid(
		a,
		"resetLevelThree",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ResetLevelTwo() {
	_jsii_.InvokeVoid(
		a,
		"resetLevelTwo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsConnectUserHierarchyStructure_HierarchyStructurePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

