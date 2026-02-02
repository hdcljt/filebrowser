<template>
  <div class="card floating" id="auth-verify">
    <div class="card-title">
      <h2>{{ t("prompts.authVerify") }}</h2>
    </div>

    <div class="card-content">
      <div class="auth-info">
        <p>{{ t("prompts.authVerify.deadline") }}: {{ deadline }}</p>
        <p>{{ t("prompts.authVerify.remaining") }}: {{ remainingSeconds }}秒</p>
        <p class="status" :class="{ approved: status === 'approved' }">
          {{ t("prompts.authVerify.status") }}: {{ t("prompts.authVerify.status" + status) }}
        </p>
      </div>

      <div class="form-group" v-if="status === 'approved'">
        <label>{{ t("prompts.authVerify.code") }}</label>
        <div class="code-input">
          <input 
            type="text" 
            v-model="verificationCode" 
            placeholder='{{ t("prompts.authVerify.codePlaceholder") }}'
            class="input"
          />
          <button class="button" @click="verifyCode">
            {{ t("prompts.authVerify.verify") }}
          </button>
        </div>
      </div>

      <div v-if="error" class="error-message">
        {{ error }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted, computed } from "vue";
import { useI18n } from "vue-i18n";
import { useLayoutStore } from "@/stores/layout";

const layoutStore = useLayoutStore();
const { t } = useI18n();

const deadline = ref("2026-02-02 18:30:00"); // 示例时间
const remainingSeconds = ref(300); // 示例剩余时间
const status = ref("approved"); // 示例状态：approved, pending, rejected
const verificationCode = ref("");
const error = ref("");

let countdownTimer: number;

onMounted(() => {
  startCountdown();
});

onUnmounted(() => {
  if (countdownTimer) {
    clearInterval(countdownTimer);
  }
});

const startCountdown = () => {
  countdownTimer = window.setInterval(() => {
    if (remainingSeconds.value > 0) {
      remainingSeconds.value--;
    } else {
      clearInterval(countdownTimer);
      error.value = t("prompts.authVerify.timeout");
    }
  }, 1000);
};

const verifyCode = () => {
  if (!verificationCode.value.trim()) {
    error.value = t("prompts.authVerify.codeRequired");
    return;
  }

  // 模拟验证逻辑
  if (verificationCode.value.length < 6) {
    error.value = t("prompts.authVerify.codeInvalid");
    return;
  }

  // 验证成功
  layoutStore.currentPrompt?.confirm({ code: verificationCode.value });
};
</script>

<style scoped>
.auth-info {
  margin-bottom: 1.5rem;
}

.auth-info p {
  margin-bottom: 0.5rem;
}

.status {
  font-weight: bold;
}

.status.approved {
  color: #4caf50;
}

.form-group {
  margin-bottom: 1rem;
}

.code-input {
  display: flex;
  gap: 0.5rem;
}

.code-input .input {
  flex: 1;
}

.error-message {
  color: #f44336;
  margin-top: 1rem;
}
</style>