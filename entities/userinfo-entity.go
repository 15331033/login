package entities

import (
    "time"
    "fmt"
)

// UserInfo .
type UserInfo struct {
    UID        int   `xorm:"pk autoincr 'id'"` //语义标签
    UserName   string
    DepartName string
    CreateAt   *time.Time   `xorm:"created"`
}

type Admin struct {
    Admin_id        int   `xorm:"pk autoincr 'id'"` //语义标签
    Admin_account   string    `xorm:"unique"`
    Admin_password  string
    Admin_type      string
}

type Soldiers struct {
    Soldier_id      int   `xorm:"pk autoincr 'id'"` //语义标签
    Rank            string
    Id_num          string
    Name            string
    Phone_num       string
    Wechat_openid   string
    Commander_id    int
    Serve_office_id int
}

type Task struct {
    Task_id         int   `xorm:"pk autoincr 'id'"`
    Title           string
    Mem_count       int
    Launch_admin_id int 
    Launch_datetime *time.Time   
    Gather_datetime *time.Time
    Detail          string
    Gather_place_id int
    Finish_datetime *time.Time    
}

type BroadcastMessages struct {
    Bm_id           int   `xorm:"pk autoincr 'id'"`
    Title           string
    Detail          string
    Bm_type         string
    Wechat_notice   bool
    Sms_notice      bool
    Voice_notice    bool
}

type BcMsgOffices struct {
    Bmo_id          int   `xorm:"pk autoincr 'id'"`
    Msg_id          int
    Msg_office_id   int
}

type BcMsgOrgs struct {
    Bmo_id          int   `xorm:"pk autoincr 'id'"`
    Msg_id          int
    Msg_org_id      int
}

type Organizations struct {
    Org_id          int   `xorm:"pk autoincr 'id'"`
    Serve_office_id int
    Leader_sid      int
    Name            string
}

type Offices struct {
    Office_id           int   `xorm:"pk autoincr 'id'"`
    Office_level        int
    Higher_office_id    int
    Name                string
}

type CommonNotifications struct {
    Cn_id           int   `xorm:"pk autoincr 'id'"`
    Cn_bm_id        int
    Recv_soldier_id int
}

type CmNtReceipts struct {
    Cnr_id          int   `xorm:"pk autoincr 'id'"`
    Cn_id           int
    Rec_content     string
}

func init() {
    err := engine.Sync2(new(UserInfo))
    checkErr(err)
    err = engine.Sync2(new(Task))
    checkErr(err)
    err = engine.Sync2(new(Admin))
    checkErr(err)
    err = engine.Sync2(new(Soldiers))
    checkErr(err)
    err = engine.Sync2(new(BroadcastMessages))
    checkErr(err)
    err = engine.Sync2(new(BcMsgOrgs))
    checkErr(err)
    err = engine.Sync2(new(BcMsgOffices))
    checkErr(err)
    err = engine.Sync2(new(Organizations))
    checkErr(err)
    err = engine.Sync2(new(Offices))
    checkErr(err)
    err = engine.Sync2(new(CommonNotifications))
    checkErr(err)
    err = engine.Sync2(new(CmNtReceipts))
    checkErr(err)
    fmt.Println("111")
}

