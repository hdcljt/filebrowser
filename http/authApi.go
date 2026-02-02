package fbhttp

import (
	"encoding/json"
	"net/http"
	"time"
)

// AuthStatusResponse 授权状态响应
type AuthStatusResponse struct {
	NeedAuth    bool `json:"needAuth"`    // 是否需要授权
	HasApplied  bool `json:"hasApplied"`  // 是否已申请授权
	IsApproved  bool `json:"isApproved"`  // 授权是否已批准
	Deadline    string `json:"deadline"`   // 授权截止时间
	RemainingSeconds int `json:"remainingSeconds"` // 剩余秒数
}

// AuthApplyRequest 授权申请请求
type AuthApplyRequest struct {
	Condition   string `json:"condition"`   // 授权条件：count或time
	TimeDuration string `json:"timeDuration"` // 授权时长（小时）
	Mode        string `json:"mode"`        // 授权模式：remote
	Approver    string `json:"approver"`    // 审批人
	Reason      string `json:"reason"`      // 申请原因
}

// AuthApplyResponse 授权申请响应
type AuthApplyResponse struct {
	Success bool   `json:"success"` // 是否成功
	Message string `json:"message"` // 消息
	ApplyID string `json:"applyID"` // 申请ID
}

// AuthVerifyRequest 授权验证请求
type AuthVerifyRequest struct {
	Code string `json:"code"` // 授权码
}

// AuthVerifyResponse 授权验证响应
type AuthVerifyResponse struct {
	Success bool   `json:"success"` // 是否成功
	Message string `json:"message"` // 消息
}

// authCheckHandler 检查授权状态
var authCheckHandler = withUser(func(w http.ResponseWriter, r *http.Request, d *data) (int, error) {
	// 模拟授权状态检查
	// 在实际实现中，应该从数据库或缓存中获取用户的授权状态
	response := AuthStatusResponse{
		NeedAuth:    true, // 假设所有下载都需要授权
		HasApplied:  false, // 假设用户尚未申请授权
		IsApproved:  false, // 假设授权尚未批准
		Deadline:    time.Now().Add(30 * time.Minute).Format("2006-01-02 15:04:05"),
		RemainingSeconds: 1800,
	}

	w.Header().Set("Content-Type", "application/json")
	return 0, json.NewEncoder(w).Encode(response)
})

// authApplyHandler 处理授权申请
var authApplyHandler = withUser(func(w http.ResponseWriter, r *http.Request, d *data) (int, error) {
	if r.Body == nil {
		return http.StatusBadRequest, nil
	}

	req := &AuthApplyRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, err
	}

	// 模拟授权申请处理
	// 在实际实现中，应该将申请信息保存到数据库，并通知审批人
	response := AuthApplyResponse{
		Success: true,
		Message: "授权申请已提交，请等待审批",
		ApplyID: "apply_" + time.Now().Format("20060102150405"),
	}

	w.Header().Set("Content-Type", "application/json")
	return 0, json.NewEncoder(w).Encode(response)
})

// authVerifyHandler 处理授权验证
var authVerifyHandler = withUser(func(w http.ResponseWriter, r *http.Request, d *data) (int, error) {
	if r.Body == nil {
		return http.StatusBadRequest, nil
	}

	req := &AuthVerifyRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return http.StatusBadRequest, err
	}

	// 模拟授权验证
	// 在实际实现中，应该验证授权码的有效性
	if len(req.Code) < 6 {
		response := AuthVerifyResponse{
			Success: false,
			Message: "授权码无效",
		}
		w.Header().Set("Content-Type", "application/json")
		return 0, json.NewEncoder(w).Encode(response)
	}

	response := AuthVerifyResponse{
		Success: true,
		Message: "授权验证成功",
	}

	w.Header().Set("Content-Type", "application/json")
	return 0, json.NewEncoder(w).Encode(response)
})
