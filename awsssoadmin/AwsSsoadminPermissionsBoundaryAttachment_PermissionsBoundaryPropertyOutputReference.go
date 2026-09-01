package awsssoadmin

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsssoadmin/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsssoadmin/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference interface {
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
	CustomerManagedPolicyReference() AwsSsoadminPermissionsBoundaryAttachment_CustomerManagedPolicyReferencePropertyOutputReference
	// Experimental.
	CustomerManagedPolicyReferenceInput() *AwsSsoadminPermissionsBoundaryAttachment_CustomerManagedPolicyReferenceProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryProperty
	// Experimental.
	SetInternalValue(val *AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryProperty)
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
	PutCustomerManagedPolicyReference(value *AwsSsoadminPermissionsBoundaryAttachment_CustomerManagedPolicyReferenceProperty)
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

// The jsii proxy struct for AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference
type jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) CustomerManagedPolicyReference() AwsSsoadminPermissionsBoundaryAttachment_CustomerManagedPolicyReferencePropertyOutputReference {
	var returns AwsSsoadminPermissionsBoundaryAttachment_CustomerManagedPolicyReferencePropertyOutputReference
	_jsii_.Get(
		j,
		"customerManagedPolicyReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) CustomerManagedPolicyReferenceInput() *AwsSsoadminPermissionsBoundaryAttachment_CustomerManagedPolicyReferenceProperty {
	var returns *AwsSsoadminPermissionsBoundaryAttachment_CustomerManagedPolicyReferenceProperty
	_jsii_.Get(
		j,
		"customerManagedPolicyReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) InternalValue() *AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryProperty {
	var returns *AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) ManagedPolicyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPolicyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) ManagedPolicyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPolicyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sso-admin.AwsSsoadminPermissionsBoundaryAttachment.PermissionsBoundaryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference_Override(a AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sso-admin.AwsSsoadminPermissionsBoundaryAttachment.PermissionsBoundaryPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference)SetInternalValue(val *AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference)SetManagedPolicyArn(val *string) {
	if err := j.validateSetManagedPolicyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedPolicyArn",
		val,
	)
}

func (j *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) PutCustomerManagedPolicyReference(value *AwsSsoadminPermissionsBoundaryAttachment_CustomerManagedPolicyReferenceProperty) {
	if err := a.validatePutCustomerManagedPolicyReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomerManagedPolicyReference",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) ResetCustomerManagedPolicyReference() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerManagedPolicyReference",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) ResetManagedPolicyArn() {
	_jsii_.InvokeVoid(
		a,
		"resetManagedPolicyArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSsoadminPermissionsBoundaryAttachment_PermissionsBoundaryPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

