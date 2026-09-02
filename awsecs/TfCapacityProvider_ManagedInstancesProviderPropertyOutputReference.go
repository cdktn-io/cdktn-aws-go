package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference interface {
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
	InfrastructureOptimization() TfCapacityProvider_InfrastructureOptimizationPropertyOutputReference
	// Experimental.
	InfrastructureOptimizationInput() *TfCapacityProvider_InfrastructureOptimizationProperty
	// Experimental.
	InfrastructureRoleArn() *string
	// Experimental.
	SetInfrastructureRoleArn(val *string)
	// Experimental.
	InfrastructureRoleArnInput() *string
	// Experimental.
	InstanceLaunchTemplate() TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference
	// Experimental.
	InstanceLaunchTemplateInput() *TfCapacityProvider_InstanceLaunchTemplateProperty
	// Experimental.
	InternalValue() *TfCapacityProvider_ManagedInstancesProviderProperty
	// Experimental.
	SetInternalValue(val *TfCapacityProvider_ManagedInstancesProviderProperty)
	// Experimental.
	PropagateTags() *string
	// Experimental.
	SetPropagateTags(val *string)
	// Experimental.
	PropagateTagsInput() *string
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
	PutInfrastructureOptimization(value *TfCapacityProvider_InfrastructureOptimizationProperty)
	// Experimental.
	PutInstanceLaunchTemplate(value *TfCapacityProvider_InstanceLaunchTemplateProperty)
	// Experimental.
	ResetInfrastructureOptimization()
	// Experimental.
	ResetPropagateTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference
type jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InfrastructureOptimization() TfCapacityProvider_InfrastructureOptimizationPropertyOutputReference {
	var returns TfCapacityProvider_InfrastructureOptimizationPropertyOutputReference
	_jsii_.Get(
		j,
		"infrastructureOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InfrastructureOptimizationInput() *TfCapacityProvider_InfrastructureOptimizationProperty {
	var returns *TfCapacityProvider_InfrastructureOptimizationProperty
	_jsii_.Get(
		j,
		"infrastructureOptimizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InfrastructureRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InfrastructureRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InstanceLaunchTemplate() TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference {
	var returns TfCapacityProvider_InstanceLaunchTemplatePropertyOutputReference
	_jsii_.Get(
		j,
		"instanceLaunchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InstanceLaunchTemplateInput() *TfCapacityProvider_InstanceLaunchTemplateProperty {
	var returns *TfCapacityProvider_InstanceLaunchTemplateProperty
	_jsii_.Get(
		j,
		"instanceLaunchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InternalValue() *TfCapacityProvider_ManagedInstancesProviderProperty {
	var returns *TfCapacityProvider_ManagedInstancesProviderProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) PropagateTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) PropagateTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCapacityProvider_ManagedInstancesProviderPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCapacityProvider_ManagedInstancesProviderPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.TfCapacityProvider.ManagedInstancesProviderPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCapacityProvider_ManagedInstancesProviderPropertyOutputReference_Override(t TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.TfCapacityProvider.ManagedInstancesProviderPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetInfrastructureRoleArn(val *string) {
	if err := j.validateSetInfrastructureRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"infrastructureRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetInternalValue(val *TfCapacityProvider_ManagedInstancesProviderProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetPropagateTags(val *string) {
	if err := j.validateSetPropagateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) PutInfrastructureOptimization(value *TfCapacityProvider_InfrastructureOptimizationProperty) {
	if err := t.validatePutInfrastructureOptimizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInfrastructureOptimization",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) PutInstanceLaunchTemplate(value *TfCapacityProvider_InstanceLaunchTemplateProperty) {
	if err := t.validatePutInstanceLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInstanceLaunchTemplate",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ResetInfrastructureOptimization() {
	_jsii_.InvokeVoid(
		t,
		"resetInfrastructureOptimization",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ResetPropagateTags() {
	_jsii_.InvokeVoid(
		t,
		"resetPropagateTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

