package awsarcregionswitch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsarcregionswitch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsarcregionswitch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CapacityMonitoringApproach() *string
	// Experimental.
	SetCapacityMonitoringApproach(val *string)
	// Experimental.
	CapacityMonitoringApproachInput() *string
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
	EksClusters() AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigEksClustersPropertyList
	// Experimental.
	EksClustersInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	KubernetesResourceType() AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigKubernetesResourceTypePropertyList
	// Experimental.
	KubernetesResourceTypeInput() interface{}
	// Experimental.
	ScalingResources() AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigScalingResourcesPropertyList
	// Experimental.
	ScalingResourcesInput() interface{}
	// Experimental.
	TargetPercent() *float64
	// Experimental.
	SetTargetPercent(val *float64)
	// Experimental.
	TargetPercentInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeoutMinutes() *float64
	// Experimental.
	SetTimeoutMinutes(val *float64)
	// Experimental.
	TimeoutMinutesInput() *float64
	// Experimental.
	Ungraceful() AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigUngracefulPropertyList
	// Experimental.
	UngracefulInput() interface{}
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
	PutEksClusters(value interface{})
	// Experimental.
	PutKubernetesResourceType(value interface{})
	// Experimental.
	PutScalingResources(value interface{})
	// Experimental.
	PutUngraceful(value interface{})
	// Experimental.
	ResetEksClusters()
	// Experimental.
	ResetKubernetesResourceType()
	// Experimental.
	ResetScalingResources()
	// Experimental.
	ResetTimeoutMinutes()
	// Experimental.
	ResetUngraceful()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference
type jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) CapacityMonitoringApproach() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMonitoringApproach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) CapacityMonitoringApproachInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMonitoringApproachInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) EksClusters() AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigEksClustersPropertyList {
	var returns AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigEksClustersPropertyList
	_jsii_.Get(
		j,
		"eksClusters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) EksClustersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"eksClustersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) KubernetesResourceType() AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigKubernetesResourceTypePropertyList {
	var returns AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigKubernetesResourceTypePropertyList
	_jsii_.Get(
		j,
		"kubernetesResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) KubernetesResourceTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kubernetesResourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) ScalingResources() AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigScalingResourcesPropertyList {
	var returns AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigScalingResourcesPropertyList
	_jsii_.Get(
		j,
		"scalingResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) ScalingResourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) TargetPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) TargetPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) TimeoutMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) TimeoutMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) Ungraceful() AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigUngracefulPropertyList {
	var returns AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigUngracefulPropertyList
	_jsii_.Get(
		j,
		"ungraceful",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) UngracefulInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ungracefulInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.AwsArcregionswitchPlan.WorkflowStepEksResourceScalingConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference_Override(a AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.AwsArcregionswitchPlan.WorkflowStepEksResourceScalingConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference)SetCapacityMonitoringApproach(val *string) {
	if err := j.validateSetCapacityMonitoringApproachParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityMonitoringApproach",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference)SetTargetPercent(val *float64) {
	if err := j.validateSetTargetPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetPercent",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference)SetTimeoutMinutes(val *float64) {
	if err := j.validateSetTimeoutMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutMinutes",
		val,
	)
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) PutEksClusters(value interface{}) {
	if err := a.validatePutEksClustersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEksClusters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) PutKubernetesResourceType(value interface{}) {
	if err := a.validatePutKubernetesResourceTypeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKubernetesResourceType",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) PutScalingResources(value interface{}) {
	if err := a.validatePutScalingResourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScalingResources",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) PutUngraceful(value interface{}) {
	if err := a.validatePutUngracefulParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUngraceful",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) ResetEksClusters() {
	_jsii_.InvokeVoid(
		a,
		"resetEksClusters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) ResetKubernetesResourceType() {
	_jsii_.InvokeVoid(
		a,
		"resetKubernetesResourceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) ResetScalingResources() {
	_jsii_.InvokeVoid(
		a,
		"resetScalingResources",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) ResetTimeoutMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) ResetUngraceful() {
	_jsii_.InvokeVoid(
		a,
		"resetUngraceful",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsArcregionswitchPlan_WorkflowStepEksResourceScalingConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

