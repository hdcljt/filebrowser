<template>
  <div class="card floating" id="auth-apply">
    <div class="card-title">
      <h2>下载授权</h2>
    </div>

    <div class="card-content">
      <div class="warning" v-if="showWarning">
        <span class="warning-text">{{ remainingTime }}秒后窗口将自动关闭，请尽快操作。</span>
        <button class="close-warning" @click="hideWarning">×</button>
      </div>

      <div class="form-group">
        <label><span class="required">*</span>选择审批人：</label>
        <select v-model="approver">
          <option value="ty_test1">吴飞鹏 (18251883836)</option>
          <option value="ty_test2">张晓明 (13800138000)</option>
          <option value="ty_test3">李丽 (13900139000)</option>
        </select>
      </div>

      <div class="form-group">
        <label><span class="required">*</span>请选择申请原因类型：</label>
        <select v-model="reasonType">
          <option value="">请选择</option>
          <option value="business">业务需求</option>
          <option value="personal">个人使用</option>
          <option value="other">其他原因</option>
        </select>
      </div>

      <div class="form-group" v-if="reasonType === 'other'">
        <label><span class="required">*</span>自定义原因：</label>
        <textarea 
          v-model="customReason" 
          rows="2" 
          placeholder="请输入具体原因"
          class="input input--block"
        ></textarea>
      </div>

      <div class="form-group">
        <label><span class="required">*</span>短信验证码：</label>
        <div class="code-input-container">
          <input 
            type="text" 
            v-model="verificationCode" 
            placeholder="请输入短信验证码"
            class="input input--block"
          />
          <button 
            class="button" 
            @click="getVerificationCode"
            :disabled="isGettingCode || countdown > 0"
          >
            {{ countdown > 0 ? `${countdown}秒后重新获取` : '获取验证码' }}
          </button>
        </div>
      </div>

      <div v-if="error" class="error-message">
        {{ error }}
      </div>
    </div>

    <div class="card-action">
      <button class="button button--flat button--grey" @click="cancel">
        取消
      </button>
      <button class="button" @click="submitApply">
        提交验证
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue";
import { useLayoutStore } from "@/stores/layout";

const layoutStore = useLayoutStore();

// 基本信息
const approver = ref("ty_test1");
const reasonType = ref("");
const customReason = ref("");
const verificationCode = ref("");

// 验证码相关
const countdown = ref(0);
const isGettingCode = ref(false);
const error = ref("");

// 倒计时相关
const remainingTime = ref(3600);
const showWarning = ref(true);
let countdownTimer: number | null = null;
let warningTimer: number | null = null;

// 隐藏警告
const hideWarning = () => {
  showWarning.value = false;
  // 停止自动关闭倒计时
  if (warningTimer) {
    clearInterval(warningTimer);
    warningTimer = null;
  }
};

// 开始倒计时
const startCountdown = () => {
  // 验证码倒计时
  countdownTimer = window.setInterval(() => {
    if (countdown.value > 0) {
      countdown.value--;
    }
  }, 1000);

  // 自动关闭倒计时
  warningTimer = window.setInterval(() => {
    if (remainingTime.value > 0) {
      remainingTime.value--;
    } else {
      // 时间到，关闭弹窗
      layoutStore.closeHovers();
    }
  }, 1000);
};

// 清除倒计时
const clearCountdown = () => {
  if (countdownTimer) {
    clearInterval(countdownTimer);
    countdownTimer = null;
  }
  if (warningTimer) {
    clearInterval(warningTimer);
    warningTimer = null;
  }
};

// 获取验证码
const getVerificationCode = async () => {
  if (isGettingCode.value || countdown.value > 0) {
    return;
  }

  if (!approver.value) {
    error.value = "请选择审批人";
    return;
  }

  if (!reasonType.value) {
    error.value = "请选择申请原因类型";
    return;
  }

  if (reasonType.value === 'other' && !customReason.value.trim()) {
    error.value = "请输入自定义原因";
    return;
  }

  isGettingCode.value = true;
  error.value = "";

  try {
    // 调用授权申请接口
    const applyData = {
      approver: approver.value,
      reasonType: reasonType.value,
      customReason: reasonType.value === 'other' ? customReason.value : ''
    };

    const response = await import('@/api/auth').then(m => m.submitAuthApply(applyData));
    
    if (response.success) {
      // 模拟成功响应
      setTimeout(() => {
        isGettingCode.value = false;
        countdown.value = 60;
        alert("验证码已发送");
      }, 1000);
    } else {
      throw new Error(response.message || "获取验证码失败");
    }
  } catch (err: any) {
    isGettingCode.value = false;
    error.value = err.message || "获取验证码失败，请重试";
  }
};

// 提交申请
const submitApply = async () => {
  // 验证必填字段
  if (!approver.value) {
    error.value = "请选择审批人";
    return;
  }

  if (!reasonType.value) {
    error.value = "请选择申请原因类型";
    return;
  }

  if (reasonType.value === 'other' && !customReason.value.trim()) {
    error.value = "请输入自定义原因";
    return;
  }

  if (!verificationCode.value.trim()) {
    error.value = "请输入短信验证码";
    return;
  }

  const authData = {
    approver: approver.value,
    reasonType: reasonType.value,
    customReason: reasonType.value === 'other' ? customReason.value : '',
    code: verificationCode.value
  };

  try {
    // 调用授权认证接口
    const response = await import('@/api/auth').then(m => m.verifyAuthCode(verificationCode.value));
    
    if (response.success) {
      // 认证成功，关闭弹窗并继续下载
      layoutStore.currentPrompt?.confirm(authData);
    } else {
      throw new Error(response.message || "授权认证失败");
    }
  } catch (err: any) {
    error.value = err.message || "授权认证失败，请重试";
  }
};

// 取消
const cancel = () => {
  layoutStore.closeHovers();
};

// 组件挂载时开始倒计时
onMounted(() => {
  startCountdown();
});

// 组件卸载时清除倒计时
onUnmounted(() => {
  clearCountdown();
});
</script>

<style scoped>
/* 整体容器样式 */
#auth-apply {
  max-width: 25em;
  width: 90%;
  max-height: 95%;
}

/* 内容区域样式 */
.card-content > *:first-child:not(.warning) {
  padding-top: 1.5em;
}

.card-content > *:last-child {
  padding-bottom: 1.5em;
}

/* 警告框样式 */
.warning {
  background-color: var(--warning-bg, #2a2a1e);
  color: var(--warning-text, #d48806);
  padding: 0.5rem 1rem;
  border-radius: 2px;
  margin-bottom: 1.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border: 1px solid var(--warning-border, #4d4d33);
  margin-top: 0 !important;
  padding-top: 0.5rem !important;
}

.close-warning {
  background: none;
  border: none;
  font-size: 1rem;
  cursor: pointer;
  color: var(--warning-text, #d48806);
  padding: 0;
  width: 16px;
  height: 16px;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* 表单组样式 */
.form-group {
  margin-bottom: 1.5rem;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 0.5rem;
}

/* 标签样式 */
label {
  display: inline-block;
  font-size: 14px;
  color: var(--textSecondary);
  width: 100%;
  margin-bottom: 0.25rem;
}

/* 必填项星号样式 */
.required {
  color: #ff4d4f;
  margin-right: 4px;
}

/* 输入控件样式 */
select, input, textarea {
  width: 100%;
  padding: 8px 12px;
  border: 1px solid var(--borderPrimary);
  border-radius: 2px;
  font-size: 14px;
  transition: all 0.3s;
  background-color: var(--surfacePrimary);
  color: var(--textSecondary);
  margin-bottom: 0;
}

select, input {
  height: 36px;
}

textarea {
  min-height: 60px;
  resize: none;
}

/* 验证码输入区域样式 */
.code-input-container {
  width: 100%;
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.code-input-container input {
  flex: 1;
}

/* 错误消息样式 */
.error-message {
  color: var(--error-color, #ff4d4f);
  margin-top: 1rem;
  font-size: 14px;
}

/* 操作按钮区域样式 */
.card-action {
  text-align: right;
  padding: 1em 1em;
  border-top: 1px solid var(--borderPrimary);
}

.card-action > * {
  margin-left: 0.5rem;
}

/* 响应式调整 */
@media (max-width: 768px) {
  #auth-apply {
    width: 90%;
  }
  
  .form-group {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.5rem;
  }
  
  label {
    width: 100%;
  }
  
  select, input, textarea {
    width: 100%;
  }
  
  .card-action {
    text-align: center;
  }
  
  .card-action > * {
    margin: 0.25rem;
  }
}
</style>