package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsnetworkmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsnetworkmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AddToNetworkFunctionGroup() *string
	// Experimental.
	SetAddToNetworkFunctionGroup(val *string)
	// Experimental.
	AddToNetworkFunctionGroupInput() *string
	// Experimental.
	AssociationMethod() *string
	// Experimental.
	SetAssociationMethod(val *string)
	// Experimental.
	AssociationMethodInput() *string
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
	InternalValue() *DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionProperty
	// Experimental.
	SetInternalValue(val *DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionProperty)
	// Experimental.
	RequireAcceptance() interface{}
	// Experimental.
	SetRequireAcceptance(val interface{})
	// Experimental.
	RequireAcceptanceInput() interface{}
	// Experimental.
	Segment() *string
	// Experimental.
	SetSegment(val *string)
	// Experimental.
	SegmentInput() *string
	// Experimental.
	TagValueOfKey() *string
	// Experimental.
	SetTagValueOfKey(val *string)
	// Experimental.
	TagValueOfKeyInput() *string
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
	ResetAddToNetworkFunctionGroup()
	// Experimental.
	ResetAssociationMethod()
	// Experimental.
	ResetRequireAcceptance()
	// Experimental.
	ResetSegment()
	// Experimental.
	ResetTagValueOfKey()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference
type jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) AddToNetworkFunctionGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addToNetworkFunctionGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) AddToNetworkFunctionGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addToNetworkFunctionGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) AssociationMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) AssociationMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"associationMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) InternalValue() *DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionProperty {
	var returns *DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) RequireAcceptance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAcceptance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) RequireAcceptanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireAcceptanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) Segment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) SegmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"segmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) TagValueOfKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValueOfKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) TagValueOfKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagValueOfKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-manager.DataTfCoreNetworkPolicyDocument.AttachmentPoliciesActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference_Override(d DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-manager.DataTfCoreNetworkPolicyDocument.AttachmentPoliciesActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference)SetAddToNetworkFunctionGroup(val *string) {
	if err := j.validateSetAddToNetworkFunctionGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addToNetworkFunctionGroup",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference)SetAssociationMethod(val *string) {
	if err := j.validateSetAssociationMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"associationMethod",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference)SetInternalValue(val *DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference)SetRequireAcceptance(val interface{}) {
	if err := j.validateSetRequireAcceptanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireAcceptance",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference)SetSegment(val *string) {
	if err := j.validateSetSegmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segment",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference)SetTagValueOfKey(val *string) {
	if err := j.validateSetTagValueOfKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagValueOfKey",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) ResetAddToNetworkFunctionGroup() {
	_jsii_.InvokeVoid(
		d,
		"resetAddToNetworkFunctionGroup",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) ResetAssociationMethod() {
	_jsii_.InvokeVoid(
		d,
		"resetAssociationMethod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) ResetRequireAcceptance() {
	_jsii_.InvokeVoid(
		d,
		"resetRequireAcceptance",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) ResetSegment() {
	_jsii_.InvokeVoid(
		d,
		"resetSegment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) ResetTagValueOfKey() {
	_jsii_.InvokeVoid(
		d,
		"resetTagValueOfKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentPoliciesActionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

