package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsssm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsssm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchLogGroupName() *string
	// Experimental.
	SetCloudwatchLogGroupName(val *string)
	// Experimental.
	CloudwatchLogGroupNameInput() *string
	// Experimental.
	CloudwatchOutputEnabled() interface{}
	// Experimental.
	SetCloudwatchOutputEnabled(val interface{})
	// Experimental.
	CloudwatchOutputEnabledInput() interface{}
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
	InternalValue() *AwsSsmMaintenanceWindowTask_CloudwatchConfigProperty
	// Experimental.
	SetInternalValue(val *AwsSsmMaintenanceWindowTask_CloudwatchConfigProperty)
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
	ResetCloudwatchLogGroupName()
	// Experimental.
	ResetCloudwatchOutputEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference
type jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) CloudwatchLogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) CloudwatchLogGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudwatchLogGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) CloudwatchOutputEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchOutputEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) CloudwatchOutputEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchOutputEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) InternalValue() *AwsSsmMaintenanceWindowTask_CloudwatchConfigProperty {
	var returns *AwsSsmMaintenanceWindowTask_CloudwatchConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ssm.AwsSsmMaintenanceWindowTask.CloudwatchConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference_Override(a AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ssm.AwsSsmMaintenanceWindowTask.CloudwatchConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference)SetCloudwatchLogGroupName(val *string) {
	if err := j.validateSetCloudwatchLogGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchLogGroupName",
		val,
	)
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference)SetCloudwatchOutputEnabled(val interface{}) {
	if err := j.validateSetCloudwatchOutputEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchOutputEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference)SetInternalValue(val *AwsSsmMaintenanceWindowTask_CloudwatchConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) ResetCloudwatchLogGroupName() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchLogGroupName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) ResetCloudwatchOutputEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchOutputEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSsmMaintenanceWindowTask_CloudwatchConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

