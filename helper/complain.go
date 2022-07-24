package helper

import "github.com/smartschool/apptypes"

func MapCheckinStatus_V2E(str string) string {
	switch str {
	case apptypes.VAttend:
		return apptypes.Attend
	case apptypes.VLate:
		return apptypes.Late
	case apptypes.VAbsence:
		return apptypes.Absence
	case apptypes.VAbsenceWithPermission:
		return apptypes.AbsenceWithPermission
	case apptypes.Unknown:
		return ""
	default:
		return apptypes.CheckinType_Error
	}
}

func MapCheckinStatus_E2V(str string) string {
	switch str {
	case apptypes.Attend:
		return apptypes.VAttend
	case apptypes.Late:
		return apptypes.VLate
	case apptypes.Absence:
		return apptypes.VAbsence
	case apptypes.AbsenceWithPermission:
		return apptypes.VAbsenceWithPermission
	case apptypes.Unknown:
		return ""
	default:
		return apptypes.CheckinType_Error
	}
}

func MapFormStatus_V2E(str string) string {
	switch str {
	case apptypes.VApprove:
		return apptypes.Approve
	case apptypes.VPending:
		return apptypes.Pending
	case apptypes.VReject:
		return apptypes.Reject
	default:
		return apptypes.CheckinType_Error
	}
}

func MapFormStatus_E2V(str string) string {
	switch str {
	case apptypes.Approve:
		return apptypes.VApprove
	case apptypes.Pending:
		return apptypes.VPending
	case apptypes.Reject:
		return apptypes.VReject
	default:
		return apptypes.CheckinType_Error
	}
}
