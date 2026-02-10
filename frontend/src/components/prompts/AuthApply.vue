<template>
  <div class="card floating" id="auth-apply">
    <div class="card-title">
      <h2>{{ t("prompts.authApply.title") }}</h2>
    </div>

    <div class="card-content">
      <div class="warning">
        <span class="warning-text">{{ remainingTime }}秒后窗口将自动关闭，请尽快操作。</span>
        <button class="close-warning" @click="hideWarning">×</button>
      </div>

      <div class="form-group">
        <label>{{ t("prompts.authApply.mode") }}</label>
        <div class="radio-group">
          <label>
            <input type="radio" v-model="authMode" value="sms" checked />
            短信验证码
          </label>
        </div>
      </div>

      <div class="form-group">
        <label>{{ t("prompts.authApply.approver") }} *</label>
        <select v-model="approver">
          <option value="ty_test1">ty_test1 (18251883836)</option>
          <option value="ty_test2">ty_test2 (13800138000)</option>
          <option value="ty_test3">ty_test3 (13900139000)</option>
        </select>
      </div>

      <div class="form-group">
        <label>手机号:</label>
        <div class="phone-info">{{ phoneNumber }}</div>
      </div>

      <div class="form-group">
        <label>{{ t("prompts.authApply.reason") }} *</label>
        <select v-model="reasonType">
          <option value="">请选择</option>
          <option value="business">业务需求</option>
          <option value="personal">个人使用</option>
          <option value="other">其他原因</option>
        </select>
      </div>

      <div class="form-group">
        <label>{{ t("prompts.authApply.reason") }}详情:</label>
        <textarea v-model="reason" rows="4" :placeholder="t('prompts.authApply.reasonPlaceholder')"></textarea>
      </div>

      <div class="form-group">
        <label>短信验证码 *</label>
        <div class="code-input">
          <input 
            type="text" 
            v-model="verificationCode" 
            placeholder="请输入短信验证码"
          />
          <button 
            class="button get-code-button" 
            @click="getVerificationCode"
            :disabled="isGettingCode || countdown > 0"
          >
            {{ countdown > 0 ? `${countdown}秒后重新获取` : '获取验证码' }}
          </button>
        </div>
      </div>

      <div class="form-actions">
        <button class="button button--outline" @click="cancel">
          取消
        </button>
        <button class="button button--primary" @click="submitApply">
          提交验证
        </button>
      </div>

      <div v-if="error" class="error-message">
        {{ error }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from "vue";
import { useI18n } from "vue-i18n";
import { useLayoutStore } from "@/stores/layout";

const layoutStore = useLayoutStore();
const { t } = useI18n();

// 基本信息
const authMode = ref("sms");
const approver = ref("ty_test1");
const reasonType = ref("");
const reason = ref("");
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

// 审批人手机号映射
const approverPhones: Record<string, string> = {
  ty_test1: "18251883836",
  ty_test2: "13800138000",
  ty_test3: "13900139000"
};

// 计算当前选中审批人的手机号
const phoneNumber = computed(() => {
  return approverPhones[approver.value] || "";
});

// 隐藏警告
const hideWarning = () => {
  showWarning.value = false;
};

// 开始倒计时
const startCountdown = () => {
  countdownTimer = window.setInterval(() => {
    if (countdown.value > 0) {
      countdown.value--;
    }
  }, 1000);

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

  isGettingCode.value = true;
  error.value = "";

  try {
    // 模拟获取验证码的API调用
    console.log("获取验证码:", {
      approver: approver.value,
      phone: phoneNumber.value
    });

    // 模拟成功响应
    setTimeout(() => {
      isGettingCode.value = false;
      countdown.value = 60;
      alert("验证码已发送至" + phoneNumber.value);
    }, 1000);
  } catch (err) {
    isGettingCode.value = false;
    error.value = "获取验证码失败，请重试";
  }
};

// 提交申请
const submitApply = () => {
  // 验证必填字段
  if (!approver.value) {
    error.value = "请选择审批人";
    return;
  }

  if (!reasonType) {
    error.value = "请选择申请原因类型";
    return;
  }

  if (!verificationCode.value.trim()) {
    error.value = "请输入短信验证码";
    return;
  }

  const authData = {
    mode: authMode.value,
    approver: approver.value,
    reasonType: reasonType.value,
    reason: reason.value,
    code: verificationCode.value
  };

  layoutStore.currentPrompt?.confirm(authData);
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
.warning {
  background-color: #fff3cd;
  color: #856404;
  padding: 0.75rem;
  border-radius: 4px;
  margin-bottom: 1.5rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border: 1px solid #ffeaa7;
}

.warning-text {
  font-size: 0.875rem;
}

.close-warning {
  background: none;
  border: none;
  font-size: 1.25rem;
  cursor: pointer;
  color: #856404;
}

.form-group {
  margin-bottom: 1rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
}

.radio-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.radio-group label {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

select, textarea, input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}

textarea {
  resize: vertical;
}

.phone-info {
  padding: 0.5rem;
  background-color: #f8f9fa;
  border: 1px solid #dee2e6;
  border-radius: 4px;
}

.code-input {
  display: flex;
  gap: 0.5rem;
}

.code-input input {
  flex: 1;
}

.get-code-button {
  white-space: nowrap;
  background-color: #007bff;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 4px;
  cursor: pointer;
}

.get-code-button:disabled {
  background-color: #6c757d;
  cursor: not-allowed;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 0.75rem;
  margin-top: 1.5rem;
}

.button--outline {
  background-color: transparent;
  border: 1px solid #007bff;
  color: #007bff;
  padding: 0.5rem 1.5rem;
  border-radius: 4px;
  cursor: pointer;
}

.button--primary {
  background-color: #dc3545;
  color: white;
  border: none;
  padding: 0.5rem 1.5rem;
  border-radius: 4px;
  cursor: pointer;
}

.error-message {
  color: #dc3545;
  margin-top: 1rem;
  font-size: 0.875rem;
}
</style>