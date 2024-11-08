<script setup lang="ts">
import useThemeStore from "@/stores/theme";
import {useUserStore} from "@/stores/userinfo";
import {useRouter} from "vue-router";
import theme from "@/stores/theme";

const themeStore = useThemeStore()
const userStore = useUserStore()
const router = useRouter()

// console.log(themeStore.bambooGreen.topLogoBgColor)

let backDashBoard = () => {
  themeStore.menuSelected = userStore.thisUser.role==='admin'?'admin-summary':' user-summary'
  router.push({
    path: userStore.thisUser.role==='admin'?'/admin/summary':'/user/summary'
  })
}

</script>

<script lang="ts">
export default {
  name: 'CommonLogo',
}
</script>

<template>
  <div class="root" @click="backDashBoard" :style="{backgroundColor: themeStore.bambooGreen.topLogoBgColor as string}">
    <div class="txt-root">
      <button class="btn">
        <p class="txt" :style="{color: themeStore.bambooGreen.topLogoTextColor}">Books</p>
      </button>
    </div>
  </div>
</template>

<style scoped>
.root {
  text-align: center;
  justify-content: center;
  height: 52px;
  line-height: 52px;

  .txt-root {
    .btn {
      background-color: rgba(255, 255, 255, 0.0);
      border: none;
      line-height: 30px;

      .txt {
        position: relative;
        font-size: 1.25rem;
        font-weight: 100;
      }

      .txt::after {
        content: '';
        position: absolute;
        width: 100%;
        transform: scaleX(0);
        height: 2px;
        bottom: 0;
        left: 0;
        background-color: #fff;
        transform-origin: bottom right;
        transition: transform 0.25s ease-out;
      }

      .txt:hover::after {
        transform: scaleX(1);
        transform-origin: bottom left;
      }
    }
  }
}
</style>