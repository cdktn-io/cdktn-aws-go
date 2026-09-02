package awsiam

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsiam/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsiam/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfPolicyDocument_StatementPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Actions() *[]*string
	// Experimental.
	SetActions(val *[]*string)
	// Experimental.
	ActionsInput() *[]*string
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
	Condition() DataTfPolicyDocument_ConditionPropertyList
	// Experimental.
	ConditionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Effect() *string
	// Experimental.
	SetEffect(val *string)
	// Experimental.
	EffectInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	NotActions() *[]*string
	// Experimental.
	SetNotActions(val *[]*string)
	// Experimental.
	NotActionsInput() *[]*string
	// Experimental.
	NotPrincipals() DataTfPolicyDocument_NotPrincipalsPropertyList
	// Experimental.
	NotPrincipalsInput() interface{}
	// Experimental.
	NotResources() *[]*string
	// Experimental.
	SetNotResources(val *[]*string)
	// Experimental.
	NotResourcesInput() *[]*string
	// Experimental.
	Principals() DataTfPolicyDocument_PrincipalsPropertyList
	// Experimental.
	PrincipalsInput() interface{}
	// Experimental.
	Resources() *[]*string
	// Experimental.
	SetResources(val *[]*string)
	// Experimental.
	ResourcesInput() *[]*string
	// Experimental.
	Sid() *string
	// Experimental.
	SetSid(val *string)
	// Experimental.
	SidInput() *string
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
	PutCondition(value interface{})
	// Experimental.
	PutNotPrincipals(value interface{})
	// Experimental.
	PutPrincipals(value interface{})
	// Experimental.
	ResetActions()
	// Experimental.
	ResetCondition()
	// Experimental.
	ResetEffect()
	// Experimental.
	ResetNotActions()
	// Experimental.
	ResetNotPrincipals()
	// Experimental.
	ResetNotResources()
	// Experimental.
	ResetPrincipals()
	// Experimental.
	ResetResources()
	// Experimental.
	ResetSid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataTfPolicyDocument_StatementPropertyOutputReference
type jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) Actions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) ActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) Condition() DataTfPolicyDocument_ConditionPropertyList {
	var returns DataTfPolicyDocument_ConditionPropertyList
	_jsii_.Get(
		j,
		"condition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) ConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"conditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) Effect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) EffectInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) NotActions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) NotActionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) NotPrincipals() DataTfPolicyDocument_NotPrincipalsPropertyList {
	var returns DataTfPolicyDocument_NotPrincipalsPropertyList
	_jsii_.Get(
		j,
		"notPrincipals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) NotPrincipalsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notPrincipalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) NotResources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notResources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) NotResourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"notResourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) Principals() DataTfPolicyDocument_PrincipalsPropertyList {
	var returns DataTfPolicyDocument_PrincipalsPropertyList
	_jsii_.Get(
		j,
		"principals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) PrincipalsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"principalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) Resources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) ResourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) Sid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) SidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfPolicyDocument_StatementPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataTfPolicyDocument_StatementPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfPolicyDocument_StatementPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-iam.DataTfPolicyDocument.StatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfPolicyDocument_StatementPropertyOutputReference_Override(d DataTfPolicyDocument_StatementPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iam.DataTfPolicyDocument.StatementPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference)SetActions(val *[]*string) {
	if err := j.validateSetActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actions",
		val,
	)
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference)SetEffect(val *string) {
	if err := j.validateSetEffectParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effect",
		val,
	)
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference)SetNotActions(val *[]*string) {
	if err := j.validateSetNotActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notActions",
		val,
	)
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference)SetNotResources(val *[]*string) {
	if err := j.validateSetNotResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notResources",
		val,
	)
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference)SetResources(val *[]*string) {
	if err := j.validateSetResourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resources",
		val,
	)
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference)SetSid(val *string) {
	if err := j.validateSetSidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sid",
		val,
	)
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) PutCondition(value interface{}) {
	if err := d.validatePutConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCondition",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) PutNotPrincipals(value interface{}) {
	if err := d.validatePutNotPrincipalsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putNotPrincipals",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) PutPrincipals(value interface{}) {
	if err := d.validatePutPrincipalsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPrincipals",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) ResetActions() {
	_jsii_.InvokeVoid(
		d,
		"resetActions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) ResetCondition() {
	_jsii_.InvokeVoid(
		d,
		"resetCondition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) ResetEffect() {
	_jsii_.InvokeVoid(
		d,
		"resetEffect",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) ResetNotActions() {
	_jsii_.InvokeVoid(
		d,
		"resetNotActions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) ResetNotPrincipals() {
	_jsii_.InvokeVoid(
		d,
		"resetNotPrincipals",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) ResetNotResources() {
	_jsii_.InvokeVoid(
		d,
		"resetNotResources",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) ResetPrincipals() {
	_jsii_.InvokeVoid(
		d,
		"resetPrincipals",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) ResetResources() {
	_jsii_.InvokeVoid(
		d,
		"resetResources",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) ResetSid() {
	_jsii_.InvokeVoid(
		d,
		"resetSid",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataTfPolicyDocument_StatementPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

