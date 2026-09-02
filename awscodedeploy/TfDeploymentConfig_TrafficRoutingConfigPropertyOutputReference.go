package awscodedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscodedeploy/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscodedeploy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference interface {
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
	InternalValue() *TfDeploymentConfig_TrafficRoutingConfigProperty
	// Experimental.
	SetInternalValue(val *TfDeploymentConfig_TrafficRoutingConfigProperty)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeBasedCanary() TfDeploymentConfig_TimeBasedCanaryPropertyOutputReference
	// Experimental.
	TimeBasedCanaryInput() *TfDeploymentConfig_TimeBasedCanaryProperty
	// Experimental.
	TimeBasedLinear() TfDeploymentConfig_TimeBasedLinearPropertyOutputReference
	// Experimental.
	TimeBasedLinearInput() *TfDeploymentConfig_TimeBasedLinearProperty
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	PutTimeBasedCanary(value *TfDeploymentConfig_TimeBasedCanaryProperty)
	// Experimental.
	PutTimeBasedLinear(value *TfDeploymentConfig_TimeBasedLinearProperty)
	// Experimental.
	ResetTimeBasedCanary()
	// Experimental.
	ResetTimeBasedLinear()
	// Experimental.
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference
type jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) InternalValue() *TfDeploymentConfig_TrafficRoutingConfigProperty {
	var returns *TfDeploymentConfig_TrafficRoutingConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) TimeBasedCanary() TfDeploymentConfig_TimeBasedCanaryPropertyOutputReference {
	var returns TfDeploymentConfig_TimeBasedCanaryPropertyOutputReference
	_jsii_.Get(
		j,
		"timeBasedCanary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) TimeBasedCanaryInput() *TfDeploymentConfig_TimeBasedCanaryProperty {
	var returns *TfDeploymentConfig_TimeBasedCanaryProperty
	_jsii_.Get(
		j,
		"timeBasedCanaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) TimeBasedLinear() TfDeploymentConfig_TimeBasedLinearPropertyOutputReference {
	var returns TfDeploymentConfig_TimeBasedLinearPropertyOutputReference
	_jsii_.Get(
		j,
		"timeBasedLinear",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) TimeBasedLinearInput() *TfDeploymentConfig_TimeBasedLinearProperty {
	var returns *TfDeploymentConfig_TimeBasedLinearProperty
	_jsii_.Get(
		j,
		"timeBasedLinearInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDeploymentConfig_TrafficRoutingConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codedeploy.TfDeploymentConfig.TrafficRoutingConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference_Override(t TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codedeploy.TfDeploymentConfig.TrafficRoutingConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference)SetInternalValue(val *TfDeploymentConfig_TrafficRoutingConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) PutTimeBasedCanary(value *TfDeploymentConfig_TimeBasedCanaryProperty) {
	if err := t.validatePutTimeBasedCanaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeBasedCanary",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) PutTimeBasedLinear(value *TfDeploymentConfig_TimeBasedLinearProperty) {
	if err := t.validatePutTimeBasedLinearParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeBasedLinear",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) ResetTimeBasedCanary() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeBasedCanary",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) ResetTimeBasedLinear() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeBasedLinear",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		t,
		"resetType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

