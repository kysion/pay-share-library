package order

import "github.com/kysion/base-library/utility/enum"

type AuditStateEnum enum.IEnumCode[int]

type auditState struct {
    Reject    AuditStateEnum
    WaitAudit AuditStateEnum
    Approve   AuditStateEnum
}

var AuditState = auditState{
    Reject:    enum.New[AuditStateEnum](-1, "不通过"),
    WaitAudit: enum.New[AuditStateEnum](0, "待审核"),
    Approve:   enum.New[AuditStateEnum](1, "通过"),
}

func (e auditState) New(code int, description string) AuditStateEnum {
    if (code&AuditState.Reject.Code()) == AuditState.Reject.Code() ||
        (code&AuditState.WaitAudit.Code()) == AuditState.WaitAudit.Code() ||
        (code&AuditState.Approve.Code()) == AuditState.Approve.Code() {
        return enum.New[AuditStateEnum](code, description)

    } else {
        panic("Video.AuditState.New: error")
    }
}
