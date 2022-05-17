package apptypes

import "time"

const (
	AcceptLateMinute        time.Duration = time.Minute * time.Duration(20)
	AcceptEarlyMinute       time.Duration = time.Minute * time.Duration(20)
	QRPrefix                string        = "22"
	QRSecretKey             string        = "Keep read as your greed's lead, Each greed thing is a good news, Did you see the news of yourself, You are now cursed to the hell! - KEDY"
	AcceptRefreshQRSecond   time.Duration = time.Second * time.Duration(30)
	AcceptDeviceSignalDelay time.Duration = time.Second * time.Duration(5)

	StudentRole         uint = 1
	AcademicSectionRole uint = 2
	TeacherRole         uint = 3

	MAX_QUERY_LIMIT = 100

	CheckinType_Card  = "Card"
	CheckinType_QR    = "QR"
	CheckinType_Error = "ERROR"

	CheckinStatus_Attend                  string = "Attend"
	CheckinStatus_Late                    string = "Late"
	CheckinStatus_InvalidCheckinTime      string = "[Abnormal]: Invalid Checkin time - Delay over 5 second"
	CheckinStatus_InvalidCardUserNotFound string = "[Abnormal]: Invalid Card - User not found"
	CheckinStatus_InvalidCardRoleNotFound string = "[Abnormal]: Invalid Card - Role not found"
	CheckinStatus_InvalidQR               string = "[Abnormal]: Invalid format QR or Expired QR"
	CheckinStatus_InvalidQRRoleNotFound   string = "[Abnormal]: Invalid QR - Role not found"
	CheckinStatus_InvalidInfo             string = "[Abnormal]: Invalid format CardID"
	CheckinStatus_ErrorQueryDevice        string = "[Abnormal]: Error when query Device"
	CheckinStatus_DeviceNotFound          string = "[Abnormal]: Device does not match any room"
	CheckinStatus_ErrorQuerySchedule      string = "[Abnormal]: Error when query Schedule"
	CheckinStatus_ErrorLogic              string = "[Abnormal]: Error in Logic Check-in"
	CheckinStatus_AmbiguousUserRole       string = "[Abnormal]: Ambiguous user role"
	CheckinStatus_ErrorQueryEnrollment    string = "[Abnormal]: Error when query Enrollment of student or teacher"
	CheckinStatus_ErrorQueryAttendance    string = "[Abnormal]: Error when query Attendance"
	CheckinStatus_ErrorCreateAttendance   string = "[Abnormal]: Error when create Attendance"
	CheckinStatus_Success                 string = "[Normal]: Checkin Success"
	CheckinStatus_Exist                   string = "[Normal]: Checkin Exist"
	CheckinStatus_EnrollmentNotFound      string = "[Normal]: Student dont take this course"
	CheckinStatus_ScheduleNotFound        string = "[Normal]: Forseen time slot not in any Schedule"
	CheckinStatus_SameScheduleSpam        string = "[Normal]: Spam check-in"
)