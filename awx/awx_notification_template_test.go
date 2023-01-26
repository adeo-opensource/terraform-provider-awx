package awx

import (
	awx "github.com/denouche/goawx/client"
)

var resourceDataMapNotificationTemplate = map[string]interface{}{
	"name":                       "foo",
	"notification_template_id":   1,
	"notification_configuration": "{ \"data\":\"data\"}",
}

func (mockAwx MockAWX) ListNotificationTemplates(params map[string]string) ([]*awx.NotificationTemplate, *awx.ListNotificationTemplatesResponse, error) {
	args := mockAwx.Called(params)
	return args.Get(0).([]*awx.NotificationTemplate), args.Get(1).(*awx.ListNotificationTemplatesResponse), args.Error(2)
}

func (mockAwx MockAWX) AssociateJobTemplateNotificationTemplatesError(jobTemplateID int, notificationTemplateID int) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(jobTemplateID, notificationTemplateID)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}
func (mockAwx MockAWX) AssociateJobTemplateNotificationTemplatesSuccess(jobTemplateID int, notificationTemplateID int) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(jobTemplateID, notificationTemplateID)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}
func (mockAwx MockAWX) AssociateJobTemplateNotificationTemplatesStarted(jobTemplateID int, notificationTemplateID int) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(jobTemplateID, notificationTemplateID)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}
func (mockAwx MockAWX) DisassociateJobTemplateNotificationTemplatesError(jobTemplateID int, notificationTemplateID int) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(jobTemplateID, notificationTemplateID)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}
func (mockAwx MockAWX) DisassociateJobTemplateNotificationTemplatesSuccess(jobTemplateID int, notificationTemplateID int) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(jobTemplateID, notificationTemplateID)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}
func (mockAwx MockAWX) DisassociateJobTemplateNotificationTemplatesStarted(jobTemplateID int, notificationTemplateID int) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(jobTemplateID, notificationTemplateID)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}

func (mockAwx MockAWX) CreateNotificationTemplate(data map[string]interface{}, params map[string]string) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(data, params)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}
func (mockAwx MockAWX) GetNotificationTemplateByID(id int, params map[string]string) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}
func (mockAwx MockAWX) UpdateNotificationTemplate(id int, data map[string]interface{}, params map[string]string) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}

func (mockAwx MockAWX) DeleteNotificationTemplate(id int) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(id)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}

func (mockAwx MockAWX) AssociateWorkflowJobTemplateNotificationTemplatesError(jobTemplateID int, notificationTemplateID int) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(jobTemplateID, notificationTemplateID)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}
func (mockAwx MockAWX) AssociateWorkflowJobTemplateNotificationTemplatesSuccess(jobTemplateID int, notificationTemplateID int) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(jobTemplateID, notificationTemplateID)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}
func (mockAwx MockAWX) AssociateWorkflowJobTemplateNotificationTemplatesStarted(jobTemplateID int, notificationTemplateID int) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(jobTemplateID, notificationTemplateID)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}
func (mockAwx MockAWX) AssociateWorkflowJobTemplateNotificationTemplatesApprovals(jobTemplateID int, notificationTemplateID int) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(jobTemplateID, notificationTemplateID)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}
func (mockAwx MockAWX) DisassociateWorkflowJobTemplateNotificationTemplatesError(jobTemplateID int, notificationTemplateID int) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(jobTemplateID, notificationTemplateID)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}
func (mockAwx MockAWX) DisassociateWorkflowJobTemplateNotificationTemplatesSuccess(jobTemplateID int, notificationTemplateID int) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(jobTemplateID, notificationTemplateID)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}
func (mockAwx MockAWX) DisassociateWorkflowJobTemplateNotificationTemplatesStarted(jobTemplateID int, notificationTemplateID int) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(jobTemplateID, notificationTemplateID)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}
func (mockAwx MockAWX) DisassociateWorkflowJobTemplateNotificationTemplatesApprovals(jobTemplateID int, notificationTemplateID int) (*awx.NotificationTemplate, error) {
	args := mockAwx.Called(jobTemplateID, notificationTemplateID)
	return args.Get(0).(*awx.NotificationTemplate), args.Error(1)
}
