<script setup lang="ts">
import {onMounted, ref} from 'vue'
import {RouterView, useRouter} from 'vue-router';
import CommonHeader from "@/components/CommonHeader.vue";
import CommonAside from "@/components/CommonAside.vue";
import {useUserStore} from "@/stores/userinfo";
// import useApiAddrStore from "@/stores/useApiAddrStore";

// const apiAddrStore = useApiAddrStore();

// const themeStore = useThemeStore();
const userInfoStore = useUserStore();

const router = useRouter();

let interfaceAnimated = ref<boolean>(false)


onMounted(() => {
  // themeStore.contentPath = '/admin/dashboard/summary'
  // themeStore.menuSelected = 'dashboard'
  // console.log('新的dashboard', themeStore.contentPath)
  console.log('dashboard挂载')

  interfaceAnimated.value = true

  // getAllNotices()

  // if (!userInfoStore.isAuthed) {
  //   router.push({
  //     path: '/admin/login',
  //   })
  // }
  // setTimeout(() => {
  //   router.push({
  //     // path: '/admin/dashboard/summary',
  //     path: themeStore.userPath,
  //   })
  // }, 50)
})
</script>

<script lang="ts">
export default {
  name: "DashBoard",
}
</script>

<template>

  <!--  <transition name="slide-fade">-->
  <n-layout v-if="interfaceAnimated" has-sider style="width: 100%; position: fixed; left: 0; top: 0; z-index: 1000">
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
      <!--        <transition name="slide-fade">-->
      <n-layout-header style="position: fixed; z-index: 1001;">
        <CommonHeader></CommonHeader>
      </n-layout-header>
      <!--        </transition>-->

      <!--      <transition name="slide-fade">-->

      <n-layout-content content-style="padding: 0" class="content">
        <n-scrollbar style="max-height: 100vh">


          <!--          <n-config-provider>-->

          <RouterView></RouterView>

          <!--          </n-config-provider>-->


          <div style="height: 3rem"></div>
        </n-scrollbar>
      </n-layout-content>
      <!--      </transition>-->
      <!--      <n-layout-footer>成府路</n-layout-footer>-->
    </n-layout>
  </n-layout>
  <!--  </transition>-->


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