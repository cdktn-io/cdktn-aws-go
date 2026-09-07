//go:build !no_runtime_type_checking

package securityhub

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutAwsAccountIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_AwsAccountIdProperty:
		value := value.(*[]*AwsInsight_AwsAccountIdProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_AwsAccountIdProperty:
		value_ := value.([]*AwsInsight_AwsAccountIdProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_AwsAccountIdProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutAwsAccountNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_AwsAccountNameProperty:
		value := value.(*[]*AwsInsight_AwsAccountNameProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_AwsAccountNameProperty:
		value_ := value.([]*AwsInsight_AwsAccountNameProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_AwsAccountNameProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutCompanyNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_CompanyNameProperty:
		value := value.(*[]*AwsInsight_CompanyNameProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_CompanyNameProperty:
		value_ := value.([]*AwsInsight_CompanyNameProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_CompanyNameProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutComplianceAssociatedStandardsIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ComplianceAssociatedStandardsIdProperty:
		value := value.(*[]*AwsInsight_ComplianceAssociatedStandardsIdProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ComplianceAssociatedStandardsIdProperty:
		value_ := value.([]*AwsInsight_ComplianceAssociatedStandardsIdProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ComplianceAssociatedStandardsIdProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutComplianceSecurityControlIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ComplianceSecurityControlIdProperty:
		value := value.(*[]*AwsInsight_ComplianceSecurityControlIdProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ComplianceSecurityControlIdProperty:
		value_ := value.([]*AwsInsight_ComplianceSecurityControlIdProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ComplianceSecurityControlIdProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutComplianceSecurityControlParametersNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ComplianceSecurityControlParametersNameProperty:
		value := value.(*[]*AwsInsight_ComplianceSecurityControlParametersNameProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ComplianceSecurityControlParametersNameProperty:
		value_ := value.([]*AwsInsight_ComplianceSecurityControlParametersNameProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ComplianceSecurityControlParametersNameProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutComplianceSecurityControlParametersValueParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ComplianceSecurityControlParametersValueProperty:
		value := value.(*[]*AwsInsight_ComplianceSecurityControlParametersValueProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ComplianceSecurityControlParametersValueProperty:
		value_ := value.([]*AwsInsight_ComplianceSecurityControlParametersValueProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ComplianceSecurityControlParametersValueProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutComplianceStatusParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ComplianceStatusProperty:
		value := value.(*[]*AwsInsight_ComplianceStatusProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ComplianceStatusProperty:
		value_ := value.([]*AwsInsight_ComplianceStatusProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ComplianceStatusProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutConfidenceParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ConfidenceProperty:
		value := value.(*[]*AwsInsight_ConfidenceProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ConfidenceProperty:
		value_ := value.([]*AwsInsight_ConfidenceProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ConfidenceProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutCreatedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_CreatedAtProperty:
		value := value.(*[]*AwsInsight_CreatedAtProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_CreatedAtProperty:
		value_ := value.([]*AwsInsight_CreatedAtProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_CreatedAtProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutCriticalityParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_CriticalityProperty:
		value := value.(*[]*AwsInsight_CriticalityProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_CriticalityProperty:
		value_ := value.([]*AwsInsight_CriticalityProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_CriticalityProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutDescriptionParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_DescriptionProperty:
		value := value.(*[]*AwsInsight_DescriptionProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_DescriptionProperty:
		value_ := value.([]*AwsInsight_DescriptionProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_DescriptionProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutFindingProviderFieldsConfidenceParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_FindingProviderFieldsConfidenceProperty:
		value := value.(*[]*AwsInsight_FindingProviderFieldsConfidenceProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_FindingProviderFieldsConfidenceProperty:
		value_ := value.([]*AwsInsight_FindingProviderFieldsConfidenceProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_FindingProviderFieldsConfidenceProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutFindingProviderFieldsCriticalityParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_FindingProviderFieldsCriticalityProperty:
		value := value.(*[]*AwsInsight_FindingProviderFieldsCriticalityProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_FindingProviderFieldsCriticalityProperty:
		value_ := value.([]*AwsInsight_FindingProviderFieldsCriticalityProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_FindingProviderFieldsCriticalityProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutFindingProviderFieldsRelatedFindingsIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_FindingProviderFieldsRelatedFindingsIdProperty:
		value := value.(*[]*AwsInsight_FindingProviderFieldsRelatedFindingsIdProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_FindingProviderFieldsRelatedFindingsIdProperty:
		value_ := value.([]*AwsInsight_FindingProviderFieldsRelatedFindingsIdProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_FindingProviderFieldsRelatedFindingsIdProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutFindingProviderFieldsRelatedFindingsProductArnParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_FindingProviderFieldsRelatedFindingsProductArnProperty:
		value := value.(*[]*AwsInsight_FindingProviderFieldsRelatedFindingsProductArnProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_FindingProviderFieldsRelatedFindingsProductArnProperty:
		value_ := value.([]*AwsInsight_FindingProviderFieldsRelatedFindingsProductArnProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_FindingProviderFieldsRelatedFindingsProductArnProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutFindingProviderFieldsSeverityLabelParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_FindingProviderFieldsSeverityLabelProperty:
		value := value.(*[]*AwsInsight_FindingProviderFieldsSeverityLabelProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_FindingProviderFieldsSeverityLabelProperty:
		value_ := value.([]*AwsInsight_FindingProviderFieldsSeverityLabelProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_FindingProviderFieldsSeverityLabelProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutFindingProviderFieldsSeverityOriginalParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_FindingProviderFieldsSeverityOriginalProperty:
		value := value.(*[]*AwsInsight_FindingProviderFieldsSeverityOriginalProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_FindingProviderFieldsSeverityOriginalProperty:
		value_ := value.([]*AwsInsight_FindingProviderFieldsSeverityOriginalProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_FindingProviderFieldsSeverityOriginalProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutFindingProviderFieldsTypesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_FindingProviderFieldsTypesProperty:
		value := value.(*[]*AwsInsight_FindingProviderFieldsTypesProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_FindingProviderFieldsTypesProperty:
		value_ := value.([]*AwsInsight_FindingProviderFieldsTypesProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_FindingProviderFieldsTypesProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutFirstObservedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_FirstObservedAtProperty:
		value := value.(*[]*AwsInsight_FirstObservedAtProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_FirstObservedAtProperty:
		value_ := value.([]*AwsInsight_FirstObservedAtProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_FirstObservedAtProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutGeneratorIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_GeneratorIdProperty:
		value := value.(*[]*AwsInsight_GeneratorIdProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_GeneratorIdProperty:
		value_ := value.([]*AwsInsight_GeneratorIdProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_GeneratorIdProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_IdProperty:
		value := value.(*[]*AwsInsight_IdProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_IdProperty:
		value_ := value.([]*AwsInsight_IdProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_IdProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutKeywordParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_KeywordProperty:
		value := value.(*[]*AwsInsight_KeywordProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_KeywordProperty:
		value_ := value.([]*AwsInsight_KeywordProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_KeywordProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutLastObservedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_LastObservedAtProperty:
		value := value.(*[]*AwsInsight_LastObservedAtProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_LastObservedAtProperty:
		value_ := value.([]*AwsInsight_LastObservedAtProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_LastObservedAtProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutMalwareNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_MalwareNameProperty:
		value := value.(*[]*AwsInsight_MalwareNameProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_MalwareNameProperty:
		value_ := value.([]*AwsInsight_MalwareNameProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_MalwareNameProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutMalwarePathParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_MalwarePathProperty:
		value := value.(*[]*AwsInsight_MalwarePathProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_MalwarePathProperty:
		value_ := value.([]*AwsInsight_MalwarePathProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_MalwarePathProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutMalwareStateParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_MalwareStateProperty:
		value := value.(*[]*AwsInsight_MalwareStateProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_MalwareStateProperty:
		value_ := value.([]*AwsInsight_MalwareStateProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_MalwareStateProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutMalwareTypeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_MalwareTypeProperty:
		value := value.(*[]*AwsInsight_MalwareTypeProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_MalwareTypeProperty:
		value_ := value.([]*AwsInsight_MalwareTypeProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_MalwareTypeProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutNetworkDestinationDomainParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_NetworkDestinationDomainProperty:
		value := value.(*[]*AwsInsight_NetworkDestinationDomainProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_NetworkDestinationDomainProperty:
		value_ := value.([]*AwsInsight_NetworkDestinationDomainProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_NetworkDestinationDomainProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutNetworkDestinationIpv4Parameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_NetworkDestinationIpv4Property:
		value := value.(*[]*AwsInsight_NetworkDestinationIpv4Property)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_NetworkDestinationIpv4Property:
		value_ := value.([]*AwsInsight_NetworkDestinationIpv4Property)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_NetworkDestinationIpv4Property; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutNetworkDestinationIpv6Parameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_NetworkDestinationIpv6Property:
		value := value.(*[]*AwsInsight_NetworkDestinationIpv6Property)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_NetworkDestinationIpv6Property:
		value_ := value.([]*AwsInsight_NetworkDestinationIpv6Property)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_NetworkDestinationIpv6Property; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutNetworkDestinationPortParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_NetworkDestinationPortProperty:
		value := value.(*[]*AwsInsight_NetworkDestinationPortProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_NetworkDestinationPortProperty:
		value_ := value.([]*AwsInsight_NetworkDestinationPortProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_NetworkDestinationPortProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutNetworkDirectionParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_NetworkDirectionProperty:
		value := value.(*[]*AwsInsight_NetworkDirectionProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_NetworkDirectionProperty:
		value_ := value.([]*AwsInsight_NetworkDirectionProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_NetworkDirectionProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutNetworkProtocolParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_NetworkProtocolProperty:
		value := value.(*[]*AwsInsight_NetworkProtocolProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_NetworkProtocolProperty:
		value_ := value.([]*AwsInsight_NetworkProtocolProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_NetworkProtocolProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutNetworkSourceDomainParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_NetworkSourceDomainProperty:
		value := value.(*[]*AwsInsight_NetworkSourceDomainProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_NetworkSourceDomainProperty:
		value_ := value.([]*AwsInsight_NetworkSourceDomainProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_NetworkSourceDomainProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutNetworkSourceIpv4Parameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_NetworkSourceIpv4Property:
		value := value.(*[]*AwsInsight_NetworkSourceIpv4Property)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_NetworkSourceIpv4Property:
		value_ := value.([]*AwsInsight_NetworkSourceIpv4Property)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_NetworkSourceIpv4Property; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutNetworkSourceIpv6Parameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_NetworkSourceIpv6Property:
		value := value.(*[]*AwsInsight_NetworkSourceIpv6Property)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_NetworkSourceIpv6Property:
		value_ := value.([]*AwsInsight_NetworkSourceIpv6Property)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_NetworkSourceIpv6Property; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutNetworkSourceMacParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_NetworkSourceMacProperty:
		value := value.(*[]*AwsInsight_NetworkSourceMacProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_NetworkSourceMacProperty:
		value_ := value.([]*AwsInsight_NetworkSourceMacProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_NetworkSourceMacProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutNetworkSourcePortParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_NetworkSourcePortProperty:
		value := value.(*[]*AwsInsight_NetworkSourcePortProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_NetworkSourcePortProperty:
		value_ := value.([]*AwsInsight_NetworkSourcePortProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_NetworkSourcePortProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutNoteTextParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_NoteTextProperty:
		value := value.(*[]*AwsInsight_NoteTextProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_NoteTextProperty:
		value_ := value.([]*AwsInsight_NoteTextProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_NoteTextProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutNoteUpdatedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_NoteUpdatedAtProperty:
		value := value.(*[]*AwsInsight_NoteUpdatedAtProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_NoteUpdatedAtProperty:
		value_ := value.([]*AwsInsight_NoteUpdatedAtProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_NoteUpdatedAtProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutNoteUpdatedByParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_NoteUpdatedByProperty:
		value := value.(*[]*AwsInsight_NoteUpdatedByProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_NoteUpdatedByProperty:
		value_ := value.([]*AwsInsight_NoteUpdatedByProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_NoteUpdatedByProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutProcessLaunchedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ProcessLaunchedAtProperty:
		value := value.(*[]*AwsInsight_ProcessLaunchedAtProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ProcessLaunchedAtProperty:
		value_ := value.([]*AwsInsight_ProcessLaunchedAtProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ProcessLaunchedAtProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutProcessNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ProcessNameProperty:
		value := value.(*[]*AwsInsight_ProcessNameProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ProcessNameProperty:
		value_ := value.([]*AwsInsight_ProcessNameProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ProcessNameProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutProcessParentPidParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ProcessParentPidProperty:
		value := value.(*[]*AwsInsight_ProcessParentPidProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ProcessParentPidProperty:
		value_ := value.([]*AwsInsight_ProcessParentPidProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ProcessParentPidProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutProcessPathParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ProcessPathProperty:
		value := value.(*[]*AwsInsight_ProcessPathProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ProcessPathProperty:
		value_ := value.([]*AwsInsight_ProcessPathProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ProcessPathProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutProcessPidParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ProcessPidProperty:
		value := value.(*[]*AwsInsight_ProcessPidProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ProcessPidProperty:
		value_ := value.([]*AwsInsight_ProcessPidProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ProcessPidProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutProcessTerminatedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ProcessTerminatedAtProperty:
		value := value.(*[]*AwsInsight_ProcessTerminatedAtProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ProcessTerminatedAtProperty:
		value_ := value.([]*AwsInsight_ProcessTerminatedAtProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ProcessTerminatedAtProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutProductArnParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ProductArnProperty:
		value := value.(*[]*AwsInsight_ProductArnProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ProductArnProperty:
		value_ := value.([]*AwsInsight_ProductArnProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ProductArnProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutProductFieldsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ProductFieldsProperty:
		value := value.(*[]*AwsInsight_ProductFieldsProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ProductFieldsProperty:
		value_ := value.([]*AwsInsight_ProductFieldsProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ProductFieldsProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutProductNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ProductNameProperty:
		value := value.(*[]*AwsInsight_ProductNameProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ProductNameProperty:
		value_ := value.([]*AwsInsight_ProductNameProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ProductNameProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutRecommendationTextParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_RecommendationTextProperty:
		value := value.(*[]*AwsInsight_RecommendationTextProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_RecommendationTextProperty:
		value_ := value.([]*AwsInsight_RecommendationTextProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_RecommendationTextProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutRecordStateParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_RecordStateProperty:
		value := value.(*[]*AwsInsight_RecordStateProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_RecordStateProperty:
		value_ := value.([]*AwsInsight_RecordStateProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_RecordStateProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutRelatedFindingsIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_RelatedFindingsIdProperty:
		value := value.(*[]*AwsInsight_RelatedFindingsIdProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_RelatedFindingsIdProperty:
		value_ := value.([]*AwsInsight_RelatedFindingsIdProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_RelatedFindingsIdProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutRelatedFindingsProductArnParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_RelatedFindingsProductArnProperty:
		value := value.(*[]*AwsInsight_RelatedFindingsProductArnProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_RelatedFindingsProductArnProperty:
		value_ := value.([]*AwsInsight_RelatedFindingsProductArnProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_RelatedFindingsProductArnProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceAwsEc2InstanceIamInstanceProfileArnParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceAwsEc2InstanceIamInstanceProfileArnProperty:
		value := value.(*[]*AwsInsight_ResourceAwsEc2InstanceIamInstanceProfileArnProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceAwsEc2InstanceIamInstanceProfileArnProperty:
		value_ := value.([]*AwsInsight_ResourceAwsEc2InstanceIamInstanceProfileArnProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceAwsEc2InstanceIamInstanceProfileArnProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceAwsEc2InstanceImageIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceAwsEc2InstanceImageIdProperty:
		value := value.(*[]*AwsInsight_ResourceAwsEc2InstanceImageIdProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceAwsEc2InstanceImageIdProperty:
		value_ := value.([]*AwsInsight_ResourceAwsEc2InstanceImageIdProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceAwsEc2InstanceImageIdProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceAwsEc2InstanceIpv4AddressesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceAwsEc2InstanceIpv4AddressesProperty:
		value := value.(*[]*AwsInsight_ResourceAwsEc2InstanceIpv4AddressesProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceAwsEc2InstanceIpv4AddressesProperty:
		value_ := value.([]*AwsInsight_ResourceAwsEc2InstanceIpv4AddressesProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceAwsEc2InstanceIpv4AddressesProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceAwsEc2InstanceIpv6AddressesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceAwsEc2InstanceIpv6AddressesProperty:
		value := value.(*[]*AwsInsight_ResourceAwsEc2InstanceIpv6AddressesProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceAwsEc2InstanceIpv6AddressesProperty:
		value_ := value.([]*AwsInsight_ResourceAwsEc2InstanceIpv6AddressesProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceAwsEc2InstanceIpv6AddressesProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceAwsEc2InstanceKeyNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceAwsEc2InstanceKeyNameProperty:
		value := value.(*[]*AwsInsight_ResourceAwsEc2InstanceKeyNameProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceAwsEc2InstanceKeyNameProperty:
		value_ := value.([]*AwsInsight_ResourceAwsEc2InstanceKeyNameProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceAwsEc2InstanceKeyNameProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceAwsEc2InstanceLaunchedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceAwsEc2InstanceLaunchedAtProperty:
		value := value.(*[]*AwsInsight_ResourceAwsEc2InstanceLaunchedAtProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceAwsEc2InstanceLaunchedAtProperty:
		value_ := value.([]*AwsInsight_ResourceAwsEc2InstanceLaunchedAtProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceAwsEc2InstanceLaunchedAtProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceAwsEc2InstanceSubnetIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceAwsEc2InstanceSubnetIdProperty:
		value := value.(*[]*AwsInsight_ResourceAwsEc2InstanceSubnetIdProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceAwsEc2InstanceSubnetIdProperty:
		value_ := value.([]*AwsInsight_ResourceAwsEc2InstanceSubnetIdProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceAwsEc2InstanceSubnetIdProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceAwsEc2InstanceTypeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceAwsEc2InstanceTypeProperty:
		value := value.(*[]*AwsInsight_ResourceAwsEc2InstanceTypeProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceAwsEc2InstanceTypeProperty:
		value_ := value.([]*AwsInsight_ResourceAwsEc2InstanceTypeProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceAwsEc2InstanceTypeProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceAwsEc2InstanceVpcIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceAwsEc2InstanceVpcIdProperty:
		value := value.(*[]*AwsInsight_ResourceAwsEc2InstanceVpcIdProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceAwsEc2InstanceVpcIdProperty:
		value_ := value.([]*AwsInsight_ResourceAwsEc2InstanceVpcIdProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceAwsEc2InstanceVpcIdProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceAwsIamAccessKeyCreatedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceAwsIamAccessKeyCreatedAtProperty:
		value := value.(*[]*AwsInsight_ResourceAwsIamAccessKeyCreatedAtProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceAwsIamAccessKeyCreatedAtProperty:
		value_ := value.([]*AwsInsight_ResourceAwsIamAccessKeyCreatedAtProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceAwsIamAccessKeyCreatedAtProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceAwsIamAccessKeyStatusParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceAwsIamAccessKeyStatusProperty:
		value := value.(*[]*AwsInsight_ResourceAwsIamAccessKeyStatusProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceAwsIamAccessKeyStatusProperty:
		value_ := value.([]*AwsInsight_ResourceAwsIamAccessKeyStatusProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceAwsIamAccessKeyStatusProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceAwsIamAccessKeyUserNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceAwsIamAccessKeyUserNameProperty:
		value := value.(*[]*AwsInsight_ResourceAwsIamAccessKeyUserNameProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceAwsIamAccessKeyUserNameProperty:
		value_ := value.([]*AwsInsight_ResourceAwsIamAccessKeyUserNameProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceAwsIamAccessKeyUserNameProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceAwsS3BucketOwnerIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceAwsS3BucketOwnerIdProperty:
		value := value.(*[]*AwsInsight_ResourceAwsS3BucketOwnerIdProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceAwsS3BucketOwnerIdProperty:
		value_ := value.([]*AwsInsight_ResourceAwsS3BucketOwnerIdProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceAwsS3BucketOwnerIdProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceAwsS3BucketOwnerNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceAwsS3BucketOwnerNameProperty:
		value := value.(*[]*AwsInsight_ResourceAwsS3BucketOwnerNameProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceAwsS3BucketOwnerNameProperty:
		value_ := value.([]*AwsInsight_ResourceAwsS3BucketOwnerNameProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceAwsS3BucketOwnerNameProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceContainerImageIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceContainerImageIdProperty:
		value := value.(*[]*AwsInsight_ResourceContainerImageIdProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceContainerImageIdProperty:
		value_ := value.([]*AwsInsight_ResourceContainerImageIdProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceContainerImageIdProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceContainerImageNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceContainerImageNameProperty:
		value := value.(*[]*AwsInsight_ResourceContainerImageNameProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceContainerImageNameProperty:
		value_ := value.([]*AwsInsight_ResourceContainerImageNameProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceContainerImageNameProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceContainerLaunchedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceContainerLaunchedAtProperty:
		value := value.(*[]*AwsInsight_ResourceContainerLaunchedAtProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceContainerLaunchedAtProperty:
		value_ := value.([]*AwsInsight_ResourceContainerLaunchedAtProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceContainerLaunchedAtProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceContainerNameParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceContainerNameProperty:
		value := value.(*[]*AwsInsight_ResourceContainerNameProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceContainerNameProperty:
		value_ := value.([]*AwsInsight_ResourceContainerNameProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceContainerNameProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceDetailsOtherParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceDetailsOtherProperty:
		value := value.(*[]*AwsInsight_ResourceDetailsOtherProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceDetailsOtherProperty:
		value_ := value.([]*AwsInsight_ResourceDetailsOtherProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceDetailsOtherProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceIdParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceIdProperty:
		value := value.(*[]*AwsInsight_ResourceIdProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceIdProperty:
		value_ := value.([]*AwsInsight_ResourceIdProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceIdProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourcePartitionParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourcePartitionProperty:
		value := value.(*[]*AwsInsight_ResourcePartitionProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourcePartitionProperty:
		value_ := value.([]*AwsInsight_ResourcePartitionProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourcePartitionProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceRegionParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceRegionProperty:
		value := value.(*[]*AwsInsight_ResourceRegionProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceRegionProperty:
		value_ := value.([]*AwsInsight_ResourceRegionProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceRegionProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceTagsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceTagsProperty:
		value := value.(*[]*AwsInsight_ResourceTagsProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceTagsProperty:
		value_ := value.([]*AwsInsight_ResourceTagsProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceTagsProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutResourceTypeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ResourceTypeProperty:
		value := value.(*[]*AwsInsight_ResourceTypeProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ResourceTypeProperty:
		value_ := value.([]*AwsInsight_ResourceTypeProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ResourceTypeProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutSeverityLabelParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_SeverityLabelProperty:
		value := value.(*[]*AwsInsight_SeverityLabelProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_SeverityLabelProperty:
		value_ := value.([]*AwsInsight_SeverityLabelProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_SeverityLabelProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutSourceUrlParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_SourceUrlProperty:
		value := value.(*[]*AwsInsight_SourceUrlProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_SourceUrlProperty:
		value_ := value.([]*AwsInsight_SourceUrlProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_SourceUrlProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutThreatIntelIndicatorCategoryParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ThreatIntelIndicatorCategoryProperty:
		value := value.(*[]*AwsInsight_ThreatIntelIndicatorCategoryProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ThreatIntelIndicatorCategoryProperty:
		value_ := value.([]*AwsInsight_ThreatIntelIndicatorCategoryProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ThreatIntelIndicatorCategoryProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutThreatIntelIndicatorLastObservedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ThreatIntelIndicatorLastObservedAtProperty:
		value := value.(*[]*AwsInsight_ThreatIntelIndicatorLastObservedAtProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ThreatIntelIndicatorLastObservedAtProperty:
		value_ := value.([]*AwsInsight_ThreatIntelIndicatorLastObservedAtProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ThreatIntelIndicatorLastObservedAtProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutThreatIntelIndicatorSourceParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ThreatIntelIndicatorSourceProperty:
		value := value.(*[]*AwsInsight_ThreatIntelIndicatorSourceProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ThreatIntelIndicatorSourceProperty:
		value_ := value.([]*AwsInsight_ThreatIntelIndicatorSourceProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ThreatIntelIndicatorSourceProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutThreatIntelIndicatorSourceUrlParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ThreatIntelIndicatorSourceUrlProperty:
		value := value.(*[]*AwsInsight_ThreatIntelIndicatorSourceUrlProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ThreatIntelIndicatorSourceUrlProperty:
		value_ := value.([]*AwsInsight_ThreatIntelIndicatorSourceUrlProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ThreatIntelIndicatorSourceUrlProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutThreatIntelIndicatorTypeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ThreatIntelIndicatorTypeProperty:
		value := value.(*[]*AwsInsight_ThreatIntelIndicatorTypeProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ThreatIntelIndicatorTypeProperty:
		value_ := value.([]*AwsInsight_ThreatIntelIndicatorTypeProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ThreatIntelIndicatorTypeProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutThreatIntelIndicatorValueParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_ThreatIntelIndicatorValueProperty:
		value := value.(*[]*AwsInsight_ThreatIntelIndicatorValueProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_ThreatIntelIndicatorValueProperty:
		value_ := value.([]*AwsInsight_ThreatIntelIndicatorValueProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_ThreatIntelIndicatorValueProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutTitleParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_TitleProperty:
		value := value.(*[]*AwsInsight_TitleProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_TitleProperty:
		value_ := value.([]*AwsInsight_TitleProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_TitleProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutTypeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_TypeProperty:
		value := value.(*[]*AwsInsight_TypeProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_TypeProperty:
		value_ := value.([]*AwsInsight_TypeProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_TypeProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutUpdatedAtParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_UpdatedAtProperty:
		value := value.(*[]*AwsInsight_UpdatedAtProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_UpdatedAtProperty:
		value_ := value.([]*AwsInsight_UpdatedAtProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_UpdatedAtProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutUserDefinedValuesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_UserDefinedValuesProperty:
		value := value.(*[]*AwsInsight_UserDefinedValuesProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_UserDefinedValuesProperty:
		value_ := value.([]*AwsInsight_UserDefinedValuesProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_UserDefinedValuesProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutVerificationStateParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_VerificationStateProperty:
		value := value.(*[]*AwsInsight_VerificationStateProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_VerificationStateProperty:
		value_ := value.([]*AwsInsight_VerificationStateProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_VerificationStateProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validatePutWorkflowStatusParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsInsight_WorkflowStatusProperty:
		value := value.(*[]*AwsInsight_WorkflowStatusProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsInsight_WorkflowStatusProperty:
		value_ := value.([]*AwsInsight_WorkflowStatusProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsInsight_WorkflowStatusProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validateSetInternalValueParameters(val *AwsInsight_FiltersProperty) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsInsight_FiltersPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsInsight_FiltersPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

