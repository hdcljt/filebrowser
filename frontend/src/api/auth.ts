// 检查授权状态
export async function checkAuthStatus() {
  /*
  const res = await fetchURL("/api/auth/check", {
    method: "GET"
  });
  return res.json();
  */
  // 返回模拟数据，默认授权状态为需要申请
  return {
    success: true,
    needAuth: true,
    hasApplied: false,
    message: "需要申请授权"
  };
}

// 提交授权申请
export async function submitAuthApply(data: any) {
  /*
  const res = await fetchURL("/api/auth/apply", {
    method: "POST",
    body: JSON.stringify(data),
    headers: {
      "Content-Type": "application/json"
    }
  });
  return res.json();
  */
  // 返回模拟数据，默认申请成功
  return {
    success: true,
    message: "授权申请提交成功，请等待审批"
  };
}

// 验证授权码
export async function verifyAuthCode(code: string) {
  /*
  const res = await fetchURL("/api/auth/verify", {
    method: "POST",
    body: JSON.stringify({ code }),
    headers: {
      "Content-Type": "application/json"
    }
  });
  return res.json();
  */
  // 返回模拟数据，默认验证成功
  return {
    success: true,
    message: "授权验证成功"
  };
}
