<script setup lang="ts">
import {ref, onMounted} from 'vue'

let animated = ref<boolean>(false)


onMounted(() => {

  animated.value = true
})

</script>

<script lang="ts">
export default {
  name: 'AdminSummary'
}
</script>

<template>
  <div style="padding: 20px 20px 0 20px;">
    <n-card hoverable :embedded="true" :bordered="false">
      <div style="display: flex; flex-direction: row; justify-content: space-between; align-items: center">
        <p style="font-size: 1.1rem; font-weight: 500">总览</p>
        <!--        <n-button size="large" type="primary" @click="editType='add'; showDrawer=true">添加图书</n-button>-->
      </div>
    </n-card>
  </div>

  <transition name="slide-fade">
    <div class="root" v-if="animated">
      <n-card hoverable :embedded="true" :bordered="false">
        <n-statistic label="图书馆总用户有" tabular-nums>
          <n-number-animation ref="numberAnimationInstRef" :from="0" :to="1229" />
          <template #suffix>
            人
          </template>
        </n-statistic>
      </n-card>

      <n-card style="margin-top: 20px" hoverable :embedded="true" :bordered="false">
        <n-statistic label="总借阅量" tabular-nums>
          <n-number-animation ref="numberAnimationInstRef" :from="0" :to="3245245" />
          <template #suffix>
            本书
          </template>
        </n-statistic>
      </n-card>
    </div>
  </transition>


</template>

<style scoped lang="less">
.root {
  padding: 20px;
}

</style>