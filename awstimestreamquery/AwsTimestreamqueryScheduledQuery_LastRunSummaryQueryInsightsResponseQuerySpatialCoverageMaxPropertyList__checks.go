//go:build !no_runtime_type_checking

package awstimestreamquery

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxPropertyList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxPropertyList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxProperty:
		val := val.(*[]*AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxProperty)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxProperty:
		val_ := val.([]*AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxProperty)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *[]*AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxPropertyList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsTimestreamqueryScheduledQuery_LastRunSummaryQueryInsightsResponseQuerySpatialCoverageMaxPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

