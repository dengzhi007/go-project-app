package common

import (
    "fmt"
)

type AlarmedStatus struct {
    last_alarmed_status     int // 0 recover 1 abnormal
    last_value_status       int // 0 abnormal 1 normal
    already_alarmed_count   int
    continue_abnormal_count int

    //status constants
    alarmed_recover         int // 0
    alarmed_abnormal        int // 1
    value_normal            int // 1 
    value_abnormal          int // 0
}

func NewAlarmedStatus() *AlarmedStatus {
    
    return &AlarmedStatus{
	last_alarmed_status: 0,
	last_value_status: 0,
	already_alarmed_count: 0,
	continue_abnormal_count: 0,
	alarmed_recover: 0,
	alarmed_abnormal: 1,
	value_normal: 1,
	value_abnormal: 0,    
    }
}

func (as *AlarmedStatus)UpdateFromRedisStr(redis_str string) {
    if len(redis_str) == 0 {
	return
    }

    fmt.Sscanf(redis_str, "%d,%d,%d,%d", &as.last_alarmed_status, &as.last_value_status, &as.already_alarmed_count, &as.continue_abnormal_count)
}

func (as *AlarmedStatus)ConvertToRedisStr() string {
    return fmt.Sprintf("%d,%d,%d,%d", as.last_alarmed_status, as.last_value_status, as.already_alarmed_count, as.continue_abnormal_count)
}

//----------------------------------------------------------------------------------------------------------------------------------------------------

type AlarmPackage struct {
    hostname string
    rule_id  int
    item_id  int
    value    int 
    status   int 
}



func (ap *AlarmPackage)UpdateFromBytes(bytes []byte) {
    if  len(bytes) == 0 || len(bytes) > 102400 {
	return
    }

    fmt.Sscanf(string(bytes), "%s\x01%d\x01%d\x01%d\x01%d", &ap.hostname, &ap.rule_id, &ap.item_id, &ap.value, &ap.status)
}

func (ap *AlarmPackage)ConvertToBytes() ([]byte, error) {
    if len(ap.hostname) == 0 {
	return nil, fmt.Errorf("hostname empty")
    }

    body_str := fmt.Sprintf("%s\x01%d\x01%d\x01%d\x01%d", ap.hostname, ap.rule_id, ap.item_id, ap.value, ap.status)
    body := []byte(body_str)
    body_len := len(body)
    if body_len == 0 || body_len > 102400 {
	return nil, fmt.Errorf("body length must between 0 ~ 100k")
    }

    head_str := fmt.Sprintf("%010d", body_len)
    head := []byte(head_str)
    if len(head) != 10 {
	return nil, fmt.Errorf("head length must equals to 10")
    }

    body_len += 10
    bytes := make([]byte, body_len) 
    i := 0
    for i < 10 {
	bytes[i] = head[i] 
	i += 1
    }
    j := 0
    for i < body_len {
	bytes[i] = body[j]
	i += 1
	j += 1
    }
    return bytes, nil
}

//------------------------------------------------------------------------------------------------------------------------------------------------------
type Task struct {
    task_type int // 1 tcp 2 http 3 ping 4 dns 5 udp
    item_id   int
    next_time int //
    hostname  string
    has_ran   bool
    value     int 
    executing bool
}
