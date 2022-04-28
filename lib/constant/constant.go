package constant

import "time"

const AcceptLateMinute time.Duration = time.Minute * time.Duration(20)
const AcceptEarlyMinute time.Duration = time.Minute * time.Duration(20)
const QRPrefix string = "22"
const QRSecretKey string = "Keep read as your greed's lead, Each greed thing is a good news, Did you see the news of yourself, You are now cursed to the hell! - KEDY"
const AcceptRefreshQRSecond time.Duration = time.Second * time.Duration(30)
const AcceptDeviceSignalDelay time.Duration = time.Second * time.Duration(5)
const StudentRole string = "Student"
const TeacherRole string = "Teacher"
