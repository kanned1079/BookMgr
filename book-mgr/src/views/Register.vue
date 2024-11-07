<script lang="ts" setup>
import {onMounted, Ref, ref} from 'vue'
import {useRouter} from "vue-router";
import {ChevronForward} from "@vicons/ionicons5"
import {FormInst, FormRules, useMessage} from 'naive-ui'
import instance from "@/axios";
import {encodeToBase64} from "@/utils/base64";

let animated = ref<boolean>(false)
const router = useRouter()
const formRef = ref<FormInst | null>(null)
const message = useMessage()

interface RegisterForm {
  email: string
  password: string
  confirmPassword: string
}

const form: Ref<RegisterForm> = ref({
  email: '',
  password: '',
  confirmPassword: ''
})

const rules: FormRules = {
  email: [
    {required: true, message: '请输入邮箱', trigger: 'blur'},
    {type: 'email', message: '请输入有效的邮箱地址', trigger: 'blur'},
  ],
  password: [
    {required: true, message: '请输入密码', trigger: 'blur'},
    {min: 6, message: '密码长度至少6位', trigger: 'blur'},
  ],
  confirmPassword: [
    {required: true, message: '请再次输入密码', trigger: 'blur'},
    {validator: (rule, value) => value === form.value.password, message: '两次密码输入不一致', trigger: 'blur'},
  ]
}

const onSubmit = () => {
  formRef.value?.validate((errors) => {
    if (!errors) {
      handleRegisterReq()
    } else {
      message.error('请检查输入的内容')
    }
  })
}

const handleRegisterReq = async () => {
  try {
    form.value.password = encodeToBase64(form.value.password)
    let {data} = await instance.post('/api/user/v1/register', {
      email: form.value.email.trim(),
      password: encodeToBase64(form.value.password.trim())
    })
    if (data.code === 200) {
      message.success('注册成功')
      await router.push({path: '/login'})
    } else if (data.code === 409) {
      message.error('该邮箱已注册')
    }
  } catch (err: any) {
    message.error(err + '')
  }
}

onMounted(() => {
  animated.value = true
})
</script>

<script lang="ts">
export default {
  name: 'Register'
}
</script>

<template>
  <div class="root">
    <transition name="slide-fade">
      <n-card class="register-card" :embedded="true" v-if="animated" content-style="padding: 0px">
        <div class="register-card-root">
          <p class="register-card-root-title">创建新帐户</p>
          <p class="register-card-root-app"></p>
        </div>
        <n-form :model="form" :rules="rules" ref="formRef" label-width="100px">
          <n-form-item label="邮箱" path="email">
            <n-input v-model:value="form.email" placeholder="请输入邮箱"/>
          </n-form-item>
          <n-form-item label="密码" path="password">
            <n-input type="password" v-model:value="form.password" placeholder="请输入密码"/>
          </n-form-item>
          <n-form-item label="确认密码" path="confirmPassword">
            <n-input type="password" v-model:value="form.confirmPassword" placeholder="请再次输入密码"/>
          </n-form-item>
          <n-form-item>
            <n-button style="width: 100%" type="primary" @click="onSubmit">注册</n-button>
          </n-form-item>
        </n-form>
        <div style="display: flex; flex-direction: row; justify-content: end; align-items: center">
          <p style="opacity: 0.8">已经有账号？</p>
          <n-button type="primary" text @click="router.replace({ path: '/login' })">
            返回登录
            <n-icon size="16" style="margin-left: 3px">
              <ChevronForward/>
            </n-icon>
          </n-button>
        </div>
      </n-card>
    </transition>
  </div>
</template>

<style lang="less" scoped>
.root {
  display: flex;
  height: 100vh;
  background-color: #e3e5e7;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  .register-card {
    padding: 30px;
    width: 400px;

    .register-card-root {
      margin-bottom: 30px;
      display: flex;
      flex-direction: row;
      align-items: end;

      .register-card-root-title {
        font-size: 1.25rem;
        font-weight: 500;
        margin-right: 3px;
      }

      .register-card-root-app {
        font-size: 1.2rem;
      }
    }
  }
}
</style>