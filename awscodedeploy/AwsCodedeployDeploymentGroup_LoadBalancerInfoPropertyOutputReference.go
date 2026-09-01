package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodedeploy/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodedeploy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference interface {
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
	ElbInfo() AwsCodedeployDeploymentGroup_ElbInfoPropertyList
	// Experimental.
	ElbInfoInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsCodedeployDeploymentGroup_LoadBalancerInfoProperty
	// Experimental.
	SetInternalValue(val *AwsCodedeployDeploymentGroup_LoadBalancerInfoProperty)
	// Experimental.
	TargetGroupInfo() AwsCodedeployDeploymentGroup_TargetGroupInfoPropertyList
	// Experimental.
	TargetGroupInfoInput() interface{}
	// Experimental.
	TargetGroupPairInfo() AwsCodedeployDeploymentGroup_TargetGroupPairInfoPropertyOutputReference
	// Experimental.
	TargetGroupPairInfoInput() *AwsCodedeployDeploymentGroup_TargetGroupPairInfoProperty
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
	PutTargetGroupPairInfo(value *AwsCodedeployDeploymentGroup_TargetGroupPairInfoProperty)
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

// The jsii proxy struct for AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference
type jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ElbInfo() AwsCodedeployDeploymentGroup_ElbInfoPropertyList {
	var returns AwsCodedeployDeploymentGroup_ElbInfoPropertyList
	_jsii_.Get(
		j,
		"elbInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ElbInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elbInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) InternalValue() *AwsCodedeployDeploymentGroup_LoadBalancerInfoProperty {
	var returns *AwsCodedeployDeploymentGroup_LoadBalancerInfoProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) TargetGroupInfo() AwsCodedeployDeploymentGroup_TargetGroupInfoPropertyList {
	var returns AwsCodedeployDeploymentGroup_TargetGroupInfoPropertyList
	_jsii_.Get(
		j,
		"targetGroupInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) TargetGroupInfoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetGroupInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) TargetGroupPairInfo() AwsCodedeployDeploymentGroup_TargetGroupPairInfoPropertyOutputReference {
	var returns AwsCodedeployDeploymentGroup_TargetGroupPairInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"targetGroupPairInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) TargetGroupPairInfoInput() *AwsCodedeployDeploymentGroup_TargetGroupPairInfoProperty {
	var returns *AwsCodedeployDeploymentGroup_TargetGroupPairInfoProperty
	_jsii_.Get(
		j,
		"targetGroupPairInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codedeploy.AwsCodedeployDeploymentGroup.LoadBalancerInfoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference_Override(a AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codedeploy.AwsCodedeployDeploymentGroup.LoadBalancerInfoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference)SetInternalValue(val *AwsCodedeployDeploymentGroup_LoadBalancerInfoProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) PutElbInfo(value interface{}) {
	if err := a.validatePutElbInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putElbInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) PutTargetGroupInfo(value interface{}) {
	if err := a.validatePutTargetGroupInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTargetGroupInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) PutTargetGroupPairInfo(value *AwsCodedeployDeploymentGroup_TargetGroupPairInfoProperty) {
	if err := a.validatePutTargetGroupPairInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTargetGroupPairInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ResetElbInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetElbInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ResetTargetGroupInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetGroupInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ResetTargetGroupPairInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetGroupPairInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCodedeployDeploymentGroup_LoadBalancerInfoPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

