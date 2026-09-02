package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodedeploy/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodedeploy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference interface {
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
	InternalValue() *TfDeploymentGroup_TargetGroupPairInfoProperty
	// Experimental.
	SetInternalValue(val *TfDeploymentGroup_TargetGroupPairInfoProperty)
	// Experimental.
	ProdTrafficRoute() TfDeploymentGroup_ProdTrafficRoutePropertyOutputReference
	// Experimental.
	ProdTrafficRouteInput() *TfDeploymentGroup_ProdTrafficRouteProperty
	// Experimental.
	TargetGroup() TfDeploymentGroup_TargetGroupPropertyList
	// Experimental.
	TargetGroupInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TestTrafficRoute() TfDeploymentGroup_TestTrafficRoutePropertyOutputReference
	// Experimental.
	TestTrafficRouteInput() *TfDeploymentGroup_TestTrafficRouteProperty
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
	PutProdTrafficRoute(value *TfDeploymentGroup_ProdTrafficRouteProperty)
	// Experimental.
	PutTargetGroup(value interface{})
	// Experimental.
	PutTestTrafficRoute(value *TfDeploymentGroup_TestTrafficRouteProperty)
	// Experimental.
	ResetTestTrafficRoute()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference
type jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) InternalValue() *TfDeploymentGroup_TargetGroupPairInfoProperty {
	var returns *TfDeploymentGroup_TargetGroupPairInfoProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) ProdTrafficRoute() TfDeploymentGroup_ProdTrafficRoutePropertyOutputReference {
	var returns TfDeploymentGroup_ProdTrafficRoutePropertyOutputReference
	_jsii_.Get(
		j,
		"prodTrafficRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) ProdTrafficRouteInput() *TfDeploymentGroup_ProdTrafficRouteProperty {
	var returns *TfDeploymentGroup_ProdTrafficRouteProperty
	_jsii_.Get(
		j,
		"prodTrafficRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) TargetGroup() TfDeploymentGroup_TargetGroupPropertyList {
	var returns TfDeploymentGroup_TargetGroupPropertyList
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) TargetGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) TestTrafficRoute() TfDeploymentGroup_TestTrafficRoutePropertyOutputReference {
	var returns TfDeploymentGroup_TestTrafficRoutePropertyOutputReference
	_jsii_.Get(
		j,
		"testTrafficRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) TestTrafficRouteInput() *TfDeploymentGroup_TestTrafficRouteProperty {
	var returns *TfDeploymentGroup_TestTrafficRouteProperty
	_jsii_.Get(
		j,
		"testTrafficRouteInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDeploymentGroup_TargetGroupPairInfoPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codedeploy.TfDeploymentGroup.TargetGroupPairInfoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference_Override(t TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codedeploy.TfDeploymentGroup.TargetGroupPairInfoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference)SetInternalValue(val *TfDeploymentGroup_TargetGroupPairInfoProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) PutProdTrafficRoute(value *TfDeploymentGroup_ProdTrafficRouteProperty) {
	if err := t.validatePutProdTrafficRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProdTrafficRoute",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) PutTargetGroup(value interface{}) {
	if err := t.validatePutTargetGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTargetGroup",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) PutTestTrafficRoute(value *TfDeploymentGroup_TestTrafficRouteProperty) {
	if err := t.validatePutTestTrafficRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTestTrafficRoute",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) ResetTestTrafficRoute() {
	_jsii_.InvokeVoid(
		t,
		"resetTestTrafficRoute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

