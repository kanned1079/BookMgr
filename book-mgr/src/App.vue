<script setup lang="ts">
import {RouterView} from 'vue-router'
import useThemeStore from "@/stores/theme";
import {type GlobalThemeOverrides, darkTheme} from 'naive-ui'
import theme from "@/stores/theme";
const themeStore = useThemeStore()

// let override = {
//   common: {
//     primaryColor: '#5dac81',
//     primaryColorPressed: '#5dac81',
//     bodyColor: '#eff2f7',
//     modalColor: '#eff2f7',
//     tableColor: '#eff2f7',
//   },
//   Button: {
//     color: '#5dac81',
//     colorFocusPrimary: '#4ca879',
//     colorHoverPrimary: '#4ca879',
//     colorPressedPrimary: '#4ca879',
//     textColorPrimary: '#ffffff',
//     textColorHoverPrimary: '#ffffff',
//     textColorPressedPrimary: '#ffffff',
//     textColorFocusPrimary: '#ffffff',
//     colorTextPrimary: '#5dac81',
//     colorTextHoverPrimary: '#4ca879',
//     textColorPrimaryDisabled: '#ffffff',
//   },
//   Input: {
//     borderHover: '1px solid #5dac81',
//     borderFocus: '1px solid #5dac81',
//   },
//   Card: {
//     colorEmbedded: '#ffffff',
//   },
//   Switch: {
//     railColorActive: '#5dac81',
//   },
//   DataTable: {
//     borderRadius: '3px',
//   },
//   Notification: {
//     color: '#fff',
//     bordered: '1px solid #5dac81',
//   },
//   Dropdown: {
//     color: '#353535',
//   },
// } as GlobalThemeOverrides

// :theme-overrides="themeStore.bambooGreen.selfOverride

</script>

<template>
  <n-config-provider :theme="darkTheme">
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
