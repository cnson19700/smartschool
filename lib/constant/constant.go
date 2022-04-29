package constant

import "time"

const AcceptLateMinute time.Duration = time.Minute * time.Duration(20)
const AcceptEarlyMinute time.Duration = time.Minute * time.Duration(20)
const QRPrefix string = "22"
const QRSecretKey string = "Keep read as your greed's lead, Each greed thing is a good news, Did you see the news of yourself, You are now cursed to the hell! - KEDY"
const AcceptRefreshQRSecond time.Duration = time.Second * time.Duration(30)
const AcceptDeviceSignalDelay time.Duration = time.Second * time.Duration(5)

const StudentRole uint = 1
const AcademicSectionRole uint = 2
const TeacherRole uint = 3

const CheckinType_Card = "Card"
const CheckinType_QR = "QR"
const CheckinType_Error = "ERROR"

const CheckinStatus_InvalidCheckinTime string = "[Abnormal]: Invalid Checkin time - Delay over 5 second"
const CheckinStatus_Attend string = "Attend"
const CheckinStatus_Late string = "Late"
const CheckinStatus_InvalidCardUserNotFound string = "[Abnormal]: Invalid Card - User not found"
const CheckinStatus_InvalidCardRoleNotFound string = "[Abnormal]: Invalid Card - Role not found"
const CheckinStatus_InvalidQR string = "[Abnormal]: Invalid format QR or Expired QR"
const CheckinStatus_InvalidQRRoleNotFound string = "[Abnormal]: Invalid QR - Role not found"
const CheckinStatus_InvalidInfo string = "[Abnormal]: Invalid format CardID"
const CheckinStatus_ErrorQueryDevice string = "[Abnormal]: Error when query Device"
const CheckinStatus_DeviceNotFound string = "[Abnormal]: Device does not match any room"
const CheckinStatus_ErrorQuerySchedule string = "[Abnormal]: Error when query Schedule"
const CheckinStatus_ScheduleNotFound string = "[Normal]: Forseen time slot not in any Schedule"
const CheckinStatus_SameScheduleSpam string = "[Normal]: Spam check-in"
const CheckinStatus_AmbiguousUserRole string = "[Abnormal]: Ambiguous user role"
const CheckinStatus_ErrorQueryEnrollment string = "[Abnormal]: Error when query Enrollment of student or teacher"
const CheckinStatus_ErrorQueryAttendance string = "[Abnormal]: Error when query Attendance"
const CheckinStatus_ErrorCreateAttendance string = "[Abnormal]: Error when create Attendance"
const CheckinStatus_Success string = "[Normal]: Checkin Success"
const CheckinStatus_Exist string = "[Normal]: Checkin Exist"
const CheckinStatus_EnrollmentNotFound string = "[Normal]: Student dont take this course"
const CheckinStatus_ErrorLogic string = "[Abnormal]: Error in Logic Check-in"
