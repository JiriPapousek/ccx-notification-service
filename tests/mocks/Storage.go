// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	"time"

	types "github.com/RedHatInsights/ccx-notification-service/types"
	mock "github.com/stretchr/testify/mock"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// CleanupNewReports provides a mock function with given fields: maxAge
func (_m *Storage) CleanupNewReports(maxAge string) (int, error) {
	ret := _m.Called(maxAge)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(maxAge)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(maxAge)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CleanupOldReports provides a mock function with given fields: maxAge
func (_m *Storage) CleanupOldReports(maxAge string) (int, error) {
	ret := _m.Called(maxAge)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(maxAge)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(maxAge)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *Storage) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteRowFromNewReports provides a mock function with given fields: orgID, clusterName, updatedAt
func (_m *Storage) DeleteRowFromNewReports(orgID types.OrgID, clusterName types.ClusterName, updatedAt types.Timestamp) (int, error) {
	ret := _m.Called(orgID, clusterName, updatedAt)

	var r0 int
	if rf, ok := ret.Get(0).(func(types.OrgID, types.ClusterName, types.Timestamp) int); ok {
		r0 = rf(orgID, clusterName, updatedAt)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.OrgID, types.ClusterName, types.Timestamp) error); ok {
		r1 = rf(orgID, clusterName, updatedAt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteRowFromReported provides a mock function with given fields: orgID, clusterName, notifiedAt
func (_m *Storage) DeleteRowFromReported(orgID types.OrgID, clusterName types.ClusterName, notifiedAt types.Timestamp) (int, error) {
	ret := _m.Called(orgID, clusterName, notifiedAt)

	var r0 int
	if rf, ok := ret.Get(0).(func(types.OrgID, types.ClusterName, types.Timestamp) int); ok {
		r0 = rf(orgID, clusterName, notifiedAt)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.OrgID, types.ClusterName, types.Timestamp) error); ok {
		r1 = rf(orgID, clusterName, notifiedAt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrintNewReportsForCleanup provides a mock function with given fields: maxAge
func (_m *Storage) PrintNewReportsForCleanup(maxAge string) error {
	ret := _m.Called(maxAge)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(maxAge)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PrintOldReportsForCleanup provides a mock function with given fields: maxAge
func (_m *Storage) PrintOldReportsForCleanup(maxAge string) error {
	ret := _m.Called(maxAge)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(maxAge)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadClusterList provides a mock function with given fields:
func (_m *Storage) ReadClusterList() ([]types.ClusterEntry, error) {
	ret := _m.Called()

	var r0 []types.ClusterEntry
	if rf, ok := ret.Get(0).(func() []types.ClusterEntry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.ClusterEntry)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadLastNotifiedRecordForClusterList provides a mock function with given fields: clusterEntries, timeOffset, eventTarget
func (_m *Storage) ReadLastNotifiedRecordForClusterList(clusterEntries []types.ClusterEntry, timeOffset string, eventTarget types.EventTarget) (types.NotifiedRecordsPerCluster, error) {
	ret := _m.Called(clusterEntries, timeOffset, eventTarget)

	var r0 types.NotifiedRecordsPerCluster
	if rf, ok := ret.Get(0).(func([]types.ClusterEntry, string, types.EventTarget) types.NotifiedRecordsPerCluster); ok {
		r0 = rf(clusterEntries, timeOffset, eventTarget)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.NotifiedRecordsPerCluster)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]types.ClusterEntry, string, types.EventTarget) error); ok {
		r1 = rf(clusterEntries, timeOffset, eventTarget)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadNotificationTypes provides a mock function with given fields:
func (_m *Storage) ReadNotificationTypes() ([]types.NotificationType, error) {
	ret := _m.Called()

	var r0 []types.NotificationType
	if rf, ok := ret.Get(0).(func() []types.NotificationType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.NotificationType)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadReportForCluster provides a mock function with given fields: orgID, clusterName
func (_m *Storage) ReadReportForCluster(orgID types.OrgID, clusterName types.ClusterName) (types.ClusterReport, types.Timestamp, error) {
	ret := _m.Called(orgID, clusterName)

	var r0 types.ClusterReport
	if rf, ok := ret.Get(0).(func(types.OrgID, types.ClusterName) types.ClusterReport); ok {
		r0 = rf(orgID, clusterName)
	} else {
		r0 = ret.Get(0).(types.ClusterReport)
	}

	var r1 types.Timestamp
	if rf, ok := ret.Get(1).(func(types.OrgID, types.ClusterName) types.Timestamp); ok {
		r1 = rf(orgID, clusterName)
	} else {
		r1 = ret.Get(1).(types.Timestamp)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(types.OrgID, types.ClusterName) error); ok {
		r2 = rf(orgID, clusterName)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ReadReportForClusterAtOffset provides a mock function with given fields: orgID, clusterName, offset
func (_m *Storage) ReadReportForClusterAtOffset(orgID types.OrgID, clusterName types.ClusterName, offset types.KafkaOffset) (types.ClusterReport, error) {
	ret := _m.Called(orgID, clusterName, offset)

	var r0 types.ClusterReport
	if rf, ok := ret.Get(0).(func(types.OrgID, types.ClusterName, types.KafkaOffset) types.ClusterReport); ok {
		r0 = rf(orgID, clusterName, offset)
	} else {
		r0 = ret.Get(0).(types.ClusterReport)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.OrgID, types.ClusterName, types.KafkaOffset) error); ok {
		r1 = rf(orgID, clusterName, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadReportForClusterAtTime provides a mock function with given fields: orgID, clusterName, updatedAt
func (_m *Storage) ReadReportForClusterAtTime(orgID types.OrgID, clusterName types.ClusterName, updatedAt types.Timestamp) (types.ClusterReport, error) {
	ret := _m.Called(orgID, clusterName, updatedAt)

	var r0 types.ClusterReport
	if rf, ok := ret.Get(0).(func(types.OrgID, types.ClusterName, types.Timestamp) types.ClusterReport); ok {
		r0 = rf(orgID, clusterName, updatedAt)
	} else {
		r0 = ret.Get(0).(types.ClusterReport)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.OrgID, types.ClusterName, types.Timestamp) error); ok {
		r1 = rf(orgID, clusterName, updatedAt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadStates provides a mock function with given fields:
func (_m *Storage) ReadStates() ([]types.State, error) {
	ret := _m.Called()

	var r0 []types.State
	if rf, ok := ret.Get(0).(func() []types.State); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.State)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WriteNotificationRecord provides a mock function with given fields: notificationRecord
func (_m *Storage) WriteNotificationRecord(notificationRecord types.NotificationRecord) error {
	ret := _m.Called(notificationRecord)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.NotificationRecord) error); ok {
		r0 = rf(notificationRecord)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteNotificationRecordForCluster provides a mock function with given fields: clusterEntry, notificationTypeID, stateID, report, notifiedAt, errorLog, eventTarget
func (_m *Storage) WriteNotificationRecordForCluster(clusterEntry types.ClusterEntry, notificationTypeID types.NotificationTypeID, stateID types.StateID, report types.ClusterReport, notifiedAt types.Timestamp, errorLog string, eventTarget types.EventTarget) error {
	ret := _m.Called(clusterEntry, notificationTypeID, stateID, report, notifiedAt, errorLog, eventTarget)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.ClusterEntry, types.NotificationTypeID, types.StateID, types.ClusterReport, types.Timestamp, string, types.EventTarget) error); ok {
		r0 = rf(clusterEntry, notificationTypeID, stateID, report, notifiedAt, errorLog, eventTarget)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteNotificationRecordImpl provides a mock function with given fields: orgID, accountNumber, clusterName, notificationTypeID, stateID, report, updatedAt, notifiedAt, errorLog, eventType
func (_m *Storage) WriteNotificationRecordImpl(orgID types.OrgID, accountNumber types.AccountNumber, clusterName types.ClusterName, notificationTypeID types.NotificationTypeID, stateID types.StateID, report types.ClusterReport, updatedAt types.Timestamp, notifiedAt types.Timestamp, errorLog string, eventType types.EventTarget) error {
	ret := _m.Called(orgID, accountNumber, clusterName, notificationTypeID, stateID, report, updatedAt, notifiedAt, errorLog, eventType)

	var r0 error
	if rf, ok := ret.Get(0).(func(types.OrgID, types.AccountNumber, types.ClusterName, types.NotificationTypeID, types.StateID, types.ClusterReport, types.Timestamp, types.Timestamp, string, types.EventTarget) error); ok {
		r0 = rf(orgID, accountNumber, clusterName, notificationTypeID, stateID, report, updatedAt, notifiedAt, errorLog, eventType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *Storage) WriteReadError(
	orgID types.OrgID,
	clusterName types.ClusterName,
	lastCheckedTime time.Time,
	e error,
) error {
	return nil
}

func (_m *Storage) ReadErrorExists(
	orgID types.OrgID,
	clusterName types.ClusterName,
	lastCheckedTime time.Time,
) (bool, error) {
	return false, nil
}

type mockConstructorTestingTNewStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewStorage creates a new instance of Storage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStorage(t mockConstructorTestingTNewStorage) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
