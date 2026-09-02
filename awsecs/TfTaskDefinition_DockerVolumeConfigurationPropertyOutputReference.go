package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Autoprovision() interface{}
	// Experimental.
	SetAutoprovision(val interface{})
	// Experimental.
	AutoprovisionInput() interface{}
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
	Driver() *string
	// Experimental.
	SetDriver(val *string)
	// Experimental.
	DriverInput() *string
	// Experimental.
	DriverOpts() *map[string]*string
	// Experimental.
	SetDriverOpts(val *map[string]*string)
	// Experimental.
	DriverOptsInput() *map[string]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfTaskDefinition_DockerVolumeConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfTaskDefinition_DockerVolumeConfigurationProperty)
	// Experimental.
	Labels() *map[string]*string
	// Experimental.
	SetLabels(val *map[string]*string)
	// Experimental.
	LabelsInput() *map[string]*string
	// Experimental.
	Scope() *string
	// Experimental.
	SetScope(val *string)
	// Experimental.
	ScopeInput() *string
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
	ResetAutoprovision()
	// Experimental.
	ResetDriver()
	// Experimental.
	ResetDriverOpts()
	// Experimental.
	ResetLabels()
	// Experimental.
	ResetScope()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference
type jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) Autoprovision() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoprovision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) AutoprovisionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoprovisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) Driver() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driver",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) DriverInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"driverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) DriverOpts() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"driverOpts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) DriverOptsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"driverOptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) InternalValue() *TfTaskDefinition_DockerVolumeConfigurationProperty {
	var returns *TfTaskDefinition_DockerVolumeConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTaskDefinition_DockerVolumeConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.TfTaskDefinition.DockerVolumeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference_Override(t TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.TfTaskDefinition.DockerVolumeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference)SetAutoprovision(val interface{}) {
	if err := j.validateSetAutoprovisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoprovision",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference)SetDriver(val *string) {
	if err := j.validateSetDriverParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driver",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference)SetDriverOpts(val *map[string]*string) {
	if err := j.validateSetDriverOptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"driverOpts",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference)SetInternalValue(val *TfTaskDefinition_DockerVolumeConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) ResetAutoprovision() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoprovision",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) ResetDriver() {
	_jsii_.InvokeVoid(
		t,
		"resetDriver",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) ResetDriverOpts() {
	_jsii_.InvokeVoid(
		t,
		"resetDriverOpts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		t,
		"resetLabels",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		t,
		"resetScope",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTaskDefinition_DockerVolumeConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

