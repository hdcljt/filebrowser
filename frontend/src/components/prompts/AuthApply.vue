<template>
  <div class="card floating" id="auth-apply">
    <div class="card-title">
      <h2>{{ t("prompts.authApply.title") }}</h2>
    </div>

    <div class="card-content">
      <div class="form-group">
        <label>{{ t("prompts.authApply.condition") }}</label>
        <div class="radio-group">
          <label>
            <input type="radio" v-model="authCondition" value="count" />
            {{ t("prompts.authApply.count") }}
          </label>
          <label>
            <input type="radio" v-model="authCondition" value="time" />
            {{ t("prompts.authApply.time") }}
            <select v-model="timeDuration" v-if="authCondition === 'time'">
              <option value="1">1小时</option>
              <option value="2">2小时</option>
              <option value="4">4小时</option>
              <option value="8">8小时</option>
              <option value="24">24小时</option>
            </select>
          </label>
        </div>
      </div>

      <div class="form-group">
        <label>{{ t("prompts.authApply.mode") }}</label>
        <select v-model="authMode">
          <option value="remote">{{ t("prompts.authApply.remote") }}</option>
        </select>
      </div>

      <div class="form-group">
        <label>{{ t("prompts.authApply.approver") }}</label>
        <select v-model="approver">
          <option value="ty_test1">ty_test1</option>
          <option value="ty_test2">ty_test2</option>
          <option value="ty_test3">ty_test3</option>
        </select>
      </div>

      <div class="form-group">
        <label>{{ t("prompts.authApply.reason") }}</label>
        <textarea v-model="reason" rows="4" :placeholder="t('prompts.authApply.reasonPlaceholder')"></textarea>
      </div>

      <button class="button button--block" @click="submitApply">
        {{ t("prompts.authApply.submit") }}
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useI18n } from "vue-i18n";
import { useLayoutStore } from "@/stores/layout";

const layoutStore = useLayoutStore();
const { t } = useI18n();

const authCondition = ref("time");
const timeDuration = ref("1");
const authMode = ref("remote");
const approver = ref("ty_test1");
const reason = ref("");

const submitApply = () => {
  if (!reason.value.trim()) {
    alert(t("prompts.authApply.reasonRequired"));
    return;
  }

  const authData = {
    condition: authCondition.value,
    timeDuration: authCondition.value === "time" ? timeDuration.value : "1",
    mode: authMode.value,
    approver: approver.value,
    reason: reason.value
  };

  layoutStore.currentPrompt?.confirm(authData);
};
</script>

<style scoped>
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

select, textarea {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}

textarea {
  resize: vertical;
}
</style>