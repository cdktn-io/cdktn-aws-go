package awsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmsk/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsmsk/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AzDistribution() *string
	// Experimental.
	SetAzDistribution(val *string)
	// Experimental.
	AzDistributionInput() *string
	// Experimental.
	ClientSubnets() *[]*string
	// Experimental.
	SetClientSubnets(val *[]*string)
	// Experimental.
	ClientSubnetsInput() *[]*string
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
	ConnectivityInfo() AwsMskCluster_ConnectivityInfoPropertyOutputReference
	// Experimental.
	ConnectivityInfoInput() *AwsMskCluster_ConnectivityInfoProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InstanceType() *string
	// Experimental.
	SetInstanceType(val *string)
	// Experimental.
	InstanceTypeInput() *string
	// Experimental.
	InternalValue() *AwsMskCluster_BrokerNodeGroupInfoProperty
	// Experimental.
	SetInternalValue(val *AwsMskCluster_BrokerNodeGroupInfoProperty)
	// Experimental.
	SecurityGroups() *[]*string
	// Experimental.
	SetSecurityGroups(val *[]*string)
	// Experimental.
	SecurityGroupsInput() *[]*string
	// Experimental.
	StorageInfo() AwsMskCluster_StorageInfoPropertyOutputReference
	// Experimental.
	StorageInfoInput() *AwsMskCluster_StorageInfoProperty
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
	PutConnectivityInfo(value *AwsMskCluster_ConnectivityInfoProperty)
	// Experimental.
	PutStorageInfo(value *AwsMskCluster_StorageInfoProperty)
	// Experimental.
	ResetAzDistribution()
	// Experimental.
	ResetConnectivityInfo()
	// Experimental.
	ResetStorageInfo()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference
type jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) AzDistribution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) AzDistributionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azDistributionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) ClientSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) ClientSubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientSubnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) ConnectivityInfo() AwsMskCluster_ConnectivityInfoPropertyOutputReference {
	var returns AwsMskCluster_ConnectivityInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"connectivityInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) ConnectivityInfoInput() *AwsMskCluster_ConnectivityInfoProperty {
	var returns *AwsMskCluster_ConnectivityInfoProperty
	_jsii_.Get(
		j,
		"connectivityInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) InternalValue() *AwsMskCluster_BrokerNodeGroupInfoProperty {
	var returns *AwsMskCluster_BrokerNodeGroupInfoProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) StorageInfo() AwsMskCluster_StorageInfoPropertyOutputReference {
	var returns AwsMskCluster_StorageInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"storageInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) StorageInfoInput() *AwsMskCluster_StorageInfoProperty {
	var returns *AwsMskCluster_StorageInfoProperty
	_jsii_.Get(
		j,
		"storageInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMskCluster_BrokerNodeGroupInfoPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-msk.AwsMskCluster.BrokerNodeGroupInfoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference_Override(a AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-msk.AwsMskCluster.BrokerNodeGroupInfoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference)SetAzDistribution(val *string) {
	if err := j.validateSetAzDistributionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azDistribution",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference)SetClientSubnets(val *[]*string) {
	if err := j.validateSetClientSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSubnets",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference)SetInternalValue(val *AwsMskCluster_BrokerNodeGroupInfoProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) PutConnectivityInfo(value *AwsMskCluster_ConnectivityInfoProperty) {
	if err := a.validatePutConnectivityInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConnectivityInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) PutStorageInfo(value *AwsMskCluster_StorageInfoProperty) {
	if err := a.validatePutStorageInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStorageInfo",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) ResetAzDistribution() {
	_jsii_.InvokeVoid(
		a,
		"resetAzDistribution",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) ResetConnectivityInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectivityInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) ResetStorageInfo() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageInfo",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMskCluster_BrokerNodeGroupInfoPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

