package awsssoadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsssoadmin/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsssoadmin/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference interface {
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
	CustomerManagedPolicyReference() TfPermissionsBoundaryAttachment_CustomerManagedPolicyReferencePropertyOutputReference
	// Experimental.
	CustomerManagedPolicyReferenceInput() *TfPermissionsBoundaryAttachment_CustomerManagedPolicyReferenceProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfPermissionsBoundaryAttachment_PermissionsBoundaryProperty
	// Experimental.
	SetInternalValue(val *TfPermissionsBoundaryAttachment_PermissionsBoundaryProperty)
	// Experimental.
	ManagedPolicyArn() *string
	// Experimental.
	SetManagedPolicyArn(val *string)
	// Experimental.
	ManagedPolicyArnInput() *string
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
	PutCustomerManagedPolicyReference(value *TfPermissionsBoundaryAttachment_CustomerManagedPolicyReferenceProperty)
	// Experimental.
	ResetCustomerManagedPolicyReference()
	// Experimental.
	ResetManagedPolicyArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference
type jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) CustomerManagedPolicyReference() TfPermissionsBoundaryAttachment_CustomerManagedPolicyReferencePropertyOutputReference {
	var returns TfPermissionsBoundaryAttachment_CustomerManagedPolicyReferencePropertyOutputReference
	_jsii_.Get(
		j,
		"customerManagedPolicyReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) CustomerManagedPolicyReferenceInput() *TfPermissionsBoundaryAttachment_CustomerManagedPolicyReferenceProperty {
	var returns *TfPermissionsBoundaryAttachment_CustomerManagedPolicyReferenceProperty
	_jsii_.Get(
		j,
		"customerManagedPolicyReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) InternalValue() *TfPermissionsBoundaryAttachment_PermissionsBoundaryProperty {
	var returns *TfPermissionsBoundaryAttachment_PermissionsBoundaryProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) ManagedPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPolicyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) ManagedPolicyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPolicyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sso-admin.TfPermissionsBoundaryAttachment.PermissionsBoundaryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference_Override(t TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sso-admin.TfPermissionsBoundaryAttachment.PermissionsBoundaryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference)SetInternalValue(val *TfPermissionsBoundaryAttachment_PermissionsBoundaryProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference)SetManagedPolicyArn(val *string) {
	if err := j.validateSetManagedPolicyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedPolicyArn",
		val,
	)
}

func (j *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) PutCustomerManagedPolicyReference(value *TfPermissionsBoundaryAttachment_CustomerManagedPolicyReferenceProperty) {
	if err := t.validatePutCustomerManagedPolicyReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomerManagedPolicyReference",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) ResetCustomerManagedPolicyReference() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomerManagedPolicyReference",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) ResetManagedPolicyArn() {
	_jsii_.InvokeVoid(
		t,
		"resetManagedPolicyArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

