package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodedeploy/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodedeploy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference interface {
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
	ElbInfo() TfDeploymentGroup_ElbInfoPropertyList
	// Experimental.
	ElbInfoInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDeploymentGroup_LoadBalancerInfoProperty
	// Experimental.
	SetInternalValue(val *TfDeploymentGroup_LoadBalancerInfoProperty)
	// Experimental.
	TargetGroupInfo() TfDeploymentGroup_TargetGroupInfoPropertyList
	// Experimental.
	TargetGroupInfoInput() interface{}
	// Experimental.
	TargetGroupPairInfo() TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference
	// Experimental.
	TargetGroupPairInfoInput() *TfDeploymentGroup_TargetGroupPairInfoProperty
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
	PutElbInfo(value interface{})
	// Experimental.
	PutTargetGroupInfo(value interface{})
	// Experimental.
	PutTargetGroupPairInfo(value *TfDeploymentGroup_TargetGroupPairInfoProperty)
	// Experimental.
	ResetElbInfo()
	// Experimental.
	ResetTargetGroupInfo()
	// Experimental.
	ResetTargetGroupPairInfo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference
type jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ElbInfo() TfDeploymentGroup_ElbInfoPropertyList {
	var returns TfDeploymentGroup_ElbInfoPropertyList
	_jsii_.Get(
		j,
		"elbInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ElbInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elbInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) InternalValue() *TfDeploymentGroup_LoadBalancerInfoProperty {
	var returns *TfDeploymentGroup_LoadBalancerInfoProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) TargetGroupInfo() TfDeploymentGroup_TargetGroupInfoPropertyList {
	var returns TfDeploymentGroup_TargetGroupInfoPropertyList
	_jsii_.Get(
		j,
		"targetGroupInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) TargetGroupInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetGroupInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) TargetGroupPairInfo() TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference {
	var returns TfDeploymentGroup_TargetGroupPairInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"targetGroupPairInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) TargetGroupPairInfoInput() *TfDeploymentGroup_TargetGroupPairInfoProperty {
	var returns *TfDeploymentGroup_TargetGroupPairInfoProperty
	_jsii_.Get(
		j,
		"targetGroupPairInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDeploymentGroup_LoadBalancerInfoPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDeploymentGroup_LoadBalancerInfoPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codedeploy.TfDeploymentGroup.LoadBalancerInfoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDeploymentGroup_LoadBalancerInfoPropertyOutputReference_Override(t TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codedeploy.TfDeploymentGroup.LoadBalancerInfoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference)SetInternalValue(val *TfDeploymentGroup_LoadBalancerInfoProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) PutElbInfo(value interface{}) {
	if err := t.validatePutElbInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putElbInfo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) PutTargetGroupInfo(value interface{}) {
	if err := t.validatePutTargetGroupInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTargetGroupInfo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) PutTargetGroupPairInfo(value *TfDeploymentGroup_TargetGroupPairInfoProperty) {
	if err := t.validatePutTargetGroupPairInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTargetGroupPairInfo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ResetElbInfo() {
	_jsii_.InvokeVoid(
		t,
		"resetElbInfo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ResetTargetGroupInfo() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetGroupInfo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ResetTargetGroupPairInfo() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetGroupPairInfo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

