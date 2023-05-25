package awx

/*
Base for job template notification template.
*/

import (
	"context"
	"strconv"

	awx "github.com/adeo-opensource/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func getResourceJobTemplateNotificationTemplateAssociateFuncForType(client awx.JobTemplateNotificationTemplateService, typ string) func(jobTemplateID int, notificationTemplateID int) (*awx.NotificationTemplate, error) {
	switch typ {
	case "error":
		return client.AssociateJobTemplateNotificationTemplatesError
	case "success":
		return client.AssociateJobTemplateNotificationTemplatesSuccess
	case "started":
		return client.AssociateJobTemplateNotificationTemplatesStarted
	}
	return nil
}

func getResourceJobTemplateNotificationTemplateDisassociateFuncForType(client awx.JobTemplateNotificationTemplateService, typ string) func(jobTemplateID int, notificationTemplateID int) (*awx.NotificationTemplate, error) {
	switch typ {
	case "error":
		return client.DisassociateJobTemplateNotificationTemplatesError
	case "success":
		return client.DisassociateJobTemplateNotificationTemplatesSuccess
	case "started":
		return client.DisassociateJobTemplateNotificationTemplatesStarted
	}
	return nil
}

func resourceJobTemplateNotificationTemplateCreateForType(typ string) func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		var diags diag.Diagnostics
		awxService := m.(awx.AWX)
		jobTemplateID := d.Get("job_template_id").(int)
		_, err := awxService.JobTemplateService.GetByID(jobTemplateID, make(map[string]string))
		if err != nil {
			return buildDiagNotFoundFail("job template", jobTemplateID, err)
		}

		notificationTemplateID := d.Get("notification_template_id").(int)
		associationFunc := getResourceJobTemplateNotificationTemplateAssociateFuncForType(awxService.JobTemplateNotificationTemplatesService, typ)
		if associationFunc == nil {
			return buildDiagnosticsMessage("Create: JobTemplate not AssociateJobTemplateNotificationTemplates", "Fail to find association function for notification_template type %s", typ)
		}

		result, err := associationFunc(jobTemplateID, notificationTemplateID)
		if err != nil {
			return buildDiagnosticsMessage("Create: JobTemplate not AssociateJobTemplateNotificationTemplates", "Fail to associate notification_template credentials with ID %v, for job_template ID %v, got error: %s", notificationTemplateID, jobTemplateID, err.Error())
		}

		d.SetId(strconv.Itoa(result.ID))
		return diags
	}
}

func resourceJobTemplateNotificationTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceJobTemplateNotificationTemplateDeleteForType(typ string) func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
		var diags diag.Diagnostics
		awxService := m.(awx.AWX)
		jobTemplateID := d.Get("job_template_id").(int)
		_, err := awxService.JobTemplateService.GetByID(jobTemplateID, make(map[string]string))
		if err != nil {
			return buildDiagNotFoundFail("job template", jobTemplateID, err)
		}

		notificationTemplateID := d.Get("notification_template_id").(int)
		disassociationFunc := getResourceJobTemplateNotificationTemplateDisassociateFuncForType(awxService.JobTemplateNotificationTemplatesService, typ)
		if disassociationFunc == nil {
			return buildDiagnosticsMessage("Create: JobTemplate not DisassociateJobTemplateNotificationTemplates", "Fail to find disassociation function for notification_template type %s", typ)
		}

		_, err = disassociationFunc(jobTemplateID, notificationTemplateID)
		if err != nil {
			return buildDiagnosticsMessage("Create: JobTemplate not DisassociateJobTemplateNotificationTemplates", "Fail to disassociate notification_template credentials with ID %v, for job_template ID %v, got error: %s", notificationTemplateID, jobTemplateID, err.Error())
		}

		d.SetId("")
		return diags
	}
}
