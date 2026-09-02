package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsnetworkmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsnetworkmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() *DataTfCoreNetworkPolicyDocument_ViaProperty
	// Experimental.
	SetInternalValue(val *DataTfCoreNetworkPolicyDocument_ViaProperty)
	// Experimental.
	NetworkFunctionGroups() *[]*string
	// Experimental.
	SetNetworkFunctionGroups(val *[]*string)
	// Experimental.
	NetworkFunctionGroupsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WithEdgeOverride() DataTfCoreNetworkPolicyDocument_WithEdgeOverridePropertyList
	// Experimental.
	WithEdgeOverrideInput() interface{}
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
	PutWithEdgeOverride(value interface{})
	// Experimental.
	ResetNetworkFunctionGroups()
	// Experimental.
	ResetWithEdgeOverride()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference
type jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) InternalValue() *DataTfCoreNetworkPolicyDocument_ViaProperty {
	var returns *DataTfCoreNetworkPolicyDocument_ViaProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) NetworkFunctionGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkFunctionGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) NetworkFunctionGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkFunctionGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) WithEdgeOverride() DataTfCoreNetworkPolicyDocument_WithEdgeOverridePropertyList {
	var returns DataTfCoreNetworkPolicyDocument_WithEdgeOverridePropertyList
	_jsii_.Get(
		j,
		"withEdgeOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) WithEdgeOverrideInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withEdgeOverrideInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfCoreNetworkPolicyDocument_ViaPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-manager.DataTfCoreNetworkPolicyDocument.ViaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference_Override(d DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-manager.DataTfCoreNetworkPolicyDocument.ViaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference)SetInternalValue(val *DataTfCoreNetworkPolicyDocument_ViaProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference)SetNetworkFunctionGroups(val *[]*string) {
	if err := j.validateSetNetworkFunctionGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkFunctionGroups",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) PutWithEdgeOverride(value interface{}) {
	if err := d.validatePutWithEdgeOverrideParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putWithEdgeOverride",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) ResetNetworkFunctionGroups() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkFunctionGroups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) ResetWithEdgeOverride() {
	_jsii_.InvokeVoid(
		d,
		"resetWithEdgeOverride",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_ViaPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

