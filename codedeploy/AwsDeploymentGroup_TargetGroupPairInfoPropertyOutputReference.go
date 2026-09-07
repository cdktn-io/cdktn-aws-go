package codedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/codedeploy/jsii"

	"github.com/cdktn-io/cdktn-aws-go/codedeploy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference interface {
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
	InternalValue() *AwsDeploymentGroup_TargetGroupPairInfoProperty
	// Experimental.
	SetInternalValue(val *AwsDeploymentGroup_TargetGroupPairInfoProperty)
	// Experimental.
	ProdTrafficRoute() AwsDeploymentGroup_ProdTrafficRoutePropertyOutputReference
	// Experimental.
	ProdTrafficRouteInput() *AwsDeploymentGroup_ProdTrafficRouteProperty
	// Experimental.
	TargetGroup() AwsDeploymentGroup_TargetGroupPropertyList
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
	TestTrafficRoute() AwsDeploymentGroup_TestTrafficRoutePropertyOutputReference
	// Experimental.
	TestTrafficRouteInput() *AwsDeploymentGroup_TestTrafficRouteProperty
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
	PutProdTrafficRoute(value *AwsDeploymentGroup_ProdTrafficRouteProperty)
	// Experimental.
	PutTargetGroup(value interface{})
	// Experimental.
	PutTestTrafficRoute(value *AwsDeploymentGroup_TestTrafficRouteProperty)
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

// The jsii proxy struct for AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference
type jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) InternalValue() *AwsDeploymentGroup_TargetGroupPairInfoProperty {
	var returns *AwsDeploymentGroup_TargetGroupPairInfoProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) ProdTrafficRoute() AwsDeploymentGroup_ProdTrafficRoutePropertyOutputReference {
	var returns AwsDeploymentGroup_ProdTrafficRoutePropertyOutputReference
	_jsii_.Get(
		j,
		"prodTrafficRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) ProdTrafficRouteInput() *AwsDeploymentGroup_ProdTrafficRouteProperty {
	var returns *AwsDeploymentGroup_ProdTrafficRouteProperty
	_jsii_.Get(
		j,
		"prodTrafficRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) TargetGroup() AwsDeploymentGroup_TargetGroupPropertyList {
	var returns AwsDeploymentGroup_TargetGroupPropertyList
	_jsii_.Get(
		j,
		"targetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) TargetGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) TestTrafficRoute() AwsDeploymentGroup_TestTrafficRoutePropertyOutputReference {
	var returns AwsDeploymentGroup_TestTrafficRoutePropertyOutputReference
	_jsii_.Get(
		j,
		"testTrafficRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) TestTrafficRouteInput() *AwsDeploymentGroup_TestTrafficRouteProperty {
	var returns *AwsDeploymentGroup_TestTrafficRouteProperty
	_jsii_.Get(
		j,
		"testTrafficRouteInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codedeploy.AwsDeploymentGroup.TargetGroupPairInfoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference_Override(a AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codedeploy.AwsDeploymentGroup.TargetGroupPairInfoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference)SetInternalValue(val *AwsDeploymentGroup_TargetGroupPairInfoProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) PutProdTrafficRoute(value *AwsDeploymentGroup_ProdTrafficRouteProperty) {
	if err := a.validatePutProdTrafficRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProdTrafficRoute",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) PutTargetGroup(value interface{}) {
	if err := a.validatePutTargetGroupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTargetGroup",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) PutTestTrafficRoute(value *AwsDeploymentGroup_TestTrafficRouteProperty) {
	if err := a.validatePutTestTrafficRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTestTrafficRoute",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) ResetTestTrafficRoute() {
	_jsii_.InvokeVoid(
		a,
		"resetTestTrafficRoute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDeploymentGroup_TargetGroupPairInfoPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

