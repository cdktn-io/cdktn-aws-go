package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Baseline() TfMonitoringSchedule_BaselinePropertyOutputReference
	// Experimental.
	BaselineInput() *TfMonitoringSchedule_BaselineProperty
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
	Environment() *map[string]*string
	// Experimental.
	SetEnvironment(val *map[string]*string)
	// Experimental.
	EnvironmentInput() *map[string]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfMonitoringSchedule_MonitoringJobDefinitionProperty
	// Experimental.
	SetInternalValue(val *TfMonitoringSchedule_MonitoringJobDefinitionProperty)
	// Experimental.
	MonitoringAppSpecification() TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference
	// Experimental.
	MonitoringAppSpecificationInput() *TfMonitoringSchedule_MonitoringAppSpecificationProperty
	// Experimental.
	MonitoringInputs() TfMonitoringSchedule_MonitoringInputsPropertyOutputReference
	// Experimental.
	MonitoringInputsInput() *TfMonitoringSchedule_MonitoringInputsProperty
	// Experimental.
	MonitoringOutputConfig() TfMonitoringSchedule_MonitoringOutputConfigPropertyOutputReference
	// Experimental.
	MonitoringOutputConfigInput() *TfMonitoringSchedule_MonitoringOutputConfigProperty
	// Experimental.
	MonitoringResources() TfMonitoringSchedule_MonitoringResourcesPropertyOutputReference
	// Experimental.
	MonitoringResourcesInput() *TfMonitoringSchedule_MonitoringResourcesProperty
	// Experimental.
	NetworkConfig() TfMonitoringSchedule_NetworkConfigPropertyOutputReference
	// Experimental.
	NetworkConfigInput() *TfMonitoringSchedule_NetworkConfigProperty
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	StoppingCondition() TfMonitoringSchedule_StoppingConditionPropertyList
	// Experimental.
	StoppingConditionInput() interface{}
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
	PutBaseline(value *TfMonitoringSchedule_BaselineProperty)
	// Experimental.
	PutMonitoringAppSpecification(value *TfMonitoringSchedule_MonitoringAppSpecificationProperty)
	// Experimental.
	PutMonitoringInputs(value *TfMonitoringSchedule_MonitoringInputsProperty)
	// Experimental.
	PutMonitoringOutputConfig(value *TfMonitoringSchedule_MonitoringOutputConfigProperty)
	// Experimental.
	PutMonitoringResources(value *TfMonitoringSchedule_MonitoringResourcesProperty)
	// Experimental.
	PutNetworkConfig(value *TfMonitoringSchedule_NetworkConfigProperty)
	// Experimental.
	PutStoppingCondition(value interface{})
	// Experimental.
	ResetBaseline()
	// Experimental.
	ResetEnvironment()
	// Experimental.
	ResetNetworkConfig()
	// Experimental.
	ResetStoppingCondition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference
type jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) Baseline() TfMonitoringSchedule_BaselinePropertyOutputReference {
	var returns TfMonitoringSchedule_BaselinePropertyOutputReference
	_jsii_.Get(
		j,
		"baseline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) BaselineInput() *TfMonitoringSchedule_BaselineProperty {
	var returns *TfMonitoringSchedule_BaselineProperty
	_jsii_.Get(
		j,
		"baselineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) InternalValue() *TfMonitoringSchedule_MonitoringJobDefinitionProperty {
	var returns *TfMonitoringSchedule_MonitoringJobDefinitionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringAppSpecification() TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference {
	var returns TfMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringAppSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringAppSpecificationInput() *TfMonitoringSchedule_MonitoringAppSpecificationProperty {
	var returns *TfMonitoringSchedule_MonitoringAppSpecificationProperty
	_jsii_.Get(
		j,
		"monitoringAppSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringInputs() TfMonitoringSchedule_MonitoringInputsPropertyOutputReference {
	var returns TfMonitoringSchedule_MonitoringInputsPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringInputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringInputsInput() *TfMonitoringSchedule_MonitoringInputsProperty {
	var returns *TfMonitoringSchedule_MonitoringInputsProperty
	_jsii_.Get(
		j,
		"monitoringInputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringOutputConfig() TfMonitoringSchedule_MonitoringOutputConfigPropertyOutputReference {
	var returns TfMonitoringSchedule_MonitoringOutputConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringOutputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringOutputConfigInput() *TfMonitoringSchedule_MonitoringOutputConfigProperty {
	var returns *TfMonitoringSchedule_MonitoringOutputConfigProperty
	_jsii_.Get(
		j,
		"monitoringOutputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringResources() TfMonitoringSchedule_MonitoringResourcesPropertyOutputReference {
	var returns TfMonitoringSchedule_MonitoringResourcesPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringResourcesInput() *TfMonitoringSchedule_MonitoringResourcesProperty {
	var returns *TfMonitoringSchedule_MonitoringResourcesProperty
	_jsii_.Get(
		j,
		"monitoringResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) NetworkConfig() TfMonitoringSchedule_NetworkConfigPropertyOutputReference {
	var returns TfMonitoringSchedule_NetworkConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) NetworkConfigInput() *TfMonitoringSchedule_NetworkConfigProperty {
	var returns *TfMonitoringSchedule_NetworkConfigProperty
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) StoppingCondition() TfMonitoringSchedule_StoppingConditionPropertyList {
	var returns TfMonitoringSchedule_StoppingConditionPropertyList
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) StoppingConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfMonitoringSchedule.MonitoringJobDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference_Override(t TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfMonitoringSchedule.MonitoringJobDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetInternalValue(val *TfMonitoringSchedule_MonitoringJobDefinitionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutBaseline(value *TfMonitoringSchedule_BaselineProperty) {
	if err := t.validatePutBaselineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBaseline",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutMonitoringAppSpecification(value *TfMonitoringSchedule_MonitoringAppSpecificationProperty) {
	if err := t.validatePutMonitoringAppSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMonitoringAppSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutMonitoringInputs(value *TfMonitoringSchedule_MonitoringInputsProperty) {
	if err := t.validatePutMonitoringInputsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMonitoringInputs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutMonitoringOutputConfig(value *TfMonitoringSchedule_MonitoringOutputConfigProperty) {
	if err := t.validatePutMonitoringOutputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMonitoringOutputConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutMonitoringResources(value *TfMonitoringSchedule_MonitoringResourcesProperty) {
	if err := t.validatePutMonitoringResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMonitoringResources",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutNetworkConfig(value *TfMonitoringSchedule_NetworkConfigProperty) {
	if err := t.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutStoppingCondition(value interface{}) {
	if err := t.validatePutStoppingConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStoppingCondition",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ResetBaseline() {
	_jsii_.InvokeVoid(
		t,
		"resetBaseline",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		t,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ResetStoppingCondition() {
	_jsii_.InvokeVoid(
		t,
		"resetStoppingCondition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

