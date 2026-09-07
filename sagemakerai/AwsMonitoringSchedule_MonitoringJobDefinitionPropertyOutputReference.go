package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Baseline() AwsMonitoringSchedule_BaselinePropertyOutputReference
	// Experimental.
	BaselineInput() *AwsMonitoringSchedule_BaselineProperty
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
	InternalValue() *AwsMonitoringSchedule_MonitoringJobDefinitionProperty
	// Experimental.
	SetInternalValue(val *AwsMonitoringSchedule_MonitoringJobDefinitionProperty)
	// Experimental.
	MonitoringAppSpecification() AwsMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference
	// Experimental.
	MonitoringAppSpecificationInput() *AwsMonitoringSchedule_MonitoringAppSpecificationProperty
	// Experimental.
	MonitoringInputs() AwsMonitoringSchedule_MonitoringInputsPropertyOutputReference
	// Experimental.
	MonitoringInputsInput() *AwsMonitoringSchedule_MonitoringInputsProperty
	// Experimental.
	MonitoringOutputConfig() AwsMonitoringSchedule_MonitoringOutputConfigPropertyOutputReference
	// Experimental.
	MonitoringOutputConfigInput() *AwsMonitoringSchedule_MonitoringOutputConfigProperty
	// Experimental.
	MonitoringResources() AwsMonitoringSchedule_MonitoringResourcesPropertyOutputReference
	// Experimental.
	MonitoringResourcesInput() *AwsMonitoringSchedule_MonitoringResourcesProperty
	// Experimental.
	NetworkConfig() AwsMonitoringSchedule_NetworkConfigPropertyOutputReference
	// Experimental.
	NetworkConfigInput() *AwsMonitoringSchedule_NetworkConfigProperty
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	StoppingCondition() AwsMonitoringSchedule_StoppingConditionPropertyList
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
	PutBaseline(value *AwsMonitoringSchedule_BaselineProperty)
	// Experimental.
	PutMonitoringAppSpecification(value *AwsMonitoringSchedule_MonitoringAppSpecificationProperty)
	// Experimental.
	PutMonitoringInputs(value *AwsMonitoringSchedule_MonitoringInputsProperty)
	// Experimental.
	PutMonitoringOutputConfig(value *AwsMonitoringSchedule_MonitoringOutputConfigProperty)
	// Experimental.
	PutMonitoringResources(value *AwsMonitoringSchedule_MonitoringResourcesProperty)
	// Experimental.
	PutNetworkConfig(value *AwsMonitoringSchedule_NetworkConfigProperty)
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

// The jsii proxy struct for AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference
type jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) Baseline() AwsMonitoringSchedule_BaselinePropertyOutputReference {
	var returns AwsMonitoringSchedule_BaselinePropertyOutputReference
	_jsii_.Get(
		j,
		"baseline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) BaselineInput() *AwsMonitoringSchedule_BaselineProperty {
	var returns *AwsMonitoringSchedule_BaselineProperty
	_jsii_.Get(
		j,
		"baselineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) InternalValue() *AwsMonitoringSchedule_MonitoringJobDefinitionProperty {
	var returns *AwsMonitoringSchedule_MonitoringJobDefinitionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringAppSpecification() AwsMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference {
	var returns AwsMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringAppSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringAppSpecificationInput() *AwsMonitoringSchedule_MonitoringAppSpecificationProperty {
	var returns *AwsMonitoringSchedule_MonitoringAppSpecificationProperty
	_jsii_.Get(
		j,
		"monitoringAppSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringInputs() AwsMonitoringSchedule_MonitoringInputsPropertyOutputReference {
	var returns AwsMonitoringSchedule_MonitoringInputsPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringInputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringInputsInput() *AwsMonitoringSchedule_MonitoringInputsProperty {
	var returns *AwsMonitoringSchedule_MonitoringInputsProperty
	_jsii_.Get(
		j,
		"monitoringInputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringOutputConfig() AwsMonitoringSchedule_MonitoringOutputConfigPropertyOutputReference {
	var returns AwsMonitoringSchedule_MonitoringOutputConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringOutputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringOutputConfigInput() *AwsMonitoringSchedule_MonitoringOutputConfigProperty {
	var returns *AwsMonitoringSchedule_MonitoringOutputConfigProperty
	_jsii_.Get(
		j,
		"monitoringOutputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringResources() AwsMonitoringSchedule_MonitoringResourcesPropertyOutputReference {
	var returns AwsMonitoringSchedule_MonitoringResourcesPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringResourcesInput() *AwsMonitoringSchedule_MonitoringResourcesProperty {
	var returns *AwsMonitoringSchedule_MonitoringResourcesProperty
	_jsii_.Get(
		j,
		"monitoringResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) NetworkConfig() AwsMonitoringSchedule_NetworkConfigPropertyOutputReference {
	var returns AwsMonitoringSchedule_NetworkConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) NetworkConfigInput() *AwsMonitoringSchedule_NetworkConfigProperty {
	var returns *AwsMonitoringSchedule_NetworkConfigProperty
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) StoppingCondition() AwsMonitoringSchedule_StoppingConditionPropertyList {
	var returns AwsMonitoringSchedule_StoppingConditionPropertyList
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) StoppingConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsMonitoringSchedule.MonitoringJobDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference_Override(a AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsMonitoringSchedule.MonitoringJobDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetInternalValue(val *AwsMonitoringSchedule_MonitoringJobDefinitionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutBaseline(value *AwsMonitoringSchedule_BaselineProperty) {
	if err := a.validatePutBaselineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBaseline",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutMonitoringAppSpecification(value *AwsMonitoringSchedule_MonitoringAppSpecificationProperty) {
	if err := a.validatePutMonitoringAppSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonitoringAppSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutMonitoringInputs(value *AwsMonitoringSchedule_MonitoringInputsProperty) {
	if err := a.validatePutMonitoringInputsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonitoringInputs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutMonitoringOutputConfig(value *AwsMonitoringSchedule_MonitoringOutputConfigProperty) {
	if err := a.validatePutMonitoringOutputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonitoringOutputConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutMonitoringResources(value *AwsMonitoringSchedule_MonitoringResourcesProperty) {
	if err := a.validatePutMonitoringResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonitoringResources",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutNetworkConfig(value *AwsMonitoringSchedule_NetworkConfigProperty) {
	if err := a.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutStoppingCondition(value interface{}) {
	if err := a.validatePutStoppingConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStoppingCondition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ResetBaseline() {
	_jsii_.InvokeVoid(
		a,
		"resetBaseline",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ResetStoppingCondition() {
	_jsii_.InvokeVoid(
		a,
		"resetStoppingCondition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

