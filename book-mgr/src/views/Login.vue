<script lang="ts" setup>
import {onMounted, type Ref, ref} from 'vue'
import {useRouter} from "vue-router";
import {ChevronForward} from "@vicons/ionicons5"
import {type FormInst, type FormRules, useMessage} from 'naive-ui'
import instance from "@/axios";
import {useUserStore} from "@/stores/userinfo";
import {encodeToBase64} from "@/utils/base64";
import useThemeStore from "@/stores/theme";
import theme from "@/stores/theme";

const themeStore = useThemeStore()
const userStore = useUserStore()
const router = useRouter()
let animated = ref<boolean>(false)
const formRef = ref<FormInst | null>(null)
const message = useMessage()

interface LoginForm {
  email: string
  password: string
  role: 'user' | 'admin'
}

const form: Ref<LoginForm> = ref({
  email: '',
  password: '',
  role: 'user',
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
}


const onSubmit = () => {
  formRef.value?.validate((errors) => {
    if (!errors) {
      handleLoginReq()
    } else {
      message.error('请检查输入的内容')
    }
  })
}

let handleLoginReq = async () => {
  try {
    form.value.password = encodeToBase64(form.value.password)
    let {data} = await instance.post(`/api/${form.value.role === 'admin' ? 'admin' : 'user'}/v1/login`, {
      email: form.value.email.trim(),
      password: encodeToBase64(form.value.password.trim()),
      role: form.value.role
    })
    if (data.code === 200 && data.authed) {
      message.success('登录成功')
      Object.assign(userStore.thisUser, data.user)
      sessionStorage.setItem('authed', JSON.stringify(true))
      userStore.authed = true
      themeStore.menuSelected = form.value.role === 'admin' ? 'admin-summary' : 'user-summary'
      await router.push({
        path: form.value.role === 'admin' ? '/admin/summary' : '/user/summary'
      })
    } else if (data.code === 401) {
      message.error('密码不正确')
    } else if (data.code === 403) {
      message.error('非法访问')
    } else if (data.code === 404) {
      message.warning('用户不存在，请先注册。')
    }
  } catch (err: any) {
    message.error(err + '')
  }
}

onMounted(() => {
  let authed = JSON.parse(sessionStorage.getItem('authed') || 'false')
  userStore.authed = authed
  if (authed)
    router.push({
      path: form.value.role === 'admin' ? '/admin/summary' : '/user/borrow'
    })

  animated.value = true
})
</script>

<script lang="ts">
export default {
  name: 'Login'
}
</script>

<template>
  <div class="root">
    <transition name="slide-fade">
      <n-card class="login-card" :embedded="true" v-if="animated" content-style="padding: 0px">
        <div class="login-card-root">
          <p class="login-card-root-title">登录到</p>
          <p class="login-card-root-app">图书管理系统</p>
        </div>
        <n-form :model="form" :rules="rules" ref="formRef" label-width="80px">
          <n-form-item label="邮箱" path="email">
            <n-input v-model:value="form.email" placeholder="请输入邮箱"/>
          </n-form-item>
          <n-form-item label="密码" path="password">
            <n-input type="password" v-model:value="form.password" placeholder="请输入密码"/>
          </n-form-item>
          <n-form-item label="角色" path="role">
            <n-select :options="[{label: '用户', value: 'user'}, {label: '管理员', value: 'admin'}]"
                      :default-value="'user'" v-model:value="form.role" placeholder="请选择登录角色"/>
          </n-form-item>
          <n-form-item>
            <n-button style="width: 100%" type="primary" @click="onSubmit">登录</n-button>
          </n-form-item>
        </n-form>
        <div v-if="form.role!=='admin'"
             style="display: flex; flex-direction: row; justify-content: end; align-items: center">
          <p style="opacity: 0.8">还没有账号？</p>
          <n-button type="primary" text @click="router.push({path: '/register'})">
            立即注册
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

  .login-card {
    padding: 30px;
    width: 400px;

    .login-card-root {
      margin-bottom: 30px;
      display: flex;
      flex-direction: row;
      align-items: end;

      .login-card-root-title {
        font-size: 1.25rem;
        font-weight: 500;
        margin-right: 3px;
      }

      .login-card-root-app {
        font-size: 1.2rem;
      }
    }
  }
}


</style>