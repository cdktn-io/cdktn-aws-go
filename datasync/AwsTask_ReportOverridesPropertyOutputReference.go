package datasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/datasync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/datasync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTask_ReportOverridesPropertyOutputReference interface {
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
	DeletedOverride() *string
	// Experimental.
	SetDeletedOverride(val *string)
	// Experimental.
	DeletedOverrideInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsTask_ReportOverridesProperty
	// Experimental.
	SetInternalValue(val *AwsTask_ReportOverridesProperty)
	// Experimental.
	SkippedOverride() *string
	// Experimental.
	SetSkippedOverride(val *string)
	// Experimental.
	SkippedOverrideInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TransferredOverride() *string
	// Experimental.
	SetTransferredOverride(val *string)
	// Experimental.
	TransferredOverrideInput() *string
	// Experimental.
	VerifiedOverride() *string
	// Experimental.
	SetVerifiedOverride(val *string)
	// Experimental.
	VerifiedOverrideInput() *string
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
	ResetDeletedOverride()
	// Experimental.
	ResetSkippedOverride()
	// Experimental.
	ResetTransferredOverride()
	// Experimental.
	ResetVerifiedOverride()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsTask_ReportOverridesPropertyOutputReference
type jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) DeletedOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletedOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) DeletedOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deletedOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) InternalValue() *AwsTask_ReportOverridesProperty {
	var returns *AwsTask_ReportOverridesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) SkippedOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skippedOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) SkippedOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skippedOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) TransferredOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transferredOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) TransferredOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transferredOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) VerifiedOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) VerifiedOverrideInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedOverrideInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTask_ReportOverridesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsTask_ReportOverridesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTask_ReportOverridesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-datasync.AwsTask.ReportOverridesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTask_ReportOverridesPropertyOutputReference_Override(a AwsTask_ReportOverridesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-datasync.AwsTask.ReportOverridesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference)SetDeletedOverride(val *string) {
	if err := j.validateSetDeletedOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletedOverride",
		val,
	)
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference)SetInternalValue(val *AwsTask_ReportOverridesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference)SetSkippedOverride(val *string) {
	if err := j.validateSetSkippedOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skippedOverride",
		val,
	)
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference)SetTransferredOverride(val *string) {
	if err := j.validateSetTransferredOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transferredOverride",
		val,
	)
}

func (j *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference)SetVerifiedOverride(val *string) {
	if err := j.validateSetVerifiedOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifiedOverride",
		val,
	)
}

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) ResetDeletedOverride() {
	_jsii_.InvokeVoid(
		a,
		"resetDeletedOverride",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) ResetSkippedOverride() {
	_jsii_.InvokeVoid(
		a,
		"resetSkippedOverride",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) ResetTransferredOverride() {
	_jsii_.InvokeVoid(
		a,
		"resetTransferredOverride",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) ResetVerifiedOverride() {
	_jsii_.InvokeVoid(
		a,
		"resetVerifiedOverride",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTask_ReportOverridesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

