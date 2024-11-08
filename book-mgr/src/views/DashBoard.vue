<script setup lang="ts">
import {onMounted} from 'vue'
import {RouterView} from 'vue-router';
import CommonHeader from "@/components/CommonHeader.vue";
import CommonAside from "@/components/CommonAside.vue";


onMounted(() => {
  console.log('dashboard挂载')

})
</script>

<script lang="ts">
export default {
  name: "DashBoard",
}
</script>

<template>
  <n-layout v-if="true" has-sider style="width: 100%; position: fixed; left: 0; top: 0; z-index: 1000">
    <n-layout-sider
        v-show="true"
        content-style="padding: 0px;"
        :collapsed-width="64"
        :width="240"
        :collapsed="false"
    >
      <CommonAside></CommonAside>
    </n-layout-sider>
    <n-layout>
      <n-layout-header style="position: fixed; z-index: 1001;">
        <CommonHeader></CommonHeader>
      </n-layout-header>
      <n-layout-content content-style="padding: 0" class="content">
        <n-scrollbar style="max-height: 100vh">
          <RouterView></RouterView>
          <div style="height: 3rem"></div>
        </n-scrollbar>
      </n-layout-content>
    </n-layout>
  </n-layout>
</template>

<style lang="less" scoped>
.n-layout-header,
.n-layout-footer {
  padding: 0;
  height: 3.25rem;
  width: 100%;
  overflow: hidden
}

@media screen and (min-width: 768px) {
  .n-layout-header {
    width: calc(100% - 240px);
  }
}

.n-layout-sider {
  //background-color: v-bind('nowTheme.commonAsideBgColor');
  height: 100vh;
  align-items: center;
  //background-color: v-bind('themeStore.getTheme.globeTheme.asideBgColor');
}

.n-layout-content {
  //background-color: v-bind('nowTheme.commonContentBgColor');

  height: 100vh;

}

.logo {
  //width: 240px;
  padding: 0;
  overflow: hidden
}

.content {
  margin-top: 52px;
  //background-color: v-bind('themeStore.getTheme.globeTheme.contentBgColor');
}

</style>