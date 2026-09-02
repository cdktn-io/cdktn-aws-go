package awsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmsk/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsmsk/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCluster_BrokerNodeGroupInfoPropertyOutputReference interface {
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
	ConnectivityInfo() TfCluster_ConnectivityInfoPropertyOutputReference
	// Experimental.
	ConnectivityInfoInput() *TfCluster_ConnectivityInfoProperty
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
	InternalValue() *TfCluster_BrokerNodeGroupInfoProperty
	// Experimental.
	SetInternalValue(val *TfCluster_BrokerNodeGroupInfoProperty)
	// Experimental.
	SecurityGroups() *[]*string
	// Experimental.
	SetSecurityGroups(val *[]*string)
	// Experimental.
	SecurityGroupsInput() *[]*string
	// Experimental.
	StorageInfo() TfCluster_StorageInfoPropertyOutputReference
	// Experimental.
	StorageInfoInput() *TfCluster_StorageInfoProperty
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
	PutConnectivityInfo(value *TfCluster_ConnectivityInfoProperty)
	// Experimental.
	PutStorageInfo(value *TfCluster_StorageInfoProperty)
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

// The jsii proxy struct for TfCluster_BrokerNodeGroupInfoPropertyOutputReference
type jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) AzDistribution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azDistribution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) AzDistributionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azDistributionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) ClientSubnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientSubnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) ClientSubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clientSubnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) ConnectivityInfo() TfCluster_ConnectivityInfoPropertyOutputReference {
	var returns TfCluster_ConnectivityInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"connectivityInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) ConnectivityInfoInput() *TfCluster_ConnectivityInfoProperty {
	var returns *TfCluster_ConnectivityInfoProperty
	_jsii_.Get(
		j,
		"connectivityInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) InternalValue() *TfCluster_BrokerNodeGroupInfoProperty {
	var returns *TfCluster_BrokerNodeGroupInfoProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) StorageInfo() TfCluster_StorageInfoPropertyOutputReference {
	var returns TfCluster_StorageInfoPropertyOutputReference
	_jsii_.Get(
		j,
		"storageInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) StorageInfoInput() *TfCluster_StorageInfoProperty {
	var returns *TfCluster_StorageInfoProperty
	_jsii_.Get(
		j,
		"storageInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCluster_BrokerNodeGroupInfoPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCluster_BrokerNodeGroupInfoPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCluster_BrokerNodeGroupInfoPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-msk.TfCluster.BrokerNodeGroupInfoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCluster_BrokerNodeGroupInfoPropertyOutputReference_Override(t TfCluster_BrokerNodeGroupInfoPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-msk.TfCluster.BrokerNodeGroupInfoPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference)SetAzDistribution(val *string) {
	if err := j.validateSetAzDistributionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azDistribution",
		val,
	)
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference)SetClientSubnets(val *[]*string) {
	if err := j.validateSetClientSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSubnets",
		val,
	)
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference)SetInternalValue(val *TfCluster_BrokerNodeGroupInfoProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) PutConnectivityInfo(value *TfCluster_ConnectivityInfoProperty) {
	if err := t.validatePutConnectivityInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConnectivityInfo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) PutStorageInfo(value *TfCluster_StorageInfoProperty) {
	if err := t.validatePutStorageInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStorageInfo",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) ResetAzDistribution() {
	_jsii_.InvokeVoid(
		t,
		"resetAzDistribution",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) ResetConnectivityInfo() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectivityInfo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) ResetStorageInfo() {
	_jsii_.InvokeVoid(
		t,
		"resetStorageInfo",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCluster_BrokerNodeGroupInfoPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

