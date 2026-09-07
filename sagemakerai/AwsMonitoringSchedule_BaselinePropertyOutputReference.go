package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMonitoringSchedule_BaselinePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BaseliningJobName() *string
	// Experimental.
	SetBaseliningJobName(val *string)
	// Experimental.
	BaseliningJobNameInput() *string
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
	// Experimental.
	ConstraintsResource() AwsMonitoringSchedule_ConstraintsResourcePropertyOutputReference
	// Experimental.
	ConstraintsResourceInput() *AwsMonitoringSchedule_ConstraintsResourceProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsMonitoringSchedule_BaselineProperty
	// Experimental.
	SetInternalValue(val *AwsMonitoringSchedule_BaselineProperty)
	// Experimental.
	StatisticsResource() AwsMonitoringSchedule_StatisticsResourcePropertyOutputReference
	// Experimental.
	StatisticsResourceInput() *AwsMonitoringSchedule_StatisticsResourceProperty
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
	PutConstraintsResource(value *AwsMonitoringSchedule_ConstraintsResourceProperty)
	// Experimental.
	PutStatisticsResource(value *AwsMonitoringSchedule_StatisticsResourceProperty)
	// Experimental.
	ResetBaseliningJobName()
	// Experimental.
	ResetConstraintsResource()
	// Experimental.
	ResetStatisticsResource()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMonitoringSchedule_BaselinePropertyOutputReference
type jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) BaseliningJobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseliningJobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) BaseliningJobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseliningJobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) ConstraintsResource() AwsMonitoringSchedule_ConstraintsResourcePropertyOutputReference {
	var returns AwsMonitoringSchedule_ConstraintsResourcePropertyOutputReference
	_jsii_.Get(
		j,
		"constraintsResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) ConstraintsResourceInput() *AwsMonitoringSchedule_ConstraintsResourceProperty {
	var returns *AwsMonitoringSchedule_ConstraintsResourceProperty
	_jsii_.Get(
		j,
		"constraintsResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) InternalValue() *AwsMonitoringSchedule_BaselineProperty {
	var returns *AwsMonitoringSchedule_BaselineProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) StatisticsResource() AwsMonitoringSchedule_StatisticsResourcePropertyOutputReference {
	var returns AwsMonitoringSchedule_StatisticsResourcePropertyOutputReference
	_jsii_.Get(
		j,
		"statisticsResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) StatisticsResourceInput() *AwsMonitoringSchedule_StatisticsResourceProperty {
	var returns *AwsMonitoringSchedule_StatisticsResourceProperty
	_jsii_.Get(
		j,
		"statisticsResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMonitoringSchedule_BaselinePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMonitoringSchedule_BaselinePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMonitoringSchedule_BaselinePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsMonitoringSchedule.BaselinePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMonitoringSchedule_BaselinePropertyOutputReference_Override(a AwsMonitoringSchedule_BaselinePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsMonitoringSchedule.BaselinePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference)SetBaseliningJobName(val *string) {
	if err := j.validateSetBaseliningJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseliningJobName",
		val,
	)
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference)SetInternalValue(val *AwsMonitoringSchedule_BaselineProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) PutConstraintsResource(value *AwsMonitoringSchedule_ConstraintsResourceProperty) {
	if err := a.validatePutConstraintsResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConstraintsResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) PutStatisticsResource(value *AwsMonitoringSchedule_StatisticsResourceProperty) {
	if err := a.validatePutStatisticsResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStatisticsResource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) ResetBaseliningJobName() {
	_jsii_.InvokeVoid(
		a,
		"resetBaseliningJobName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) ResetConstraintsResource() {
	_jsii_.InvokeVoid(
		a,
		"resetConstraintsResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) ResetStatisticsResource() {
	_jsii_.InvokeVoid(
		a,
		"resetStatisticsResource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMonitoringSchedule_BaselinePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

