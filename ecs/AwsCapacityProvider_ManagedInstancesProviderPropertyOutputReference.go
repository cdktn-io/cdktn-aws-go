package ecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/ecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference interface {
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
	InfrastructureOptimization() AwsCapacityProvider_InfrastructureOptimizationPropertyOutputReference
	// Experimental.
	InfrastructureOptimizationInput() *AwsCapacityProvider_InfrastructureOptimizationProperty
	// Experimental.
	InfrastructureRoleArn() *string
	// Experimental.
	SetInfrastructureRoleArn(val *string)
	// Experimental.
	InfrastructureRoleArnInput() *string
	// Experimental.
	InstanceLaunchTemplate() AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference
	// Experimental.
	InstanceLaunchTemplateInput() *AwsCapacityProvider_InstanceLaunchTemplateProperty
	// Experimental.
	InternalValue() *AwsCapacityProvider_ManagedInstancesProviderProperty
	// Experimental.
	SetInternalValue(val *AwsCapacityProvider_ManagedInstancesProviderProperty)
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
	PutInfrastructureOptimization(value *AwsCapacityProvider_InfrastructureOptimizationProperty)
	// Experimental.
	PutInstanceLaunchTemplate(value *AwsCapacityProvider_InstanceLaunchTemplateProperty)
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

// The jsii proxy struct for AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference
type jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InfrastructureOptimization() AwsCapacityProvider_InfrastructureOptimizationPropertyOutputReference {
	var returns AwsCapacityProvider_InfrastructureOptimizationPropertyOutputReference
	_jsii_.Get(
		j,
		"infrastructureOptimization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InfrastructureOptimizationInput() *AwsCapacityProvider_InfrastructureOptimizationProperty {
	var returns *AwsCapacityProvider_InfrastructureOptimizationProperty
	_jsii_.Get(
		j,
		"infrastructureOptimizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InfrastructureRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InfrastructureRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InstanceLaunchTemplate() AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference {
	var returns AwsCapacityProvider_InstanceLaunchTemplatePropertyOutputReference
	_jsii_.Get(
		j,
		"instanceLaunchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InstanceLaunchTemplateInput() *AwsCapacityProvider_InstanceLaunchTemplateProperty {
	var returns *AwsCapacityProvider_InstanceLaunchTemplateProperty
	_jsii_.Get(
		j,
		"instanceLaunchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InternalValue() *AwsCapacityProvider_ManagedInstancesProviderProperty {
	var returns *AwsCapacityProvider_ManagedInstancesProviderProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) PropagateTags() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) PropagateTagsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"propagateTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCapacityProvider_ManagedInstancesProviderPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsCapacityProvider.ManagedInstancesProviderPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference_Override(a AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsCapacityProvider.ManagedInstancesProviderPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetInfrastructureRoleArn(val *string) {
	if err := j.validateSetInfrastructureRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"infrastructureRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetInternalValue(val *AwsCapacityProvider_ManagedInstancesProviderProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetPropagateTags(val *string) {
	if err := j.validateSetPropagateTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"propagateTags",
		val,
	)
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) PutInfrastructureOptimization(value *AwsCapacityProvider_InfrastructureOptimizationProperty) {
	if err := a.validatePutInfrastructureOptimizationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInfrastructureOptimization",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) PutInstanceLaunchTemplate(value *AwsCapacityProvider_InstanceLaunchTemplateProperty) {
	if err := a.validatePutInstanceLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInstanceLaunchTemplate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ResetInfrastructureOptimization() {
	_jsii_.InvokeVoid(
		a,
		"resetInfrastructureOptimization",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ResetPropagateTags() {
	_jsii_.InvokeVoid(
		a,
		"resetPropagateTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCapacityProvider_ManagedInstancesProviderPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

