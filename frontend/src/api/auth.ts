import { fetchURL } from "./utils";

// 检查授权状态
export async function checkAuthStatus() {
  const res = await fetchURL("/api/auth/check", {
    method: "GET"
  });
  return res.json();
}

// 提交授权申请
export async function submitAuthApply(data: any) {
  const res = await fetchURL("/api/auth/apply", {
    method: "POST",
    body: JSON.stringify(data),
    headers: {
      "Content-Type": "application/json"
    }
  });
  return res.json();
}

// 验证授权码
export async function verifyAuthCode(code: string) {
  const res = await fetchURL("/api/auth/verify", {
    method: "POST",
    body: JSON.stringify({ code }),
    headers: {
      "Content-Type": "application/json"
    }
  });
  return res.json();
}
