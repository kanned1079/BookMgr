<script setup lang="ts">
import {onMounted, ref} from 'vue'

let animated = ref<boolean>(false)


onMounted(() => {

  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'UserSummary'
}
</script>

<template>
  <div style="padding: 20px 20px 0 20px;">
    <n-card hoverable :embedded="true" :bordered="false">
      <div style="display: flex; flex-direction: row; justify-content: space-between; align-items: center">
        <p style="font-size: 1.1rem; font-weight: 500">用户总览</p>
      </div>
    </n-card>
  </div>

  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card hoverable :embedded="true" :bordered="false">
        <n-statistic label="未归还书目数量" tabular-nums>
          <n-number-animation ref="numberAnimationInstRef" :from="0" :to="4"/>
          <template #suffix>
            本
          </template>
        </n-statistic>
      </n-card>

      <n-card style="margin-top: 20px;" hoverable :embedded="true" :bordered="false">
        <div style="display: flex; flex-direction: row; justify-content: left; align-items: center;">
          <n-statistic label="总借阅量 / 在所有用户的百分比" tabular-nums>
            <n-number-animation ref="numberAnimationInstRef" :from="0" :to="229"/>
            本&nbsp;/&nbsp;
            <n-number-animation ref="numberAnimationInstRef" :from="0" :to="32.20" :precision="2"/>
            %
          </n-statistic>
        </div>
      </n-card>
    </div>
  </transition>
</template>

<style scoped lang="less">
.root {
  padding: 20px;
}

</style>