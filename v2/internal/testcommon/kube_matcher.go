/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"context"

	"github.com/onsi/gomega/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// TODO: Would we rather these just be on testcontext? Might read better
type KubeMatcher struct {
	ctx    context.Context
	verify *Verify
}

func NewKubeMatcher(verify *Verify, ctx context.Context) *KubeMatcher {
	return &KubeMatcher{
		verify: verify,
		ctx:    ctx,
	}
}

func (m *KubeMatcher) BeProvisioned() types.GomegaMatcher {
	return &DesiredStateMatcher{
		verify:            m.verify,
		ctx:               m.ctx,
		readyGoalStatus:   metav1.ConditionTrue,
		readyGoalSeverity: conditions.ConditionSeverityNone,
	}
}

func (m *KubeMatcher) BeProvisionedAfter(previousReadyCondition conditions.Condition) types.GomegaMatcher {
	return &DesiredStateMatcher{
		verify:              m.verify,
		ctx:                 m.ctx,
		readyGoalStatus:     metav1.ConditionTrue,
		readyGoalSeverity:   conditions.ConditionSeverityNone,
		readyPreviousStatus: &previousReadyCondition.Status,
	}
}

func (m *KubeMatcher) BeFailed() types.GomegaMatcher {
	return &DesiredStateMatcher{
		verify:            m.verify,
		ctx:               m.ctx,
		readyGoalStatus:   metav1.ConditionFalse,
		readyGoalSeverity: conditions.ConditionSeverityError,
	}
}

func (m *KubeMatcher) BeDeleted() types.GomegaMatcher {
	return &BeDeletedMatcher{
		verify: m.verify,
		ctx:    m.ctx,
	}
}
