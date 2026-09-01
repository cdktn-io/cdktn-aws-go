package awsconnectcustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconnectcustomerprofiles/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsconnectcustomerprofiles/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCustomerprofilesDomain_MatchingPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AutoMerging() AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference
	// Experimental.
	AutoMergingInput() *AwsCustomerprofilesDomain_AutoMergingProperty
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
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	ExportingConfig() AwsCustomerprofilesDomain_MatchingExportingConfigPropertyOutputReference
	// Experimental.
	ExportingConfigInput() *AwsCustomerprofilesDomain_MatchingExportingConfigProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsCustomerprofilesDomain_MatchingProperty
	// Experimental.
	SetInternalValue(val *AwsCustomerprofilesDomain_MatchingProperty)
	// Experimental.
	JobSchedule() AwsCustomerprofilesDomain_JobSchedulePropertyOutputReference
	// Experimental.
	JobScheduleInput() *AwsCustomerprofilesDomain_JobScheduleProperty
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
	PutAutoMerging(value *AwsCustomerprofilesDomain_AutoMergingProperty)
	// Experimental.
	PutExportingConfig(value *AwsCustomerprofilesDomain_MatchingExportingConfigProperty)
	// Experimental.
	PutJobSchedule(value *AwsCustomerprofilesDomain_JobScheduleProperty)
	// Experimental.
	ResetAutoMerging()
	// Experimental.
	ResetExportingConfig()
	// Experimental.
	ResetJobSchedule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCustomerprofilesDomain_MatchingPropertyOutputReference
type jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) AutoMerging() AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference {
	var returns AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference
	_jsii_.Get(
		j,
		"autoMerging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) AutoMergingInput() *AwsCustomerprofilesDomain_AutoMergingProperty {
	var returns *AwsCustomerprofilesDomain_AutoMergingProperty
	_jsii_.Get(
		j,
		"autoMergingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) ExportingConfig() AwsCustomerprofilesDomain_MatchingExportingConfigPropertyOutputReference {
	var returns AwsCustomerprofilesDomain_MatchingExportingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"exportingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) ExportingConfigInput() *AwsCustomerprofilesDomain_MatchingExportingConfigProperty {
	var returns *AwsCustomerprofilesDomain_MatchingExportingConfigProperty
	_jsii_.Get(
		j,
		"exportingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) InternalValue() *AwsCustomerprofilesDomain_MatchingProperty {
	var returns *AwsCustomerprofilesDomain_MatchingProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) JobSchedule() AwsCustomerprofilesDomain_JobSchedulePropertyOutputReference {
	var returns AwsCustomerprofilesDomain_JobSchedulePropertyOutputReference
	_jsii_.Get(
		j,
		"jobSchedule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) JobScheduleInput() *AwsCustomerprofilesDomain_JobScheduleProperty {
	var returns *AwsCustomerprofilesDomain_JobScheduleProperty
	_jsii_.Get(
		j,
		"jobScheduleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCustomerprofilesDomain_MatchingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCustomerprofilesDomain_MatchingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCustomerprofilesDomain_MatchingPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.AwsCustomerprofilesDomain.MatchingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCustomerprofilesDomain_MatchingPropertyOutputReference_Override(a AwsCustomerprofilesDomain_MatchingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.AwsCustomerprofilesDomain.MatchingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference)SetInternalValue(val *AwsCustomerprofilesDomain_MatchingProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) PutAutoMerging(value *AwsCustomerprofilesDomain_AutoMergingProperty) {
	if err := a.validatePutAutoMergingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutoMerging",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) PutExportingConfig(value *AwsCustomerprofilesDomain_MatchingExportingConfigProperty) {
	if err := a.validatePutExportingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExportingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) PutJobSchedule(value *AwsCustomerprofilesDomain_JobScheduleProperty) {
	if err := a.validatePutJobScheduleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJobSchedule",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) ResetAutoMerging() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoMerging",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) ResetExportingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetExportingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) ResetJobSchedule() {
	_jsii_.InvokeVoid(
		a,
		"resetJobSchedule",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_MatchingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

