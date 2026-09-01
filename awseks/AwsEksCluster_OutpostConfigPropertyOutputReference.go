package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseks/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEksCluster_OutpostConfigPropertyOutputReference interface {
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
	// Experimental.
	ControlPlaneInstanceType() *string
	// Experimental.
	SetControlPlaneInstanceType(val *string)
	// Experimental.
	ControlPlaneInstanceTypeInput() *string
	// Experimental.
	ControlPlanePlacement() AwsEksCluster_ControlPlanePlacementPropertyOutputReference
	// Experimental.
	ControlPlanePlacementInput() *AwsEksCluster_ControlPlanePlacementProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EtcdInstanceType() *string
	// Experimental.
	SetEtcdInstanceType(val *string)
	// Experimental.
	EtcdInstanceTypeInput() *string
	// Experimental.
	EtcdPlacement() AwsEksCluster_EtcdPlacementPropertyOutputReference
	// Experimental.
	EtcdPlacementInput() *AwsEksCluster_EtcdPlacementProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsEksCluster_OutpostConfigProperty
	// Experimental.
	SetInternalValue(val *AwsEksCluster_OutpostConfigProperty)
	// Experimental.
	OutpostArns() *[]*string
	// Experimental.
	SetOutpostArns(val *[]*string)
	// Experimental.
	OutpostArnsInput() *[]*string
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
	PutControlPlanePlacement(value *AwsEksCluster_ControlPlanePlacementProperty)
	// Experimental.
	PutEtcdPlacement(value *AwsEksCluster_EtcdPlacementProperty)
	// Experimental.
	ResetControlPlanePlacement()
	// Experimental.
	ResetEtcdInstanceType()
	// Experimental.
	ResetEtcdPlacement()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEksCluster_OutpostConfigPropertyOutputReference
type jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) ControlPlaneInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) ControlPlaneInstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPlaneInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) ControlPlanePlacement() AwsEksCluster_ControlPlanePlacementPropertyOutputReference {
	var returns AwsEksCluster_ControlPlanePlacementPropertyOutputReference
	_jsii_.Get(
		j,
		"controlPlanePlacement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) ControlPlanePlacementInput() *AwsEksCluster_ControlPlanePlacementProperty {
	var returns *AwsEksCluster_ControlPlanePlacementProperty
	_jsii_.Get(
		j,
		"controlPlanePlacementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) EtcdInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etcdInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) EtcdInstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etcdInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) EtcdPlacement() AwsEksCluster_EtcdPlacementPropertyOutputReference {
	var returns AwsEksCluster_EtcdPlacementPropertyOutputReference
	_jsii_.Get(
		j,
		"etcdPlacement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) EtcdPlacementInput() *AwsEksCluster_EtcdPlacementProperty {
	var returns *AwsEksCluster_EtcdPlacementProperty
	_jsii_.Get(
		j,
		"etcdPlacementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) InternalValue() *AwsEksCluster_OutpostConfigProperty {
	var returns *AwsEksCluster_OutpostConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) OutpostArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outpostArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) OutpostArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outpostArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEksCluster_OutpostConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEksCluster_OutpostConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEksCluster_OutpostConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eks.AwsEksCluster.OutpostConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEksCluster_OutpostConfigPropertyOutputReference_Override(a AwsEksCluster_OutpostConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.AwsEksCluster.OutpostConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference)SetControlPlaneInstanceType(val *string) {
	if err := j.validateSetControlPlaneInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controlPlaneInstanceType",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference)SetEtcdInstanceType(val *string) {
	if err := j.validateSetEtcdInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"etcdInstanceType",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference)SetInternalValue(val *AwsEksCluster_OutpostConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference)SetOutpostArns(val *[]*string) {
	if err := j.validateSetOutpostArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outpostArns",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) PutControlPlanePlacement(value *AwsEksCluster_ControlPlanePlacementProperty) {
	if err := a.validatePutControlPlanePlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putControlPlanePlacement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) PutEtcdPlacement(value *AwsEksCluster_EtcdPlacementProperty) {
	if err := a.validatePutEtcdPlacementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEtcdPlacement",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) ResetControlPlanePlacement() {
	_jsii_.InvokeVoid(
		a,
		"resetControlPlanePlacement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) ResetEtcdInstanceType() {
	_jsii_.InvokeVoid(
		a,
		"resetEtcdInstanceType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) ResetEtcdPlacement() {
	_jsii_.InvokeVoid(
		a,
		"resetEtcdPlacement",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEksCluster_OutpostConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

