<script lang="ts" setup>
import {useRouter} from "vue-router";
import useThemeStore from "@/stores/theme";
import {useUserStore} from "@/stores/userinfo";
import {useMessage} from "naive-ui";
import renderIcon from "@/utils";
import {
  ChevronDownOutline as downIcon,
  LogOutOutline as LogoutIcon,
  Pencil as EditIcon,
  PersonCircleOutline as UserIcon,
} from '@vicons/ionicons5'

const message = useMessage()
const router = useRouter()
const user_options = [
  // {
  //   label: '用户资料',
  //   key: 'profile',
  //   icon: renderIcon(UserIcon)
  // },
  {
    label: '编辑用户资料',
    key: 'editProfile',
    icon: renderIcon(EditIcon)
  },
  {
    label: '退出登录',
    key: 'logout',
    icon: renderIcon(LogoutIcon)
  }
]

const admin_options = [
  {
    label: '退出登录',
    key: 'logout',
    icon: renderIcon(LogoutIcon)
  }
]

let handleSelect = (key: string) => {
  switch (key) {
    case 'logout': {
      // sessionStorage.removeItem('authed')
      sessionStorage.setItem('authed', JSON.stringify(false))
      router.replace({
        path: '/login'
      })
      break
    }
    case 'editProfile': {
      message.info('功能正在开发别急')
      break
    }
  }
}

const themeStore = useThemeStore()
const userStore = useUserStore()



</script>

<script lang="ts">
export default {
  name: 'CommonHeader'
}
</script>

<template>
  <div class="root" :style="{backgroundColor: themeStore.bambooGreen.topHeaderBgColor}">
    <div class="root-inner">
      <div class="l-content">
        {{ userStore.thisUser.role==='admin'?'管理员':'用户' }}
      </div>
      <div class="r-content">
        <n-dropdown @select="handleSelect" :options="userStore.thisUser.role==='admin'?admin_options:user_options" width="180px" style="color: white">
          <n-button :bordered="false" type="primary" style="color: white">
            <n-icon style="margin-right: 5px" size="18"><UserIcon/></n-icon>
            {{ userStore.thisUser.email }}
            <n-icon style="margin-left: 10px" size="15">
              <downIcon/>
            </n-icon>
          </n-button>
        </n-dropdown>
      </div>
    </div>
  </div>

</template>

<style lang="less" scoped>
.root {
  height: 52px;
  color: white;

  .root-inner {
    display: flex;
    flex-direction: row;
    justify-content: space-between;


    line-height: 52px;
  }

  .l-content {
    margin-left: 20px;
  }

  .r-content {
    margin-right: 20px;
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

}
</style>