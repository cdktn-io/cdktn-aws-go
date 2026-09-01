package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Baseline() AwsSagemakerMonitoringSchedule_BaselinePropertyOutputReference
	// Experimental.
	BaselineInput() *AwsSagemakerMonitoringSchedule_BaselineProperty
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
	InternalValue() *AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionProperty)
	// Experimental.
	MonitoringAppSpecification() AwsSagemakerMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference
	// Experimental.
	MonitoringAppSpecificationInput() *AwsSagemakerMonitoringSchedule_MonitoringAppSpecificationProperty
	// Experimental.
	MonitoringInputs() AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference
	// Experimental.
	MonitoringInputsInput() *AwsSagemakerMonitoringSchedule_MonitoringInputsProperty
	// Experimental.
	MonitoringOutputConfig() AwsSagemakerMonitoringSchedule_MonitoringOutputConfigPropertyOutputReference
	// Experimental.
	MonitoringOutputConfigInput() *AwsSagemakerMonitoringSchedule_MonitoringOutputConfigProperty
	// Experimental.
	MonitoringResources() AwsSagemakerMonitoringSchedule_MonitoringResourcesPropertyOutputReference
	// Experimental.
	MonitoringResourcesInput() *AwsSagemakerMonitoringSchedule_MonitoringResourcesProperty
	// Experimental.
	NetworkConfig() AwsSagemakerMonitoringSchedule_NetworkConfigPropertyOutputReference
	// Experimental.
	NetworkConfigInput() *AwsSagemakerMonitoringSchedule_NetworkConfigProperty
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	StoppingCondition() AwsSagemakerMonitoringSchedule_StoppingConditionPropertyList
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
	PutBaseline(value *AwsSagemakerMonitoringSchedule_BaselineProperty)
	// Experimental.
	PutMonitoringAppSpecification(value *AwsSagemakerMonitoringSchedule_MonitoringAppSpecificationProperty)
	// Experimental.
	PutMonitoringInputs(value *AwsSagemakerMonitoringSchedule_MonitoringInputsProperty)
	// Experimental.
	PutMonitoringOutputConfig(value *AwsSagemakerMonitoringSchedule_MonitoringOutputConfigProperty)
	// Experimental.
	PutMonitoringResources(value *AwsSagemakerMonitoringSchedule_MonitoringResourcesProperty)
	// Experimental.
	PutNetworkConfig(value *AwsSagemakerMonitoringSchedule_NetworkConfigProperty)
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

// The jsii proxy struct for AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference
type jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) Baseline() AwsSagemakerMonitoringSchedule_BaselinePropertyOutputReference {
	var returns AwsSagemakerMonitoringSchedule_BaselinePropertyOutputReference
	_jsii_.Get(
		j,
		"baseline",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) BaselineInput() *AwsSagemakerMonitoringSchedule_BaselineProperty {
	var returns *AwsSagemakerMonitoringSchedule_BaselineProperty
	_jsii_.Get(
		j,
		"baselineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) InternalValue() *AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionProperty {
	var returns *AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringAppSpecification() AwsSagemakerMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference {
	var returns AwsSagemakerMonitoringSchedule_MonitoringAppSpecificationPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringAppSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringAppSpecificationInput() *AwsSagemakerMonitoringSchedule_MonitoringAppSpecificationProperty {
	var returns *AwsSagemakerMonitoringSchedule_MonitoringAppSpecificationProperty
	_jsii_.Get(
		j,
		"monitoringAppSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringInputs() AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference {
	var returns AwsSagemakerMonitoringSchedule_MonitoringInputsPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringInputs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringInputsInput() *AwsSagemakerMonitoringSchedule_MonitoringInputsProperty {
	var returns *AwsSagemakerMonitoringSchedule_MonitoringInputsProperty
	_jsii_.Get(
		j,
		"monitoringInputsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringOutputConfig() AwsSagemakerMonitoringSchedule_MonitoringOutputConfigPropertyOutputReference {
	var returns AwsSagemakerMonitoringSchedule_MonitoringOutputConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringOutputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringOutputConfigInput() *AwsSagemakerMonitoringSchedule_MonitoringOutputConfigProperty {
	var returns *AwsSagemakerMonitoringSchedule_MonitoringOutputConfigProperty
	_jsii_.Get(
		j,
		"monitoringOutputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringResources() AwsSagemakerMonitoringSchedule_MonitoringResourcesPropertyOutputReference {
	var returns AwsSagemakerMonitoringSchedule_MonitoringResourcesPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) MonitoringResourcesInput() *AwsSagemakerMonitoringSchedule_MonitoringResourcesProperty {
	var returns *AwsSagemakerMonitoringSchedule_MonitoringResourcesProperty
	_jsii_.Get(
		j,
		"monitoringResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) NetworkConfig() AwsSagemakerMonitoringSchedule_NetworkConfigPropertyOutputReference {
	var returns AwsSagemakerMonitoringSchedule_NetworkConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"networkConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) NetworkConfigInput() *AwsSagemakerMonitoringSchedule_NetworkConfigProperty {
	var returns *AwsSagemakerMonitoringSchedule_NetworkConfigProperty
	_jsii_.Get(
		j,
		"networkConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) StoppingCondition() AwsSagemakerMonitoringSchedule_StoppingConditionPropertyList {
	var returns AwsSagemakerMonitoringSchedule_StoppingConditionPropertyList
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) StoppingConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerMonitoringSchedule.MonitoringJobDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference_Override(a AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerMonitoringSchedule.MonitoringJobDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetInternalValue(val *AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutBaseline(value *AwsSagemakerMonitoringSchedule_BaselineProperty) {
	if err := a.validatePutBaselineParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBaseline",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutMonitoringAppSpecification(value *AwsSagemakerMonitoringSchedule_MonitoringAppSpecificationProperty) {
	if err := a.validatePutMonitoringAppSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonitoringAppSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutMonitoringInputs(value *AwsSagemakerMonitoringSchedule_MonitoringInputsProperty) {
	if err := a.validatePutMonitoringInputsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonitoringInputs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutMonitoringOutputConfig(value *AwsSagemakerMonitoringSchedule_MonitoringOutputConfigProperty) {
	if err := a.validatePutMonitoringOutputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonitoringOutputConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutMonitoringResources(value *AwsSagemakerMonitoringSchedule_MonitoringResourcesProperty) {
	if err := a.validatePutMonitoringResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonitoringResources",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutNetworkConfig(value *AwsSagemakerMonitoringSchedule_NetworkConfigProperty) {
	if err := a.validatePutNetworkConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) PutStoppingCondition(value interface{}) {
	if err := a.validatePutStoppingConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStoppingCondition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ResetBaseline() {
	_jsii_.InvokeVoid(
		a,
		"resetBaseline",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ResetNetworkConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ResetStoppingCondition() {
	_jsii_.InvokeVoid(
		a,
		"resetStoppingCondition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

