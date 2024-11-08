<script setup lang="ts">
import {onMounted} from "vue";
import {RouterView} from 'vue-router'
import useThemeStore from "@/stores/theme";
import {useUserStore} from "@/stores/userinfo";
import {type GlobalThemeOverrides, darkTheme} from 'naive-ui'
import theme from "@/stores/theme";
const themeStore = useThemeStore()
const userStore = useUserStore()

// onMounted(() => {
//   // console.log(JSON.parse(sessionStorage.getItem('auth') || false))
//   // userStore.authed = JSON.parse(sessionStorage.getItem('auth') || false)
// })

</script>

<template>
  <n-config-provider :theme="null" :theme-overrides="themeStore.bambooGreen.selfOverride">
    <n-message-provider>
      <n-notification-provider>
        <n-dialog-provider>
          <RouterView></RouterView>
        </n-dialog-provider>
      </n-notification-provider>
    </n-message-provider>
  </n-config-provider>

</template>

<style lang="less">
.slide-fade-enter-active {
  transition: all 250ms ease;
}

.slide-fade-leave-active {
  transition: all 250ms ease;
}

.slide-fade-enter-from {
  transform: translateY(10px);
  opacity: 0;
}
.slide-fade-leave-to {
  //transform: translateY(10px);
  opacity: 0;
}

* {
  padding: 0;
  margin: 0;
  scrollbar-width: none;
  -ms-overflow-style: none;
}

*::-webkit-scrollbar {
  display: none;
}
</style>
