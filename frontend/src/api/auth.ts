import { fetchURL } from "./utils";

// 检查授权状态
export async function checkAuthStatus() {
  const res = await fetchURL("/api/authorization/check", {
    method: "GET"
  });
  return res.json();
}

// 提交授权申请（获取验证码）
export async function submitAuthApply(data: any) {
  /* const res = await fetchURL("/api/authorization/apply", {
    method: "POST",
    body: JSON.stringify(data),
    headers: {
      "Content-Type": "application/json"
    }
  });
  return res.json(); */
  return {
    success: true,
    requestId: "123456",
    message: "验证码已发送",
  }
}

// 验证授权码（提交认证）
export async function verifyAuthCode(data: any) {
  /* const res = await fetchURL("/api/authorization/authenticate", {
    method: "POST",
    body: JSON.stringify(data),
    headers: {
      "Content-Type": "application/json"
    }
  });
  return res.json(); */
  return {
    success: true,
    message: "认证成功",
  }
}
