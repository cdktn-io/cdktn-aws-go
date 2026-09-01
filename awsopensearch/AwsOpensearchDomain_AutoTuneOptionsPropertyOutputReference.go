package awsopensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsopensearch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsopensearch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference interface {
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
	DesiredState() *string
	// Experimental.
	SetDesiredState(val *string)
	// Experimental.
	DesiredStateInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsOpensearchDomain_AutoTuneOptionsProperty
	// Experimental.
	SetInternalValue(val *AwsOpensearchDomain_AutoTuneOptionsProperty)
	// Experimental.
	MaintenanceSchedule() AwsOpensearchDomain_MaintenanceSchedulePropertyList
	// Experimental.
	MaintenanceScheduleInput() interface{}
	// Experimental.
	RollbackOnDisable() *string
	// Experimental.
	SetRollbackOnDisable(val *string)
	// Experimental.
	RollbackOnDisableInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UseOffPeakWindow() interface{}
	// Experimental.
	SetUseOffPeakWindow(val interface{})
	// Experimental.
	UseOffPeakWindowInput() interface{}
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
	PutMaintenanceSchedule(value interface{})
	// Experimental.
	ResetMaintenanceSchedule()
	// Experimental.
	ResetRollbackOnDisable()
	// Experimental.
	ResetUseOffPeakWindow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference
type jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) DesiredState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) DesiredStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"desiredStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) InternalValue() *AwsOpensearchDomain_AutoTuneOptionsProperty {
	var returns *AwsOpensearchDomain_AutoTuneOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) MaintenanceSchedule() AwsOpensearchDomain_MaintenanceSchedulePropertyList {
	var returns AwsOpensearchDomain_MaintenanceSchedulePropertyList
	_jsii_.Get(
		j,
		"maintenanceSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) MaintenanceScheduleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) RollbackOnDisable() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rollbackOnDisable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) RollbackOnDisableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rollbackOnDisableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) UseOffPeakWindow() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOffPeakWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) UseOffPeakWindowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOffPeakWindowInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsOpensearchDomain_AutoTuneOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-opensearch.AwsOpensearchDomain.AutoTuneOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference_Override(a AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-opensearch.AwsOpensearchDomain.AutoTuneOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference)SetDesiredState(val *string) {
	if err := j.validateSetDesiredStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredState",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference)SetInternalValue(val *AwsOpensearchDomain_AutoTuneOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference)SetRollbackOnDisable(val *string) {
	if err := j.validateSetRollbackOnDisableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rollbackOnDisable",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference)SetUseOffPeakWindow(val interface{}) {
	if err := j.validateSetUseOffPeakWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useOffPeakWindow",
		val,
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) PutMaintenanceSchedule(value interface{}) {
	if err := a.validatePutMaintenanceScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMaintenanceSchedule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) ResetMaintenanceSchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetMaintenanceSchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) ResetRollbackOnDisable() {
	_jsii_.InvokeVoid(
		a,
		"resetRollbackOnDisable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) ResetUseOffPeakWindow() {
	_jsii_.InvokeVoid(
		a,
		"resetUseOffPeakWindow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsOpensearchDomain_AutoTuneOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

