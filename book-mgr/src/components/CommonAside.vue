<script setup lang="ts">
import CommonLogo from "@/components/CommonLogo.vue"
import renderIcon from "@/utils";
import {useRouter} from "vue-router";
import useThemeStore from "@/stores/theme";
import {useUserStore} from "@/stores/userinfo";
import {
  MenuOutline,
  BarChartOutline,
  BookOutline,
  PeopleOutline,
  ClipboardOutline
} from '@vicons/ionicons5'

const themeStore = useThemeStore()
const userStore = useUserStore()
const router = useRouter()

let MenuOption = [
  {
    label: '总览',
    key: 'admin-summary',
    icon: renderIcon(BarChartOutline)
  },
  {
    label: '图书管理',
    key: 'book-mgr',
    icon: renderIcon(BookOutline)
  },
  {
    label: '用户管理',
    key: 'user-mgr',
    icon: renderIcon(PeopleOutline)
  },
  {
    label: '借阅历史',
    key: 'borrow-log',
    icon: renderIcon(ClipboardOutline)
  },
]

let MenuOptionUser = [
  {
    label: '总览',
    key: 'user-summary',
    icon: renderIcon(BarChartOutline)
  },
  {
    label: '借书',
    key: 'bring-books',
    icon: renderIcon(BookOutline)
  },
  {
    label: '我借的书',
    key: 'my-books',
    icon: renderIcon(MenuOutline)
  },
]

let adminUpdate = (value: string) => {
  themeStore.menuSelected = value
  console.log(value)
  switch (value.trim()) {
    // admin
    case 'admin-summary': {
      router.push({path: '/admin/summary'})
      break
    }
    case 'book-mgr': {
      router.push({path: '/admin/book'})
      break
    }
    case 'user-mgr': {
      router.push({path: '/admin/user'})
      break
    }
    case 'borrow-log': {
      router.push({path: '/admin/log'})
      break
    }
    // user
    case 'user-summary': {
      router.push({path: '/user/summary'})
      break
    }
    case 'my-books': {
      router.push({path: '/user/borrowed'})
      break
    }
    case 'bring-books': {
      router.push({path: '/user/borrow'})
      break
    }
  }
}

</script>

<script lang="ts">
export default {
  name: 'CommonAside',
}
</script>

<template>
  <div class="root">
    <CommonLogo></CommonLogo>
   <div class="root-inner">
     <n-menu
         class="menu"
         :accordion="false"
         :default-expand-all="true"
         :options="userStore.thisUser.role==='admin'?MenuOption:MenuOptionUser"
         @update:value="adminUpdate"
         :value="themeStore.menuSelected"
     />
   </div>
  </div>
</template>

<style lang="less" scoped>
.root {
  display: flex;
  flex-direction: column;
  .logo {
  }
  .menu {
    padding: 15px;
  }
  .root-inner {
    height: calc(100vh - 52px);
  }
}
</style>